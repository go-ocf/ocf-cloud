package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/plgd-dev/cloud/pkg/security/oauth/manager"
	kitSync "github.com/plgd-dev/kit/sync"
	"go.uber.org/zap"

	"github.com/plgd-dev/cloud/pkg/sync/task/queue"

	"google.golang.org/grpc"

	pbAS "github.com/plgd-dev/cloud/authorization/pb"
	"github.com/plgd-dev/cloud/coap-gateway/uri"
	pbGRPC "github.com/plgd-dev/cloud/grpc-gateway/pb"
	kitNetGrpc "github.com/plgd-dev/cloud/pkg/net/grpc"
	grpcClient "github.com/plgd-dev/cloud/pkg/net/grpc/client"
	certManagerServer "github.com/plgd-dev/cloud/pkg/security/certManager/server"
	raClient "github.com/plgd-dev/cloud/resource-aggregate/client"
	"github.com/plgd-dev/cloud/resource-aggregate/cqrs/eventbus/nats/subscriber"
	"github.com/plgd-dev/cloud/resource-aggregate/cqrs/utils"
	coapCodes "github.com/plgd-dev/go-coap/v2/message/codes"
	"github.com/plgd-dev/go-coap/v2/mux"
	"github.com/plgd-dev/go-coap/v2/net"
	"github.com/plgd-dev/go-coap/v2/net/blockwise"
	"github.com/plgd-dev/go-coap/v2/net/monitor/inactivity"
	"github.com/plgd-dev/go-coap/v2/tcp"
	"github.com/plgd-dev/go-coap/v2/tcp/message/pool"

	cache "github.com/patrickmn/go-cache"
	"github.com/plgd-dev/cloud/pkg/log"
)

var authCtxKey = "AuthCtx"

//Service a configuration of coapgateway
type Service struct {
	config Config

	KeepaliveOnInactivity   func(cc inactivity.ClientConn)
	BlockWiseTransferSZX    blockwise.SZX
	raClient                *raClient.Client
	asClient                pbAS.AuthorizationServiceClient
	rdClient                pbGRPC.GrpcGatewayClient
	oauthMgr                *manager.Manager
	expirationClientCache   *cache.Cache
	coapServer              *tcp.Server
	listener                tcp.Listener
	authInterceptor         Interceptor
	ctx                     context.Context
	cancel                  context.CancelFunc
	taskQueue               *queue.Queue
	userDeviceSubscriptions *kitSync.Map
	devicesStatusUpdater    *devicesStatusUpdater
	resourceSubscriber      *subscriber.Subscriber
	sigs                    chan os.Signal
}

type DialCertManager = interface {
	GetClientTLSConfig() *tls.Config
}

type ListenCertManager = interface {
	GetServerTLSConfig() *tls.Config
}

