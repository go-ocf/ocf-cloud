package subscription

import (
	"context"
	"fmt"
	"io"
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
	"go.uber.org/atomic"
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
	resourceDirectory     pb.GrpcGatewayClient
	closed                atomic.Bool
	deduplicateInitEvents map[string]uint64

	eventsCache   chan eventbus.EventUnmarshaler
	ownerCache    chan *ownerEvents.Event
	done          chan struct{}
	mutex         sync.Mutex
	eventsCacheWg sync.WaitGroup
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

func isFilteredBit(filteredEventTypes FilterBitmask, bit FilterBitmask) bool {
	return filteredEventTypes&bit != 0
}

func isFilteredEventype(filteredEventTypes FilterBitmask, eventType string) bool {
	bit, ok := eventTypeToBitmaks[eventType]
	if !ok {
		return false
	}
	return isFilteredBit(filteredEventTypes, bit)
}

func (s *Sub) deinitDeviceLocked(deviceID string) error {
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
	if !s.isFiltered(eu) {
		return
	}
	handler, ok := eventToHandler[eu.EventType()]
	if !ok {
		log.Errorf("unhandled event type %v", eu.EventType())
		return
	}

	ev, err := handler(eu)
	if err != nil {
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

func (s *Sub) isFiltered(ev event) bool {
	if s.closed.Load() {
		return false
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return !s.isDuplicatedInitEventLocked(ev)
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

		select {
		case <-s.done:
			return nil
		case s.eventsCache <- eu:
		}
	}

	return iter.Err()
}

func (s *Sub) SetContext(ctx context.Context) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.ctx = ctx
}

func (s *Sub) Id() string {
	return s.id
}

func (s *Sub) subscribeToOwnerEvents(ownerCache *client.OwnerCache) error {
	owner, err := grpc.OwnerFromTokenMD(s.ctx, ownerCache.OwnerClaim())
	if err != nil {
		return grpc.ForwardFromError(codes.InvalidArgument, err)
	}
	close, err := ownerCache.Subscribe(owner, s.onOwnerEvent)
	if err != nil {
		return err
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.ownerSubClose = close
	return nil
}

func (s *Sub) initOwnerSubscription(ownerCache *client.OwnerCache) ([]string, error) {
	err := s.subscribeToOwnerEvents(ownerCache)
	if err != nil {
		_ = s.Close()
		return nil, err
	}
	devices, err := ownerCache.GetDevices(s.ctx)
	if err != nil {
		_ = s.Close()
		return nil, err
	}

	s.mutex.Lock()
	err = s.initEventSubscriptionsLocked(devices)
	s.mutex.Unlock()
	if err != nil {
		_ = s.Close()
		return nil, err
	}
	return devices, nil
}

func (s *Sub) start() {
	s.eventsCacheWg.Add(1)
	go func() {
		defer s.eventsCacheWg.Done()
		s.run()
	}()
}

func (s *Sub) Init(ownerCache *client.OwnerCache) error {
	devices, err := s.initOwnerSubscription(ownerCache)
	if err != nil {
		return err
	}
	if !s.req.GetIncludeCurrentState() {
		s.start()
		return nil
	}
	var initEventFuncs = []func([]string) error{
		s.sendDevicesRegistered,
		s.initDeviceMetadataUpdated,
		s.initResourcesPublished,
		s.initResourceChanged,
		s.initPendingCommands,
	}
	var errors []error
	for _, f := range initEventFuncs {
		err := f(devices)
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		_ = s.Close()
		return fmt.Errorf("%v", errors)
	}
	s.start()
	return nil
}

func (s *Sub) filterDevices(devices []string) []string {
	filteredDevices := make([]string, 0, len(devices))
	for _, d := range devices {
		if isFilteredDevice(s.filteredDeviceIDs, d) {
			filteredDevices = append(filteredDevices, d)
		}
	}
	return filteredDevices
}

func (s *Sub) initEventSubscriptionsLocked(deviceIDs []string) error {
	var errors []error
	for _, deviceID := range deviceIDs {
		if _, ok := s.devicesEventsObserver[deviceID]; ok {
			continue
		}
		obs, err := s.eventsSub.Subscribe(s.ctx, deviceID+"."+s.id, utils.GetDeviceSubject(deviceID), s)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		s.devicesEventsObserver[deviceID] = obs
	}
	if len(errors) > 0 {
		return fmt.Errorf("cannot init events subscription for devices[%v]: %v", deviceIDs, errors)
	}
	return nil
}

func (s *Sub) sendDevicesRegistered(deviceIDs []string) error {
	if !isFilteredBit(s.filter, FilterBitmaskDeviceRegistered) {
		return nil
	}
	err := s.send(&pb.Event{
		SubscriptionId: s.id,
		CorrelationId:  s.correlationID,
		Type: &pb.Event_DeviceRegistered_{
			DeviceRegistered: &pb.Event_DeviceRegistered{
				DeviceIds: deviceIDs,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("cannot send devices registered for devices %v: %w", deviceIDs, err)
	}
	return nil
}

func (s *Sub) initResourceChanged(deviceIDs []string) error {
	if !isFilteredBit(s.filter, FilterBitmaskResourceChanged) {
		return nil
	}
	errFunc := func(err error) error {
		return fmt.Errorf("cannot init resources changed events for devices %v: %w", deviceIDs, err)
	}
	deviceIdFilter := deviceIDs
	if len(s.req.GetResourceIdFilter()) > 0 {
		deviceIdFilter = nil
	}
	resourcesClient, err := s.resourceDirectory.GetResources(s.ctx, &pb.GetResourcesRequest{
		DeviceIdFilter:   deviceIdFilter,
		ResourceIdFilter: s.req.GetResourceIdFilter(),
	})
	if err != nil {
		return errFunc(fmt.Errorf("cannot get resources: %w", err))
	}
	for {
		recv, err := resourcesClient.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return errFunc(fmt.Errorf("cannot receive resource: %w", err))
		}
		if recv.GetData() == nil {
			// event doesn't contains data - resource is not initialized yet
			continue
		}
		ev := &pb.Event{
			SubscriptionId: s.id,
			CorrelationId:  s.correlationID,
			Type: &pb.Event_ResourceChanged{
				ResourceChanged: recv.GetData(),
			},
		}
		s.fillDeduplicateInitEvent(ev.GetResourceChanged())
		err = s.send(ev)
		if err != nil {
			return errFunc(fmt.Errorf("cannot send a resource: %w", err))
		}
	}
}

func (s *Sub) initDeviceMetadataUpdated(deviceIDs []string) error {
	if !isFilteredBit(s.filter, FilterBitmaskDeviceMetadataUpdated) {
		return nil
	}
	errFunc := func(err error) error {
		return fmt.Errorf("cannot init devices metadata for devices %v: %w", deviceIDs, err)
	}
	linksClient, err := s.resourceDirectory.GetDevicesMetadata(s.ctx, &pb.GetDevicesMetadataRequest{
		DeviceIdFilter: deviceIDs,
	})
	if err != nil {
		return errFunc(fmt.Errorf("cannot get devices metadata: %w", err))
	}
	for {
		recv, err := linksClient.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return errFunc(fmt.Errorf("cannot receive devices metadata: %w", err))
		}
		ev := &pb.Event{
			SubscriptionId: s.id,
			CorrelationId:  s.correlationID,
			Type: &pb.Event_DeviceMetadataUpdated{
				DeviceMetadataUpdated: recv,
			},
		}
		s.fillDeduplicateInitEvent(ev.GetDeviceMetadataUpdated())
		err = s.send(ev)
		if err != nil {
			return errFunc(fmt.Errorf("cannot send a devices metadata: %w", err))
		}
	}
}

func (s *Sub) initResourcesPublished(deviceIDs []string) error {
	if !isFilteredBit(s.filter, FilterBitmaskResourcesPublished) {
		return nil
	}
	errFunc := func(err error) error {
		return fmt.Errorf("cannot init resources published events for devices %v: %w", deviceIDs, err)
	}
	linksClient, err := s.resourceDirectory.GetResourceLinks(s.ctx, &pb.GetResourceLinksRequest{
		DeviceIdFilter: deviceIDs,
	})
	if err != nil {
		return errFunc(fmt.Errorf("cannot get resource links: %w", err))
	}
	for {
		recv, err := linksClient.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return errFunc(fmt.Errorf("cannot receive resource links: %w", err))
		}
		ev := &pb.Event{
			SubscriptionId: s.id,
			CorrelationId:  s.correlationID,
			Type: &pb.Event_ResourcePublished{
				ResourcePublished: recv,
			},
		}
		s.fillDeduplicateInitEvent(ev.GetResourcePublished())
		err = s.send(ev)
		if err != nil {
			return errFunc(fmt.Errorf("cannot send a resource links: %w", err))
		}
	}
}

func pendingCommandToEvent(cmd *pb.PendingCommand) (*pb.Event, event) {
	switch c := cmd.GetCommand().(type) {
	case *pb.PendingCommand_DeviceMetadataUpdatePending:
		return &pb.Event{
			Type: &pb.Event_DeviceMetadataUpdatePending{
				DeviceMetadataUpdatePending: c.DeviceMetadataUpdatePending,
			},
		}, c.DeviceMetadataUpdatePending
	case *pb.PendingCommand_ResourceCreatePending:
		return &pb.Event{
			Type: &pb.Event_ResourceCreatePending{
				ResourceCreatePending: c.ResourceCreatePending,
			},
		}, c.ResourceCreatePending
	case *pb.PendingCommand_ResourceDeletePending:
		return &pb.Event{
			Type: &pb.Event_ResourceDeletePending{
				ResourceDeletePending: c.ResourceDeletePending,
			},
		}, c.ResourceDeletePending
	case *pb.PendingCommand_ResourceRetrievePending:
		return &pb.Event{
			Type: &pb.Event_ResourceRetrievePending{
				ResourceRetrievePending: c.ResourceRetrievePending,
			},
		}, c.ResourceRetrievePending
	case *pb.PendingCommand_ResourceUpdatePending:
		return &pb.Event{
			Type: &pb.Event_ResourceUpdatePending{
				ResourceUpdatePending: c.ResourceUpdatePending,
			},
		}, c.ResourceUpdatePending
	}
	return nil, nil
}

type event = interface {
	EventType() string
	AggregateID() string
	Version() uint64
}

func deduplicateEventKey(ev event) string {
	return ev.AggregateID() + ev.EventType()
}

func (s *Sub) isDuplicatedInitEventLocked(ev event) bool {
	key := deduplicateEventKey(ev)
	version, ok := s.deduplicateInitEvents[key]
	if !ok {
		return false
	}
	if version >= ev.Version() {
		return true
	}
	return false
}

func (s *Sub) fillDeduplicateInitEvent(v event) {
	key := deduplicateEventKey(v)
	version, ok := s.deduplicateInitEvents[key]
	if !ok || version < v.Version() {
		s.deduplicateInitEvents[key] = v.Version()
	}
}

func (s *Sub) initPendingCommands(deviceIDs []string) error {
	if !isFilteredBit(s.filter,
		FilterBitmaskDeviceMetadataUpdatePending|
			FilterBitmaskResourceCreatePending|
			FilterBitmaskResourceRetrievePending|
			FilterBitmaskResourceUpdatePending|
			FilterBitmaskResourceDeletePending) {
		return nil
	}
	errFunc := func(err error) error {
		return fmt.Errorf("cannot init pending commands for devices %v: %w", deviceIDs, err)
	}

	deviceIdFilter := deviceIDs
	if len(s.req.GetResourceIdFilter()) > 0 {
		deviceIdFilter = nil
	}

	pendingCommands, err := s.resourceDirectory.GetPendingCommands(s.ctx, &pb.GetPendingCommandsRequest{
		DeviceIdFilter:   deviceIdFilter,
		ResourceIdFilter: s.req.GetResourceIdFilter(),
		CommandFilter:    BitmaskToFilterPendingsCommands(s.filter),
	})
	if err != nil {
		return errFunc(fmt.Errorf("cannot get pending commands: %w", err))
	}
	for {
		recv, err := pendingCommands.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return errFunc(fmt.Errorf("cannot receive pending command: %w", err))
		}
		ev, deduplicateEvent := pendingCommandToEvent(recv)
		if ev == nil {
			s.errFunc(errFunc(fmt.Errorf("unknown recv command %T", recv.GetCommand())))
			continue
		}
		ev.CorrelationId = s.correlationID
		ev.SubscriptionId = s.id

		s.fillDeduplicateInitEvent(deduplicateEvent)

		err = s.send(ev)
		if err != nil {
			return errFunc(fmt.Errorf("cannot send a pending command: %w", err))
		}
	}
}

func (s *Sub) onRegisteredEventLocked(e *ownerEvents.DevicesRegistered) {
	devices := s.filterDevices(e.GetDeviceIds())
	if len(devices) == 0 {
		return
	}
	err := s.initEventSubscriptionsLocked(devices)
	if err != nil {
		s.errFunc(err)
		return
	}
	err = s.sendDevicesRegistered(devices)
	if err != nil {
		s.errFunc(err)
	}
}

func (s *Sub) onUnregisteredEventLocked(e *ownerEvents.DevicesUnregistered) {
	devices := s.filterDevices(e.GetDeviceIds())
	if len(devices) == 0 {
		return
	}
	for _, deviceID := range devices {
		err := s.deinitDeviceLocked(deviceID)
		if err != nil {
			s.errFunc(fmt.Errorf("cannot deinit device %v: %w", deviceID, err))
		}
	}
	if isFilteredBit(s.filter, s.filter&FilterBitmaskDeviceUnregistered) {
		err := s.send(&pb.Event{
			SubscriptionId: s.id,
			CorrelationId:  s.correlationID,
			Type: &pb.Event_DeviceUnregistered_{
				DeviceUnregistered: &pb.Event_DeviceUnregistered{
					DeviceIds: devices,
				},
			},
		})
		if err != nil {
			s.errFunc(fmt.Errorf("cannot send device unregistered event for devices %v: %w", devices, err))
		}
	}
}

func (s *Sub) onOwnerEvent(e *ownerEvents.Event) {
	select {
	case <-s.done:
	case s.ownerCache <- e:
	}
}

func (s *Sub) handleOwnerEvent(e *ownerEvents.Event) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.closed.Load() {
		return
	}
	switch {
	case e.GetDevicesRegistered() != nil:
		s.onRegisteredEventLocked(e.GetDevicesRegistered())
	case e.GetDevicesUnregistered() != nil:
		s.onUnregisteredEventLocked(e.GetDevicesUnregistered())
	}
}

func (s *Sub) setClosed() bool {
	return s.closed.CAS(false, true)
}

func (s *Sub) pullOutDevicesEventsObserver() map[string]eventbus.Observer {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	devicesEventsObserver := s.devicesEventsObserver
	s.devicesEventsObserver = make(map[string]eventbus.Observer)
	return devicesEventsObserver
}

func cleanUpDevicesEventsObservers(devicesEventsObserver map[string]eventbus.Observer) error {
	var errors []error
	for _, obs := range devicesEventsObserver {
		err := obs.Close()
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", errors)
	}
	return nil
}

func (s *Sub) cleanUp(devicesEventsObserver map[string]eventbus.Observer) error {
	close(s.done)
	s.eventsCacheWg.Wait()
	if s.ownerSubClose != nil {
		s.ownerSubClose()
	}
	return cleanUpDevicesEventsObservers(devicesEventsObserver)
}

func (s *Sub) Close() error {
	if !s.setClosed() {
		return nil
	}
	return s.cleanUp(s.pullOutDevicesEventsObserver())
}

func (s *Sub) CleanUpDeduplicationInitEventsCache() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.deduplicateInitEvents = make(map[string]uint64)
}

