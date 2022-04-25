package server

import (
	context "context"
	"math"
	"strings"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/plgd-dev/hub/v2/grpc-gateway/pb"
	"github.com/plgd-dev/hub/v2/pkg/log"
	kitNetGrpc "github.com/plgd-dev/hub/v2/pkg/net/grpc"
	"github.com/plgd-dev/hub/v2/pkg/security/jwt"
	"github.com/plgd-dev/hub/v2/resource-aggregate/commands"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const grpcPrefixKey = "grpc"
const requestKey = "request"

var defaultCodeToLevel = map[codes.Code]zapcore.Level{
	codes.OK:                 zap.DebugLevel,
	codes.Canceled:           zap.DebugLevel,
	codes.Unknown:            zap.ErrorLevel,
	codes.InvalidArgument:    zap.DebugLevel,
	codes.DeadlineExceeded:   zap.WarnLevel,
	codes.NotFound:           zap.DebugLevel,
	codes.AlreadyExists:      zap.DebugLevel,
	codes.PermissionDenied:   zap.WarnLevel,
	codes.Unauthenticated:    zap.DebugLevel,
	codes.ResourceExhausted:  zap.WarnLevel,
	codes.FailedPrecondition: zap.WarnLevel,
	codes.Aborted:            zap.WarnLevel,
	codes.OutOfRange:         zap.WarnLevel,
	codes.Unimplemented:      zap.ErrorLevel,
	codes.Internal:           zap.ErrorLevel,
	codes.Unavailable:        zap.WarnLevel,
	codes.DataLoss:           zap.ErrorLevel,
}

// DefaultCodeToLevel is the default implementation of gRPC return codes and interceptor log level for server side.
func DefaultCodeToLevel(code codes.Code) zapcore.Level {
	lvl, ok := defaultCodeToLevel[code]
	if ok {
		return lvl
	}
	return zap.ErrorLevel
}

func setLogBasicLabels(m map[string]interface{}, req interface{}) {
	if d, ok := req.(interface{ GetDeviceId() string }); ok && d.GetDeviceId() != "" {
		log.SetLogValue(m, log.DeviceIDKey, d.GetDeviceId())
	}
	if r, ok := req.(interface{ GetResourceId() *commands.ResourceId }); ok {
		log.SetLogValue(m, log.DeviceIDKey, r.GetResourceId().GetDeviceId())
		log.SetLogValue(m, log.ResourceHrefKey, r.GetResourceId().GetHref())
	}
	if r, ok := req.(interface{ GetCorrelationId() string }); ok && r.GetCorrelationId() != "" {
		log.SetLogValue(m, log.CorrelationIDKey, r.GetCorrelationId())
	}
}

func setLogFilterLabels(m map[string]interface{}, req interface{}) {
	if req == nil {
		return
	}
	if r, ok := req.(interface {
		GetCommandFilter() []pb.GetPendingCommandsRequest_Command
	}); ok {
		commandFiler := make([]string, 0, len(r.GetCommandFilter()))
		for _, f := range r.GetCommandFilter() {
			commandFiler = append(commandFiler, f.String())
		}
		log.SetLogValue(m, log.CommandFilterKey, commandFiler)
	}
	if r, ok := req.(interface{ GetResourceIdFilter() []string }); ok {
		log.SetLogValue(m, log.ResourceIDFilterKey, r.GetResourceIdFilter())
	}
	if r, ok := req.(interface{ GetDeviceIdFilter() []string }); ok {
		log.SetLogValue(m, log.DeviceIDFilterKey, r.GetDeviceIdFilter())
	}
	if r, ok := req.(interface{ GetTypeFilter() []string }); ok {
		log.SetLogValue(m, log.TypeFilterKey, r.GetTypeFilter())
	}
}

func setLogSubscriptionLabels(m map[string]interface{}, sub *pb.SubscribeToEvents) {
	switch sub.GetAction().(type) {
	case *pb.SubscribeToEvents_CreateSubscription_:
		m[log.SubActionKey] = "createSubscription"
	case *pb.SubscribeToEvents_CancelSubscription_:
		m[log.SubActionKey] = "cancelSubscription"
	}
	setLogFilterLabels(m, sub.GetCreateSubscription())
	eventFilter := make([]string, 0, len(sub.GetCreateSubscription().GetEventFilter()))
	for _, e := range sub.GetCreateSubscription().GetEventFilter() {
		eventFilter = append(eventFilter, e.String())
	}
	log.SetLogValue(m, log.EventFilterKey, eventFilter)
	log.SetLogValue(m, log.CorrelationIDKey, sub.GetCorrelationId())
}

// CodeGenRequestFieldExtractor is a function that relies on code-generated functions that export log fields from requests.
// These are usually coming from a protoc-plugin that generates additional information based on custom field options.
func CodeGenRequestFieldExtractor(fullMethod string, req interface{}) map[string]interface{} {
	m := grpc_ctxtags.CodeGenRequestFieldExtractor(fullMethod, req)
	if m == nil {
		m = make(map[string]interface{})
	}
	setLogBasicLabels(m, req)
	setLogFilterLabels(m, req)
	if sub, ok := req.(*pb.SubscribeToEvents); ok {
		setLogSubscriptionLabels(m, sub)
	}
	method := strings.SplitAfterN(fullMethod, "/", 3)
	if len(method) == 3 {
		m["service"] = strings.ReplaceAll(method[1], "/", "")
		m[log.MethodKey] = strings.ReplaceAll(method[2], "/", "")
	}
	m[log.StartTimeKey] = time.Now()

	if len(m) > 0 {
		return m
	}
	return nil
}

func defaultMessageProducer(ctx context.Context, ctxLogger context.Context, msg string, level zapcore.Level, code codes.Code, err error, duration zapcore.Field) {
	req := make(map[string]interface{})
	resp := make(map[string]interface{})
	resp["code"] = code.String()
	if err != nil {
		resp["error"] = err.Error()
	}

	if sub, err := kitNetGrpc.OwnerFromTokenMD(ctx, "sub"); err == nil {
		req[log.JWTKey] = map[string]string{
			log.SubKey: sub,
		}
	}
	tags := grpc_ctxtags.Extract(ctx)
	newTags := grpc_ctxtags.NewTags()
	newTags.Set(log.DurationMSKey, math.Float32frombits(uint32(duration.Integer)))
	newTags.Set(log.ProtocolKey, "GRPC")
	for k, v := range tags.Values() {
		if strings.EqualFold(k, grpcPrefixKey+"."+requestKey+"."+log.StartTimeKey) {
			newTags.Set(log.StartTimeKey, v)
			continue
		}
		if strings.EqualFold(k, grpcPrefixKey+"."+requestKey+"."+log.DeviceIDKey) {
			newTags.Set(log.DeviceIDKey, v)
			continue
		}
		if strings.HasPrefix(k, grpcPrefixKey+"."+requestKey+".") {
			req[strings.TrimPrefix(k, grpcPrefixKey+"."+requestKey+".")] = v
		}
	}
	if len(req) > 0 {
		newTags.Set(log.RequestKey, req)
	}
	if len(resp) > 0 {
		newTags.Set(log.ResponseKey, resp)
	}
	if deadline, ok := ctx.Deadline(); ok {
		newTags.Set(log.DeadlineKey, deadline)
	}
	ctx = grpc_ctxtags.SetInContext(ctxLogger, newTags)

	ctxzap.Extract(ctx).Check(level, msg).Write()
}

func MakeDefaultMessageProducer(logger *zap.Logger) func(ctx context.Context, msg string, level zapcore.Level, code codes.Code, err error, duration zapcore.Field) {
	ctxLogger := ctxzap.ToContext(context.Background(), logger)
	return func(ctx context.Context, msg string, level zapcore.Level, code codes.Code, err error, duration zapcore.Field) {
		defaultMessageProducer(ctx, ctxLogger, msg, level, code, err, duration)
	}
}

type GetDeviceIDPb interface {
	GetDeviceId() string
}

func MakeDefaultOptions(auth kitNetGrpc.AuthInterceptors, logger log.Logger) ([]grpc.ServerOption, error) {
	streamInterceptors := []grpc.StreamServerInterceptor{}
	unaryInterceptors := []grpc.UnaryServerInterceptor{}
	zapLogger, ok := logger.Unwrap().(*zap.SugaredLogger)
	if ok {
		cfg := logger.Config()
		if cfg.EncoderConfig.EncodeTime.TimeEncoder == nil {
			cfg.EncoderConfig.EncodeTime = log.MakeDefaultConfig().EncoderConfig.EncodeTime
		}
		streamInterceptors = append(streamInterceptors, grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(zapLogger.Desugar(), grpc_zap.WithTimestampFormat(cfg.EncoderConfig.EncodeTime.TimeEncoder.TimeString()), grpc_zap.WithLevels(DefaultCodeToLevel), grpc_zap.WithMessageProducer(MakeDefaultMessageProducer(zapLogger.Desugar()))))
		unaryInterceptors = append(unaryInterceptors, grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(zapLogger.Desugar(), grpc_zap.WithTimestampFormat(cfg.EncoderConfig.EncodeTime.TimeEncoder.TimeString()), grpc_zap.WithLevels(DefaultCodeToLevel), grpc_zap.WithMessageProducer(MakeDefaultMessageProducer(zapLogger.Desugar()))))
	}
	streamInterceptors = append(streamInterceptors, auth.Stream())
	unaryInterceptors = append(unaryInterceptors, auth.Unary())

	return []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			streamInterceptors...,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			unaryInterceptors...,
		)),
	}, nil
}

