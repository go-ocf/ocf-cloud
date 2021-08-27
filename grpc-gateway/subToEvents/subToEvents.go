package subtoevents

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/plgd-dev/cloud/authorization/client"
	ownerEvents "github.com/plgd-dev/cloud/authorization/events"
	"github.com/plgd-dev/cloud/grpc-gateway/pb"
	"github.com/plgd-dev/cloud/pkg/log"
	"github.com/plgd-dev/cloud/pkg/net/grpc"
	"github.com/plgd-dev/cloud/resource-aggregate/commands"
	"github.com/plgd-dev/cloud/resource-aggregate/cqrs/eventbus"
	"github.com/plgd-dev/cloud/resource-aggregate/cqrs/eventbus/nats/subscriber"
	"github.com/plgd-dev/cloud/resource-aggregate/cqrs/utils"
	"github.com/plgd-dev/cloud/resource-aggregate/events"
	"github.com/plgd-dev/kit/strings"
	"google.golang.org/grpc/codes"
)

type SendEventFunc = func(e *pb.Event) error
type ErrFunc = func(err error)

type Sub struct {
	ctx                   context.Context
	filter                FilterBitmask
	send                  SendEventFunc
	req                   *pb.SubscribeToEvents_CreateSubscription
	id                    string
	correlationID         string
	errFunc               ErrFunc
	eventsSub             *subscriber.Subscriber
	ownerSubClose         func()
	devicesEventsObserver map[string]eventbus.Observer
	filteredDeviceIDs     strings.Set
	filteredResourceIDs   strings.Set

	sync.Mutex
}

func isFilteredDevice(filteredDeviceIDs strings.Set, deviceID string) bool {
	if len(filteredDeviceIDs) == 0 {
		return true
	}
	return filteredDeviceIDs.HasOneOf(deviceID)
}

func isFilteredResourceIDs(filteredResourceIDs strings.Set, resourceID string) bool {
	if len(filteredResourceIDs) == 0 {
		return true
	}
	return filteredResourceIDs.HasOneOf(resourceID)
}

var eventTypeToBitmaks = map[string]FilterBitmask{
	(&events.ResourceCreatePending{}).EventType():       FilterBitmaskResourceCreatePending,
	(&events.ResourceCreated{}).EventType():             FilterBitmaskResourceCreated,
	(&events.ResourceRetrievePending{}).EventType():     FilterBitmaskResourceRetrievePending,
	(&events.ResourceRetrieved{}).EventType():           FilterBitmaskResourceRetrieved,
	(&events.ResourceUpdatePending{}).EventType():       FilterBitmaskResourceUpdatePending,
	(&events.ResourceUpdated{}).EventType():             FilterBitmaskResourceUpdated,
	(&events.ResourceDeletePending{}).EventType():       FilterBitmaskResourceDeletePending,
	(&events.ResourceDeleted{}).EventType():             FilterBitmaskResourceDeleted,
	(&events.DeviceMetadataUpdatePending{}).EventType(): FilterBitmaskDeviceMetadataUpdatePending,
	(&events.DeviceMetadataUpdated{}).EventType():       FilterBitmaskDeviceMetadataUpdated,
	(&events.ResourceChanged{}).EventType():             FilterBitmaskResourceChanged,
	(&events.ResourceLinksPublished{}).EventType():      FilterBitmaskResourcesPublished,
	(&events.ResourceLinksUnpublished{}).EventType():    FilterBitmaskResourcesUnpublished,
}

func isFilteredEventype(filteredEventTypes FilterBitmask, eventType string) bool {
	if filteredEventTypes == 0 {
		return true
	}
	bit, ok := eventTypeToBitmaks[eventType]
	if !ok {
		return false
	}
	return filteredEventTypes&bit != 0
}

func (s *Sub) initDevice(deviceID string) error {
	if _, ok := s.devicesEventsObserver[deviceID]; ok {
		return nil
	}
	obs, err := s.eventsSub.Subscribe(s.ctx, s.id, utils.GetDeviceSubject(deviceID), s)
	if err != nil {
		return err
	}
	s.devicesEventsObserver[deviceID] = obs
	return nil
}