func (s *Sub) run() {
	for {
		select {
		case <-s.done:
			return
		case event := <-s.eventsCache:
			s.handleEvent(event)
		case ownerEvent := <-s.ownerCache:
			s.handleOwnerEvent(ownerEvent)
		}
	}
}

func New(ctx context.Context, eventsSub *subscriber.Subscriber, resourceDirectory pb.GrpcGatewayClient, send SendEventFunc, correlationID string, cacheSize int, errFunc ErrFunc, req *pb.SubscribeToEvents_CreateSubscription) *Sub {
	bitmask := EventsFilterToBitmask(req.GetEventFilter())
	filteredResourceIDs := strings.MakeSet()
	filteredDeviceIDs := strings.MakeSet(req.GetDeviceIdFilter()...)
	for _, r := range req.GetResourceIdFilter() {
		v := commands.ResourceIdFromString(r)
		if v == nil {
			continue
		}
		filteredResourceIDs.Add(v.ToUUID())
		filteredDeviceIDs.Add(v.GetDeviceId())
		if len(req.GetEventFilter()) > 0 {
			if bitmask&(FilterBitmaskDeviceMetadataUpdatePending|FilterBitmaskDeviceMetadataUpdated) != 0 {
				filteredResourceIDs.Add(commands.MakeStatusResourceUUID(v.GetDeviceId()))
			}
			if bitmask&(FilterBitmaskResourcesPublished|FilterBitmaskResourcesUnpublished) != 0 {
				filteredResourceIDs.Add(commands.MakeLinksResourceUUID(v.GetDeviceId()))
			}
		}
	}
	id := uuid.NewString()
	if errFunc == nil {
		errFunc = func(err error) {}
	} else {
		newErrFunc := func(err error) {
			errFunc(fmt.Errorf("correlationId: %v, subscriptionId: %v: %w", correlationID, id, err))
		}
		errFunc = newErrFunc
	}
	return &Sub{
		ctx:                   ctx,
		filter:                EventsFilterToBitmask(req.GetEventFilter()),
		send:                  send,
		req:                   req,
		id:                    id,
		eventsSub:             eventsSub,
		filteredDeviceIDs:     strings.MakeSet(req.GetDeviceIdFilter()...),
		filteredResourceIDs:   filteredResourceIDs,
		resourceDirectory:     resourceDirectory,
		errFunc:               errFunc,
		correlationID:         correlationID,
		devicesEventsObserver: make(map[string]eventbus.Observer),
		deduplicateInitEvents: make(map[string]uint64),
		eventsCache:           make(chan eventbus.EventUnmarshaler, cacheSize),
		ownerCache:            make(chan *ownerEvents.Event, cacheSize),
		done:                  make(chan struct{}),
	}
}