type cfg struct {
	disableTokenForwarding bool
	whiteListedMethods     []string
}

type Option func(*cfg)

func WithDisabledTokenForwarding() Option {
	return func(c *cfg) {
		c.disableTokenForwarding = true
	}
}

func WithWhiteListedMethods(method ...string) Option {
	return func(c *cfg) {
		c.whiteListedMethods = append(c.whiteListedMethods, method...)
	}
}

func NewAuth(validator kitNetGrpc.Validator, opts ...Option) kitNetGrpc.AuthInterceptors {
	interceptor := kitNetGrpc.ValidateJWTWithValidator(validator, func(ctx context.Context, method string) kitNetGrpc.Claims {
		return jwt.NewScopeClaims()
	})
	var cfg cfg
	for _, o := range opts {
		o(&cfg)
	}
	return kitNetGrpc.MakeAuthInterceptors(func(ctx context.Context, method string) (context.Context, error) {
		ctx, err := interceptor(ctx, method)
		if err != nil {
			log.Errorf("auth interceptor %v: %w", method, err)
			return ctx, err
		}

		if !cfg.disableTokenForwarding {
			if token, err := kitNetGrpc.TokenFromMD(ctx); err == nil {
				ctx = kitNetGrpc.CtxWithToken(ctx, token)
			}
		}

		return ctx, nil
	}, cfg.whiteListedMethods...)
}