func (s *Sub) deinitDevice(deviceID string) error {
	devicesEventsObserver, ok := s.devicesEventsObserver[deviceID]
	if !ok {
		return nil
	}
	delete(s.devicesEventsObserver, deviceID)
	return devicesEventsObserver.Close()
}

type resourceEventHandler func(eventbus.EventUnmarshaler) (*pb.Event, error)

func handleResourcesPublished(eu eventbus.EventUnmarshaler) (*pb.Event, error) {
	var e events.ResourceLinksPublished
	if err := eu.Unmarshal(&e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event %v: %w", eu, err)
	}
	return &pb.Event{
		Type: &pb.Event_ResourcePublished{
			ResourcePublished: &e,
		},
	}, nil
}

func handleResourcesUnpublished(eu eventbus.EventUnmarshaler) (*pb.Event, error) {
	var e events.ResourceLinksUnpublished
	if err := eu.Unmarshal(&e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event %v: %w", eu, err)
	}
	return &pb.Event{
		Type: &pb.Event_ResourceUnpublished{
			ResourceUnpublished: &e,
		},
	}, nil
}

func handleResourceChanged(eu eventbus.EventUnmarshaler) (*pb.Event, error) {
	var e events.ResourceChanged
	if err := eu.Unmarshal(&e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event %v: %w", eu, err)
	}
	return &pb.Event{
		Type: &pb.Event_ResourceChanged{
			ResourceChanged: &e,
		},
	}, nil
}

func handleResourceUpdatePending(eu eventbus.EventUnmarshaler) (*pb.Event, error) {
	var e events.ResourceUpdatePending
	if err := eu.Unmarshal(&e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event %v: %w", eu, err)
	}
	return &pb.Event{
		Type: &pb.Event_ResourceUpdatePending{
			ResourceUpdatePending: &e,
		},
	}, nil
}

func handleResourceUpdated(eu eventbus.EventUnmarshaler) (*pb.Event, error) {
	var e events.ResourceUpdated
	if err := eu.Unmarshal(&e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event %v: %w", eu, err)
	}
	return &pb.Event{
		Type: &pb.Event_ResourceUpdated{
			ResourceUpdated: &e,
		},
	}, nil
}

func handleResourceRetrievePending(eu eventbus.EventUnmarshaler) (*pb.Event, error) {
	var e events.ResourceRetrievePending
	if err := eu.Unmarshal(&e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event %v: %w", eu, err)
	}
	return &pb.Event{
		Type: &pb.Event_ResourceRetrievePending{
			ResourceRetrievePending: &e,
		},
	}, nil
}

func handleResourceRetrieved(eu eventbus.EventUnmarshaler) (*pb.Event, error) {
	var e events.ResourceRetrieved
	if err := eu.Unmarshal(&e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event %v: %w", eu, err)
	}
	return &pb.Event{
		Type: &pb.Event_ResourceRetrieved{
			ResourceRetrieved: &e,
		},
	}, nil
}

func handleResourceDeletePending(eu eventbus.EventUnmarshaler) (*pb.Event, error) {
	var e events.ResourceDeletePending
	if err := eu.Unmarshal(&e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event %v: %w", eu, err)
	}
	return &pb.Event{
		Type: &pb.Event_ResourceDeletePending{
			ResourceDeletePending: &e,
		},
	}, nil
}

func handleResourceDeleted(eu eventbus.EventUnmarshaler) (*pb.Event, error) {
	var e events.ResourceDeleted
	if err := eu.Unmarshal(&e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event %v: %w", eu, err)
	}
	return &pb.Event{
		Type: &pb.Event_ResourceDeleted{
			ResourceDeleted: &e,
		},
	}, nil
}

