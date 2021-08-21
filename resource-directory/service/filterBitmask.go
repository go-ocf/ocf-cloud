package service

import (
	"time"

	"github.com/plgd-dev/cloud/grpc-gateway/pb"
)

type filterBitmask uint64

const (
	filterBitmaskResourceCreatePending       filterBitmask = 1
	filterBitmaskResourceCreated             filterBitmask = 1 << 1
	filterBitmaskResourceRetrievePending     filterBitmask = 1 << 2
	filterBitmaskResourceRetrieved           filterBitmask = 1 << 3
	filterBitmaskResourceUpdatePending       filterBitmask = 1 << 4
	filterBitmaskResourceUpdated             filterBitmask = 1 << 5
	filterBitmaskResourceDeletePending       filterBitmask = 1 << 6
	filterBitmaskResourceDeleted             filterBitmask = 1 << 7
	filterBitmaskDeviceMetadataUpdatePending filterBitmask = 1 << 8
	filterBitmaskDeviceMetadataUpdated       filterBitmask = 1 << 9
	filterBitmaskDeviceRegistered            filterBitmask = 1 << 10
	filterBitmaskDeviceUnregistered          filterBitmask = 1 << 11
	filterBitmaskResourceChanged             filterBitmask = 1 << 12
	filterBitmaskResourcesPublished          filterBitmask = 1 << 13
	filterBitmaskResourcesUnpublished        filterBitmask = 1 << 14
)

func filterPendingCommandToBitmask(f pb.GetPendingCommandsRequest_Command) filterBitmask {
	bitmask := filterBitmask(0)
	switch f {
	case pb.GetPendingCommandsRequest_RESOURCE_CREATE:
		bitmask |= filterBitmaskResourceCreatePending
	case pb.GetPendingCommandsRequest_RESOURCE_RETRIEVE:
		bitmask |= filterBitmaskResourceRetrievePending
	case pb.GetPendingCommandsRequest_RESOURCE_UPDATE:
		bitmask |= filterBitmaskResourceUpdatePending
	case pb.GetPendingCommandsRequest_RESOURCE_DELETE:
		bitmask |= filterBitmaskResourceDeletePending
	case pb.GetPendingCommandsRequest_DEVICE_METADATA_UPDATE:
		bitmask |= filterBitmaskDeviceMetadataUpdatePending
	}
	return bitmask
}

func filterPendingsCommandsToBitmask(commandFilter []pb.GetPendingCommandsRequest_Command) filterBitmask {
	bitmask := filterBitmask(0)
	if len(commandFilter) == 0 {
		bitmask = filterBitmaskResourceCreatePending | filterBitmaskResourceRetrievePending | filterBitmaskResourceUpdatePending | filterBitmaskResourceDeletePending | filterBitmaskDeviceMetadataUpdatePending
	} else {
		for _, f := range commandFilter {
			bitmask |= filterPendingCommandToBitmask(f)
		}
	}
	return bitmask
}

const devicesEventCount = pb.SubscribeToEvents_CreateSubscription_RESOURCE_CHANGED + 1