// New creates server.
func New(ctx context.Context, config Config, logger *zap.Logger) (*Service, error) {
	p, err := queue.New(config.TaskQueue)
	if err != nil {
		return nil, fmt.Errorf("cannot job queue %w", err)
	}

	resourceSubscriber, err := subscriber.New(config.Clients.Eventbus.NATS, logger, subscriber.WithGoPool(func(f func()) error { return p.Submit(f) }), subscriber.WithUnmarshaler(utils.Unmarshal))
	if err != nil {
		p.Release()
		return nil, fmt.Errorf("cannot create eventbus subscriber: %w", err)
	}
	resourceSubscriber.AddCloseFunc(p.Release)

	expirationClientCache := cache.New(cache.NoExpiration, time.Minute)
	expirationClientCache.OnEvicted(func(deviceID string, c interface{}) {
		if c == nil {
			return
		}
		client := c.(*Client)
		authCtx, err := client.GetAuthorizationContext()
		if err != nil {
			client.Close()
			log.Debugf("device %v token has ben expired", authCtx.GetDeviceID())
		}
	})

	oauthMgr, err := manager.New(config.Clients.AuthServer.OAuth, logger)
	if err != nil {
		resourceSubscriber.Close()
		return nil, fmt.Errorf("cannot create oauth manager: %w", err)
	}
	resourceSubscriber.AddCloseFunc(oauthMgr.Close)

	raConn, err := grpcClient.New(config.Clients.ResourceAggregate.Connection, logger, grpc.WithPerRPCCredentials(kitNetGrpc.NewOAuthAccess(oauthMgr.GetToken)))
	if err != nil {
		resourceSubscriber.Close()
		return nil, fmt.Errorf("cannot create connection to resource aggregate: %w", err)
	}
	resourceSubscriber.AddCloseFunc(func() {
		err := raConn.Close()
		if err != nil {
			logger.Sugar().Errorf("error occurs during close connection to resource aggregate: %v", err)
		}
	})
	raClient := raClient.New(raConn.GRPC(), resourceSubscriber)

	asConn, err := grpcClient.New(config.Clients.AuthServer.Connection, logger, grpc.WithPerRPCCredentials(kitNetGrpc.NewOAuthAccess(oauthMgr.GetToken)))
	if err != nil {
		resourceSubscriber.Close()
		return nil, fmt.Errorf("cannot create connection to authorization server: %w", err)
	}
	resourceSubscriber.AddCloseFunc(func() {
		err := asConn.Close()
		if err != nil {
			logger.Sugar().Errorf("error occurs during close connection to authorization server: %v", err)
		}
	})
	asClient := pbAS.NewAuthorizationServiceClient(asConn.GRPC())

	rdConn, err := grpcClient.New(config.Clients.ResourceDirectory.Connection, logger, grpc.WithPerRPCCredentials(kitNetGrpc.NewOAuthAccess(oauthMgr.GetToken)))
	if err != nil {
		resourceSubscriber.Close()
		return nil, fmt.Errorf("cannot create connection to resource directory: %w", err)
	}
	resourceSubscriber.AddCloseFunc(func() {
		err := asConn.Close()
		if err != nil {
			logger.Sugar().Errorf("error occurs during close connection to resource directory: %v", err)
		}
	})
	rdClient := pbGRPC.NewGrpcGatewayClient(rdConn.GRPC())

	var listener tcp.Listener
	if !config.APIs.COAP.TLS.Enabled {
		l, err := net.NewTCPListener("tcp", config.APIs.COAP.Addr)
		if err != nil {
			resourceSubscriber.Close()
			return nil, fmt.Errorf("cannot create tcp listener: %w", err)
		}
		resourceSubscriber.AddCloseFunc(func() {
			l.Close()
		})
		listener = l
	} else {
		coapsTLS, err := certManagerServer.New(config.APIs.COAP.TLS.Embedded, logger)
		if err != nil {
			resourceSubscriber.Close()
			return nil, fmt.Errorf("cannot create tls cert manager: %w", err)
		}
		resourceSubscriber.AddCloseFunc(coapsTLS.Close)
		l, err := net.NewTLSListener("tcp", config.APIs.COAP.Addr, coapsTLS.GetTLSConfig())
		if err != nil {
			resourceSubscriber.Close()
			return nil, fmt.Errorf("cannot create tcp-tls listener: %w", err)
		}
		resourceSubscriber.AddCloseFunc(func() {
			l.Close()
		})
		listener = l
	}

	onInactivity := func(cc inactivity.ClientConn) {
		cc.Close()
		client, ok := cc.Context().Value(clientKey).(*Client)
		if ok {
			deviceID := getDeviceID(client)
			log.Errorf("DeviceId: %v: keep alive was reached fail limit:: closing connection", deviceID)
		} else {
			log.Errorf("keep alive was reached fail limit:: closing connection")
		}
	}

	blockWiseTransferSZX := blockwise.SZX1024
	if config.APIs.COAP.BlockwiseTransfer.Enabled {
		switch strings.ToLower(config.APIs.COAP.BlockwiseTransfer.SZX) {
		case "16":
			blockWiseTransferSZX = blockwise.SZX16
		case "32":
			blockWiseTransferSZX = blockwise.SZX32
		case "64":
			blockWiseTransferSZX = blockwise.SZX64
		case "128":
			blockWiseTransferSZX = blockwise.SZX128
		case "256":
			blockWiseTransferSZX = blockwise.SZX256
		case "512":
			blockWiseTransferSZX = blockwise.SZX512
		case "1024":
			blockWiseTransferSZX = blockwise.SZX1024
		case "bert":
			blockWiseTransferSZX = blockwise.SZXBERT
		default:
			resourceSubscriber.Close()
			return nil, fmt.Errorf("invalid value BlockWiseTransferSZX %v", config.APIs.COAP.BlockwiseTransfer.SZX)
		}
	}

	ctx, cancel := context.WithCancel(ctx)

	s := Service{
		config:                config,
		BlockWiseTransferSZX:  blockWiseTransferSZX,
		KeepaliveOnInactivity: onInactivity,

		raClient:              raClient,
		asClient:              asClient,
		rdClient:              rdClient,
		oauthMgr:              oauthMgr,
		expirationClientCache: expirationClientCache,
		listener:              listener,
		authInterceptor:       NewAuthInterceptor(),
		devicesStatusUpdater:  NewDevicesStatusUpdater(ctx, config.Clients.ResourceAggregate.DeviceStatusExpiration),

		sigs: make(chan os.Signal, 1),

		taskQueue:          p,
		resourceSubscriber: resourceSubscriber,

		ctx:    ctx,
		cancel: cancel,

		userDeviceSubscriptions: kitSync.NewMap(),
	}

	s.setupCoapServer()

	return &s, nil
}