func handleResourceCreatePending(eu eventbus.EventUnmarshaler) (*pb.Event, error) {
	var e events.ResourceCreatePending
	if err := eu.Unmarshal(&e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event %v: %w", eu, err)
	}
	return &pb.Event{
		Type: &pb.Event_ResourceCreatePending{
			ResourceCreatePending: &e,
		},
	}, nil
}

func handleResourceCreated(eu eventbus.EventUnmarshaler) (*pb.Event, error) {
	var e events.ResourceCreated
	if err := eu.Unmarshal(&e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event %v: %w", eu, err)
	}
	return &pb.Event{
		Type: &pb.Event_ResourceCreated{
			ResourceCreated: &e,
		},
	}, nil
}

func handleDeviceMetadataUpdatePending(eu eventbus.EventUnmarshaler) (*pb.Event, error) {
	var e events.DeviceMetadataUpdatePending
	if err := eu.Unmarshal(&e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event %v: %w", eu, err)
	}
	return &pb.Event{
		Type: &pb.Event_DeviceMetadataUpdatePending{
			DeviceMetadataUpdatePending: &e,
		},
	}, nil
}

func handleDeviceMetadataUpdated(eu eventbus.EventUnmarshaler) (*pb.Event, error) {
	var e events.DeviceMetadataUpdated
	if err := eu.Unmarshal(&e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event %v: %w", eu, err)
	}
	return &pb.Event{
		Type: &pb.Event_DeviceMetadataUpdated{
			DeviceMetadataUpdated: &e,
		},
	}, nil
}

var eventToHandler = map[string]resourceEventHandler{
	(&events.ResourceCreatePending{}).EventType():       handleResourceCreatePending,
	(&events.ResourceCreated{}).EventType():             handleResourceCreated,
	(&events.ResourceRetrievePending{}).EventType():     handleResourceRetrievePending,
	(&events.ResourceRetrieved{}).EventType():           handleResourceRetrieved,
	(&events.ResourceUpdatePending{}).EventType():       handleResourceUpdatePending,
	(&events.ResourceUpdated{}).EventType():             handleResourceUpdated,
	(&events.ResourceDeletePending{}).EventType():       handleResourceDeletePending,
	(&events.ResourceDeleted{}).EventType():             handleResourceDeleted,
	(&events.DeviceMetadataUpdatePending{}).EventType(): handleDeviceMetadataUpdatePending,
	(&events.DeviceMetadataUpdated{}).EventType():       handleDeviceMetadataUpdated,
	(&events.ResourceChanged{}).EventType():             handleResourceChanged,
	(&events.ResourceLinksPublished{}).EventType():      handleResourcesPublished,
	(&events.ResourceLinksUnpublished{}).EventType():    handleResourcesUnpublished,
}

func (s *Sub) handleEvent(eu eventbus.EventUnmarshaler) {
	log.Debug("handleEvent: EventType %v", eu)

	handler, ok := eventToHandler[eu.EventType()]
	if !ok {
		log.Errorf("unhandled event type %v", eu.EventType())
		return
	}

	ev, err := handler(eu)
	if ev != nil {
		log.Errorf("cannot get event: %w", err)
		return
	}
	ev.CorrelationId = s.correlationID
	ev.SubscriptionId = s.id
	err = s.send(ev)
	if err != nil {
		log.Errorf("cannot send event %v: %w", ev, err)
	}
}

func (s *Sub) Handle(ctx context.Context, iter eventbus.Iter) (err error) {
	for {
		eu, ok := iter.Next(ctx)
		if !ok {
			break
		}
		if !isFilteredResourceIDs(s.filteredResourceIDs, eu.AggregateID()) {
			continue
		}
		if !isFilteredEventype(s.filter, eu.EventType()) {
			continue
		}
		s.handleEvent(eu)
	}

	return iter.Err()
}

func (s *Sub) SetContext(ctx context.Context) {
	s.Lock()
	defer s.Unlock()
	s.ctx = ctx
}

