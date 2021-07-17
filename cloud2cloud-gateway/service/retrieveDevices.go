package service

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	pbGRPC "github.com/plgd-dev/cloud/grpc-gateway/pb"
	kitNetGrpc "github.com/plgd-dev/cloud/pkg/net/grpc"
	kitNetHttp "github.com/plgd-dev/cloud/pkg/net/http"
	"github.com/plgd-dev/sdk/schema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Status string

const Status_ONLINE Status = "online"
const Status_OFFLINE Status = "offline"

func toStatus(isOnline bool) Status {
	if isOnline {
		return "online"
	}
	return "offline"
}

type responseWriterEncoderFunc func(w http.ResponseWriter, v interface{}, status int) error

type Device struct {
	Device schema.Device `json:"device"`
	Status Status        `json:"status"`
}

func (rh *RequestHandler) GetDevices(ctx context.Context, deviceIdFilter []string) ([]Device, error) {
	getDevicesClient, err := rh.rdClient.GetDevices(ctx, &pbGRPC.GetDevicesRequest{
		DeviceIdFilter: deviceIdFilter,
	})

	if err != nil {
		return nil, fmt.Errorf("cannot get devices: %w", err)
	}
	defer getDevicesClient.CloseSend()

	devices := make([]Device, 0, 32)
	for {
		device, err := getDevicesClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("cannot get devices: %w", err)
		}

		devices = append(devices, Device{
			Device: device.ToSchema(),
			Status: toStatus(device.GetMetadata().GetStatus().IsOnline()),
		})
	}
	if len(devices) == 0 {
		return nil, status.Errorf(codes.NotFound, "cannot get devices: not found")
	}
	return devices, nil
}

func (rh *RequestHandler) RetrieveDevicesBase(ctx context.Context, w http.ResponseWriter, encoder responseWriterEncoderFunc) (int, error) {
	devices, err := rh.GetDevices(ctx, nil)
	if err != nil {
		return kitNetHttp.ErrToStatusWithDef(err, http.StatusForbidden), fmt.Errorf("cannot retrieve all devices[base]: %w", err)
	}
	resourceLink, err := rh.GetResourceLinks(ctx, nil)
	if err != nil {
		return kitNetHttp.ErrToStatusWithDef(err, http.StatusForbidden), fmt.Errorf("cannot retrieve all devices[base]: %w", err)
	}

	resp := make([]RetrieveDeviceWithLinksResponse, 0, 32)
	for _, dev := range devices {
		links, ok := resourceLink[dev.Device.ID]
		if ok {
			resp = append(resp, RetrieveDeviceWithLinksResponse{
				Device: dev,
				Links:  links,
			})
		}
	}

	err = encoder(w, resp, http.StatusOK)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("cannot retrieve all devices[base]: %w", err)
	}
	return http.StatusOK, nil
}

func (rh *RequestHandler) RetrieveDevicesAll(ctx context.Context, w http.ResponseWriter, encoder responseWriterEncoderFunc) (int, error) {
	devices, err := rh.GetDevices(ctx, nil)
	if err != nil {
		return kitNetHttp.ErrToStatusWithDef(err, http.StatusForbidden), fmt.Errorf("cannot retrieve all devices[base]: %w", err)
	}
	reps, err := rh.RetrieveResources(ctx, nil, nil)
	if err != nil {
		return kitNetHttp.ErrToStatusWithDef(err, http.StatusForbidden), fmt.Errorf("cannot retrieve all devices[base]: %w", err)
	}

	resp := make([]RetrieveDeviceContentAllResponse, 0, 32)
	for _, dev := range devices {
		devReps, ok := reps[dev.Device.ID]
		if ok {
			resp = append(resp, RetrieveDeviceContentAllResponse{
				Device: dev,
				Links:  devReps,
			})
		}
	}

	err = encoder(w, resp, http.StatusOK)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("cannot retrieve all devices[base]: %w", err)
	}
	return http.StatusOK, nil
}

func (rh *RequestHandler) RetrieveDevicesWithContentQuery(ctx context.Context, w http.ResponseWriter, routeVars map[string]string, contentQuery string, encoder responseWriterEncoderFunc) (statusCode int, err error) {
	switch contentQuery {
	case ContentQueryAllValue:
		statusCode, err = rh.RetrieveDevicesAll(ctx, w, encoder)
	case ContentQueryBaseValue:
		statusCode, err = rh.RetrieveDevicesBase(ctx, w, encoder)
	default:
		return http.StatusBadRequest, fmt.Errorf("invalid value '%v' of '%v' query parameter", contentQuery, ContentQuery)
	}
	if err != nil {
		statusCode = kitNetHttp.ErrToStatusWithDef(err, statusCode)
		if statusCode == http.StatusNotFound {
			// return's empty array
			errEnc := encoder(w, []interface{}{}, http.StatusOK)
			if errEnc == nil {
				return http.StatusOK, nil
			}
		}
	}
	return statusCode, err

}

type callbackFunc func(ctx context.Context, w http.ResponseWriter, routeVars map[string]string, contentQuery string, encoder responseWriterEncoderFunc) (int, error)

func (rh *RequestHandler) retrieveWithCallback(w http.ResponseWriter, r *http.Request, callback callbackFunc) (int, error) {
	_, userID, err := parseAuth(rh.ownerClaim, r.Header.Get("Authorization"))
	if err != nil {
		return http.StatusUnauthorized, err
	}

	encoder, err := getResponseWriterEncoder(strings.Split(r.Header.Get("Accept"), ","))
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("cannot retrieve: %w", err)
	}

	return callback(kitNetGrpc.CtxWithOwner(r.Context(), userID), w, mux.Vars(r), getContentQueryValue(r.URL), encoder)
}

func (rh *RequestHandler) RetrieveDevices(w http.ResponseWriter, r *http.Request) {
	statusCode, err := rh.retrieveWithCallback(w, r, rh.RetrieveDevicesWithContentQuery)
	if err != nil {
		logAndWriteErrorResponse(fmt.Errorf("cannot retrieve all devices: %w", err), statusCode, w)
	}
}