func getDeviceID(client *Client) string {
	deviceID := "unknown"
	if client != nil {
		authCtx, _ := client.GetAuthorizationContext()
		deviceID = authCtx.GetDeviceID()
		if deviceID == "" {
			deviceID = fmt.Sprintf("unknown(%v)", client.remoteAddrString())
		}
	}
	return deviceID
}

func validateCommand(s mux.ResponseWriter, req *mux.Message, server *Service, fnc func(req *mux.Message, client *Client)) {
	client, ok := s.Client().Context().Value(clientKey).(*Client)
	if !ok || client == nil {
		client = newClient(server, s.Client().ClientConn().(*tcp.ClientConn))
	}
	err := server.taskQueue.Submit(func() {
		switch req.Code {
		case coapCodes.POST, coapCodes.DELETE, coapCodes.PUT, coapCodes.GET:
			fnc(req, client)
		case coapCodes.Empty:
			if !ok {
				client.logAndWriteErrorResponse(fmt.Errorf("cannot handle command: client not found"), coapCodes.InternalServerError, req.Token)
				return
			}
			clientResetHandler(req, client)
		case coapCodes.Content:
			// Unregistered observer at a peer send us a notification
			deviceID := getDeviceID(client)
			tmp, err := pool.ConvertFrom(req.Message)
			if err != nil {
				log.Errorf("DeviceId: %v: cannot convert dropped notification: %v", deviceID, err)
			} else {
				decodeMsgToDebug(client, tmp, "DROPPED-NOTIFICATION")
			}
		default:
			deviceID := getDeviceID(client)
			log.Errorf("DeviceId: %v: received invalid code: CoapCode(%v)", deviceID, req.Code)
		}
	})
	if err != nil {
		deviceID := getDeviceID(client)
		client.Close()
		log.Errorf("DeviceId: %v: cannot handle request %v by task queue: %v", deviceID, req.String(), err)
	}
}

func defaultHandler(req *mux.Message, client *Client) {
	path, _ := req.Options.Path()

	switch {
	case strings.HasPrefix("/"+path, uri.ResourceRoute):
		resourceRouteHandler(req, client)
	default:
		deviceID := getDeviceID(client)
		client.logAndWriteErrorResponse(fmt.Errorf("DeviceId: %v: unknown path %v", deviceID, path), coapCodes.NotFound, req.Token)
	}
}

const clientKey = "client"

func (server *Service) coapConnOnNew(coapConn *tcp.ClientConn, tlscon *tls.Conn) {
	client := newClient(server, coapConn)
	coapConn.SetContextValue(clientKey, client)
	coapConn.AddOnClose(func() {
		client.OnClose()
	})
}

func (server *Service) loggingMiddleware(next mux.Handler) mux.Handler {
	return mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
		client, ok := w.Client().Context().Value(clientKey).(*Client)
		if !ok {
			client = newClient(server, w.Client().ClientConn().(*tcp.ClientConn))
		}
		tmp, err := pool.ConvertFrom(r.Message)
		if err != nil {
			client.logAndWriteErrorResponse(fmt.Errorf("cannot convert from mux.Message: %w", err), coapCodes.InternalServerError, r.Token)
			return
		}
		decodeMsgToDebug(client, tmp, "RECEIVED-COMMAND")
		next.ServeCOAP(w, r)
	})
}

