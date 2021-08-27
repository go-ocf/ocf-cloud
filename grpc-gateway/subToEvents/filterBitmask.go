package subtoevents

import (
	"github.com/plgd-dev/cloud/grpc-gateway/pb"
)

type FilterBitmask uint64

const (
	FilterBitmaskResourceCreatePending       FilterBitmask = 1
	FilterBitmaskResourceCreated             FilterBitmask = 1 << 1
	FilterBitmaskResourceRetrievePending     FilterBitmask = 1 << 2
	FilterBitmaskResourceRetrieved           FilterBitmask = 1 << 3
	FilterBitmaskResourceUpdatePending       FilterBitmask = 1 << 4
	FilterBitmaskResourceUpdated             FilterBitmask = 1 << 5
	FilterBitmaskResourceDeletePending       FilterBitmask = 1 << 6
	FilterBitmaskResourceDeleted             FilterBitmask = 1 << 7
	FilterBitmaskDeviceMetadataUpdatePending FilterBitmask = 1 << 8
	FilterBitmaskDeviceMetadataUpdated       FilterBitmask = 1 << 9
	FilterBitmaskDeviceRegistered            FilterBitmask = 1 << 10
	FilterBitmaskDeviceUnregistered          FilterBitmask = 1 << 11
	FilterBitmaskResourceChanged             FilterBitmask = 1 << 12
	FilterBitmaskResourcesPublished          FilterBitmask = 1 << 13
	FilterBitmaskResourcesUnpublished        FilterBitmask = 1 << 14
)

func FilterPendingCommandToBitmask(f pb.GetPendingCommandsRequest_Command) FilterBitmask {
	bitmask := FilterBitmask(0)
	switch f {
	case pb.GetPendingCommandsRequest_RESOURCE_CREATE:
		bitmask |= FilterBitmaskResourceCreatePending
	case pb.GetPendingCommandsRequest_RESOURCE_RETRIEVE:
		bitmask |= FilterBitmaskResourceRetrievePending
	case pb.GetPendingCommandsRequest_RESOURCE_UPDATE:
		bitmask |= FilterBitmaskResourceUpdatePending
	case pb.GetPendingCommandsRequest_RESOURCE_DELETE:
		bitmask |= FilterBitmaskResourceDeletePending
	case pb.GetPendingCommandsRequest_DEVICE_METADATA_UPDATE:
		bitmask |= FilterBitmaskDeviceMetadataUpdatePending
	}
	return bitmask
}

func FilterPendingsCommandsToBitmask(commandFilter []pb.GetPendingCommandsRequest_Command) FilterBitmask {
	bitmask := FilterBitmask(0)
	if len(commandFilter) == 0 {
		bitmask = FilterBitmaskResourceCreatePending | FilterBitmaskResourceRetrievePending | FilterBitmaskResourceUpdatePending | FilterBitmaskResourceDeletePending | FilterBitmaskDeviceMetadataUpdatePending
	} else {
		for _, f := range commandFilter {
			bitmask |= FilterPendingCommandToBitmask(f)
		}
	}
	return bitmask
}

func EventFilterToBitmask(f pb.SubscribeToEvents_CreateSubscription_Event) FilterBitmask {
	bitmask := FilterBitmask(0)
	switch f {
	case pb.SubscribeToEvents_CreateSubscription_RESOURCE_CREATE_PENDING:
		bitmask |= FilterBitmaskResourceCreatePending
	case pb.SubscribeToEvents_CreateSubscription_RESOURCE_CREATED:
		bitmask |= FilterBitmaskResourceCreated
	case pb.SubscribeToEvents_CreateSubscription_RESOURCE_RETRIEVE_PENDING:
		bitmask |= FilterBitmaskResourceRetrievePending
	case pb.SubscribeToEvents_CreateSubscription_RESOURCE_RETRIEVED:
		bitmask |= FilterBitmaskResourceRetrieved
	case pb.SubscribeToEvents_CreateSubscription_RESOURCE_UPDATE_PENDING:
		bitmask |= FilterBitmaskResourceUpdatePending
	case pb.SubscribeToEvents_CreateSubscription_RESOURCE_UPDATED:
		bitmask |= FilterBitmaskResourceUpdated
	case pb.SubscribeToEvents_CreateSubscription_RESOURCE_DELETE_PENDING:
		bitmask |= FilterBitmaskResourceDeletePending
	case pb.SubscribeToEvents_CreateSubscription_RESOURCE_DELETED:
		bitmask |= FilterBitmaskResourceDeleted
	case pb.SubscribeToEvents_CreateSubscription_DEVICE_METADATA_UPDATE_PENDING:
		bitmask |= FilterBitmaskDeviceMetadataUpdatePending
	case pb.SubscribeToEvents_CreateSubscription_DEVICE_METADATA_UPDATED:
		bitmask |= FilterBitmaskDeviceMetadataUpdated
	case pb.SubscribeToEvents_CreateSubscription_REGISTERED:
		bitmask |= FilterBitmaskDeviceRegistered
	case pb.SubscribeToEvents_CreateSubscription_UNREGISTERED:
		bitmask |= FilterBitmaskDeviceUnregistered
	case pb.SubscribeToEvents_CreateSubscription_RESOURCE_PUBLISHED:
		bitmask |= FilterBitmaskResourcesPublished
	case pb.SubscribeToEvents_CreateSubscription_RESOURCE_UNPUBLISHED:
		bitmask |= FilterBitmaskResourcesUnpublished
	case pb.SubscribeToEvents_CreateSubscription_RESOURCE_CHANGED:
		bitmask |= FilterBitmaskResourceChanged
	}
	return bitmask
}

func EventsFilterToBitmask(commandFilter []pb.SubscribeToEvents_CreateSubscription_Event) FilterBitmask {
	bitmask := FilterBitmask(0)
	if len(commandFilter) == 0 {
		bitmask = FilterBitmask(0xffffffff)
	} else {
		for _, f := range commandFilter {
			bitmask |= EventFilterToBitmask(f)
		}
	}
	return bitmask
}