var eventToBitmask = [devicesEventCount]filterBitmask{
	pb.SubscribeToEvents_CreateSubscription_REGISTERED:                     filterBitmaskDeviceRegistered,
	pb.SubscribeToEvents_CreateSubscription_UNREGISTERED:                   filterBitmaskDeviceUnregistered,
	pb.SubscribeToEvents_CreateSubscription_DEVICE_METADATA_UPDATED:        filterBitmaskDeviceMetadataUpdated,
	pb.SubscribeToEvents_CreateSubscription_DEVICE_METADATA_UPDATE_PENDING: filterBitmaskDeviceMetadataUpdatePending,
	pb.SubscribeToEvents_CreateSubscription_RESOURCE_PUBLISHED:             filterBitmaskResourcesPublished,
	pb.SubscribeToEvents_CreateSubscription_RESOURCE_UNPUBLISHED:           filterBitmaskResourcesUnpublished,
	pb.SubscribeToEvents_CreateSubscription_RESOURCE_UPDATE_PENDING:        filterBitmaskResourceUpdatePending,
	pb.SubscribeToEvents_CreateSubscription_RESOURCE_UPDATED:               filterBitmaskResourceUpdated,
	pb.SubscribeToEvents_CreateSubscription_RESOURCE_RETRIEVE_PENDING:      filterBitmaskResourceRetrievePending,
	pb.SubscribeToEvents_CreateSubscription_RESOURCE_RETRIEVED:             filterBitmaskResourceRetrieved,
	pb.SubscribeToEvents_CreateSubscription_RESOURCE_DELETE_PENDING:        filterBitmaskResourceDeletePending,
	pb.SubscribeToEvents_CreateSubscription_RESOURCE_DELETED:               filterBitmaskResourceDeleted,
	pb.SubscribeToEvents_CreateSubscription_RESOURCE_CREATE_PENDING:        filterBitmaskResourceCreatePending,
	pb.SubscribeToEvents_CreateSubscription_RESOURCE_CREATED:               filterBitmaskResourceCreated,
	pb.SubscribeToEvents_CreateSubscription_RESOURCE_CHANGED:               filterBitmaskResourceChanged,
}

func devicesEventFilterToBitmask(f pb.SubscribeToEvents_CreateSubscription_Event) filterBitmask {
	return eventToBitmask[f]
}

func devicesEventsFilterToBitmask(commandFilter []pb.SubscribeToEvents_CreateSubscription_Event) filterBitmask {
	bitmask := filterBitmask(0)
	if len(commandFilter) == 0 {
		bitmask = filterBitmask(0xffffffff)
	} else {
		for _, f := range commandFilter {
			bitmask |= devicesEventFilterToBitmask(f)
		}
	}
	return bitmask
}

func toPendingCommands(resource *Resource, commandFilter filterBitmask, now time.Time) []*pb.PendingCommand {
	if resource.projection == nil {
		return nil
	}
	pendingCmds := make([]*pb.PendingCommand, 0, 32)
	if commandFilter&filterBitmaskResourceCreatePending > 0 {
		for _, pendingCmd := range resource.projection.resourceCreatePendings {
			if pendingCmd.IsExpired(now) {
				continue
			}
			pendingCmds = append(pendingCmds, &pb.PendingCommand{
				Command: &pb.PendingCommand_ResourceCreatePending{
					ResourceCreatePending: pendingCmd,
				},
			})
		}
	}
	if commandFilter&filterBitmaskResourceRetrievePending > 0 {
		for _, pendingCmd := range resource.projection.resourceRetrievePendings {
			if pendingCmd.IsExpired(now) {
				continue
			}
			pendingCmds = append(pendingCmds, &pb.PendingCommand{
				Command: &pb.PendingCommand_ResourceRetrievePending{
					ResourceRetrievePending: pendingCmd,
				},
			})
		}
	}
	if commandFilter&filterBitmaskResourceUpdatePending > 0 {
		for _, pendingCmd := range resource.projection.resourceUpdatePendings {
			if pendingCmd.IsExpired(now) {
				continue
			}
			pendingCmds = append(pendingCmds, &pb.PendingCommand{
				Command: &pb.PendingCommand_ResourceUpdatePending{
					ResourceUpdatePending: pendingCmd,
				},
			})
		}
	}
	if commandFilter&filterBitmaskResourceDeletePending > 0 {
		for _, pendingCmd := range resource.projection.resourceDeletePendings {
			if pendingCmd.IsExpired(now) {
				continue
			}
			pendingCmds = append(pendingCmds, &pb.PendingCommand{
				Command: &pb.PendingCommand_ResourceDeletePending{
					ResourceDeletePending: pendingCmd,
				},
			})
		}
	}
	return pendingCmds
}