func (server *Service) authMiddleware(next mux.Handler) mux.Handler {
	return mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
		client, ok := w.Client().Context().Value(clientKey).(*Client)
		if !ok {
			client = newClient(server, w.Client().ClientConn().(*tcp.ClientConn))
		}
		authCtx, _ := client.GetAuthorizationContext()
		ctx := context.WithValue(r.Context, &authCtxKey, authCtx)
		path, _ := r.Options.Path()
		_, err := server.authInterceptor(ctx, r.Code, "/"+path)
		if err != nil {
			client.logAndWriteErrorResponse(fmt.Errorf("cannot handle request to path '%v': %w", path, err), coapCodes.Unauthorized, r.Token)
			client.Close()
			return
		}
		serviceToken, err := server.oauthMgr.GetToken(r.Context)
		if err != nil {
			client.logAndWriteErrorResponse(fmt.Errorf("cannot get service token: %w", err), coapCodes.InternalServerError, r.Token)
			client.Close()
			return
		}
		r.Context = kitNetGrpc.CtxWithOwner(kitNetGrpc.CtxWithToken(r.Context, serviceToken.AccessToken), authCtx.GetUserID())
		next.ServeCOAP(w, r)
	})
}

func (server *Service) ServiceRequestContext(userID string) (context.Context, error) {
	serviceToken, err := server.oauthMgr.GetToken(server.ctx)
	if err != nil {
		return nil, err
	}
	return kitNetGrpc.CtxWithOwner(kitNetGrpc.CtxWithToken(server.ctx, serviceToken.AccessToken), userID), nil
}

//setupCoapServer setup coap server
func (server *Service) setupCoapServer() {
	m := mux.NewRouter()
	m.Use(server.loggingMiddleware, server.authMiddleware)
	m.DefaultHandle(mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
		validateCommand(w, r, server, defaultHandler)
	}))
	m.Handle(uri.ResourceDirectory, mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
		validateCommand(w, r, server, resourceDirectoryHandler)
	}))
	m.Handle(uri.SignUp, mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
		validateCommand(w, r, server, signUpHandler)
	}))
	m.Handle(uri.SignIn, mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
		validateCommand(w, r, server, signInHandler)
	}))
	m.Handle(uri.ResourceDiscovery, mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
		validateCommand(w, r, server, resourceDiscoveryHandler)
	}))
	m.Handle(uri.RefreshToken, mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
		validateCommand(w, r, server, refreshTokenHandler)
	}))

	opts := make([]tcp.ServerOption, 0, 8)
	opts = append(opts, tcp.WithKeepAlive(1, server.config.APIs.COAP.KeepAlive.Timeout, server.KeepaliveOnInactivity))
	opts = append(opts, tcp.WithOnNewClientConn(server.coapConnOnNew))
	opts = append(opts, tcp.WithBlockwise(server.config.APIs.COAP.BlockwiseTransfer.Enabled, server.BlockWiseTransferSZX, server.config.APIs.COAP.KeepAlive.Timeout))
	opts = append(opts, tcp.WithMux(m))
	opts = append(opts, tcp.WithContext(server.ctx))
	opts = append(opts, tcp.WithHeartBeat(server.config.APIs.COAP.GoroutineSocketHeartbeat))
	opts = append(opts, tcp.WithMaxMessageSize(server.config.APIs.COAP.MaxMessageSize))
	opts = append(opts, tcp.WithGoPool(func(f func()) error {
		// we call directly function in connection-goroutine because
		// pairing request/response cannot be done in taskQueue for a observe resource.
		// - the observe resource creates task which wait for the response and this wait can be infinite
		// if all task goroutines are processing observations and they are waiting for the responses, which
		// will be stored in task queue.  it happens when we use task queue here.
		f()
		return nil
	}))
	server.coapServer = tcp.NewServer(opts...)
}

func (server *Service) tlsEnabled() bool {
	return server.config.APIs.COAP.TLS.Enabled
}

// Serve starts a coapgateway on the configured address in *Service.
func (server *Service) Serve() error {
	return server.serveWithHandlingSignal()
}

func (server *Service) serveWithHandlingSignal() error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)
	go func(server *Service) {
		defer wg.Done()
		err = server.coapServer.Serve(server.listener)
		server.cancel()
		server.resourceSubscriber.Close()
	}(server)

	signal.Notify(server.sigs,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-server.sigs

	server.coapServer.Stop()
	wg.Wait()

	return err
}

// Shutdown turn off server.
func (server *Service) Close() error {
	select {
	case server.sigs <- syscall.SIGTERM:
	default:
	}
	return nil
}