func (s *Sub) Init(ownerCache *client.OwnerCache, sub *subscriber.Subscriber, ownerClaim string) error {
	owner, err := grpc.OwnerFromTokenMD(s.ctx, ownerClaim)
	if err != nil {
		return grpc.ForwardFromError(codes.InvalidArgument, err)
	}
	s.Lock()
	defer s.Unlock()
	close, err := ownerCache.Subscribe(owner, s.onOwnerEvent)
	if err != nil {
		return err
	}
	devices, err := ownerCache.GetDevices(s.ctx)
	if err != nil {
		close()
		return err
	}
	s.ownerSubClose = close
	devices = s.filterDevices(devices)
	var errors []error
	for _, device := range devices {
		err := s.initDevice(device)
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", errors)
	}
	return nil
}

func (s *Sub) filterDevices(devices []string) []string {
	filteredDevices := make([]string, 0, len(devices))
	for _, d := range devices {
		if isFilteredDevice(s.filteredDeviceIDs, d) {
			filteredDevices = append(devices, d)
		}
	}
	return filteredDevices
}

func (s *Sub) onRegisteredEvent(e *ownerEvents.DevicesRegistered) {
	devices := s.filterDevices(e.GetDeviceIds())
	if len(devices) == 0 {
		return
	}

	if s.filter&FilterBitmaskDeviceRegistered != 0 {
		err := s.send(&pb.Event{
			SubscriptionId: s.id,
			CorrelationId:  s.correlationID,
			Type: &pb.Event_DeviceRegistered_{
				DeviceRegistered: &pb.Event_DeviceRegistered{
					DeviceIds: devices,
				},
			},
		})
		if err != nil {
			s.errFunc(err)
		}
	}
	for _, deviceID := range devices {
		s.initDevice(deviceID)
	}
}

func (s *Sub) onUnregisteredEvent(e *ownerEvents.DevicesUnregistered) {
	devices := s.filterDevices(e.GetDeviceIds())
	if len(devices) == 0 {
		return
	}

	if s.filter&FilterBitmaskDeviceUnregistered != 0 {
		err := s.send(&pb.Event{
			SubscriptionId: s.id,
			CorrelationId:  s.correlationID,
			Type: &pb.Event_DeviceRegistered_{
				DeviceRegistered: &pb.Event_DeviceRegistered{
					DeviceIds: devices,
				},
			},
		})
		if err != nil {
			s.errFunc(err)
		}
	}
	for _, deviceID := range devices {
		s.deinitDevice(deviceID)
	}
}

func (s *Sub) onOwnerEvent(e *ownerEvents.Event) {
	s.Lock()
	defer s.Unlock()
	switch {
	case e.GetDevicesRegistered() != nil:
		s.onRegisteredEvent(e.GetDevicesRegistered())
	case e.GetDevicesUnregistered() != nil:
		s.onUnregisteredEvent(e.GetDevicesUnregistered())
	}
}

func New(ctx context.Context, eventsSub *subscriber.Subscriber, send SendEventFunc, correlationID string, req *pb.SubscribeToEvents_CreateSubscription) *Sub {
	filteredResourceIDs := strings.MakeSet()
	for _, r := range req.GetResourceIdFilter() {
		v := commands.ResourceIdFromString(r)
		if v != nil {
			filteredResourceIDs.Add(v.ToUUID())
		}
	}
	filteredDeviceIDs := strings.MakeSet(req.GetDeviceIdFilter()...)
	for _, r := range req.GetResourceIdFilter() {
		v := commands.ResourceIdFromString(r)
		if v != nil {
			filteredDeviceIDs.Add(v.GetDeviceId())
		}
	}
	return &Sub{
		ctx:                 ctx,
		filter:              EventsFilterToBitmask(req.GetEventFilter()),
		send:                send,
		req:                 req,
		id:                  uuid.NewString(),
		eventsSub:           eventsSub,
		filteredDeviceIDs:   strings.MakeSet(req.GetDeviceIdFilter()...),
		filteredResourceIDs: filteredResourceIDs,
	}
}
