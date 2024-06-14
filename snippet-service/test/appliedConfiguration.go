package test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/plgd-dev/hub/v2/snippet-service/pb"
	"github.com/plgd-dev/hub/v2/snippet-service/store"
	hubTest "github.com/plgd-dev/hub/v2/test"
	pbTest "github.com/plgd-dev/hub/v2/test/pb"
	"github.com/stretchr/testify/require"
)

func DeviceID(i int) string {
	return "device" + strconv.Itoa(i)
}

func AppliedConfigurationID(i int) string {
	if id, ok := RuntimeConfig.appliedConfigurationIds[i]; ok {
		return id
	}
	id := uuid.NewString()
	RuntimeConfig.appliedConfigurationIds[i] = id
	return id
}

func SetAppliedConfigurationExecutedBy(ac *store.AppliedDeviceConfiguration, i int) {
	if i%RuntimeConfig.NumConfigurations == 0 {
		ac.ExecutedBy = pb.MakeExecutedByOnDemand()
		return
	}
	ac.ExecutedBy = pb.MakeExecutedByConditionId(ConditionID(i), uint64(i%RuntimeConfig.NumConditions))
}

func AppliedConfigurationResource(t *testing.T, deviceID string, start, n int) []*pb.AppliedDeviceConfiguration_Resource {
	resources := make([]*pb.AppliedDeviceConfiguration_Resource, 0, n)
	for i := start; i < start+n; i++ {
		correlationID := "corID" + strconv.Itoa(i)
		resource := &pb.AppliedDeviceConfiguration_Resource{
			Href:          hubTest.TestResourceLightInstanceHref(strconv.Itoa(i)),
			CorrelationId: correlationID,
			Status:        pb.AppliedDeviceConfiguration_Resource_Status(i % 4),
		}
		if resource.GetStatus() == pb.AppliedDeviceConfiguration_Resource_DONE {
			resource.ResourceUpdated = pbTest.MakeResourceUpdated(t,
				deviceID,
				resource.GetHref(),
				hubTest.TestResourceLightInstanceResourceTypes,
				correlationID,
				map[string]interface{}{
					"power": i,
				},
			)
		}
		resources = append(resources, resource)
	}
	return resources
}

func getAppliedConditions(t *testing.T) map[string]*store.AppliedDeviceConfiguration {
	owners := make(map[int]string, RuntimeConfig.NumConfigurations)
	acs := make(map[string]*store.AppliedDeviceConfiguration)
	i := 0
	for d := range RuntimeConfig.numDevices {
		for c := range RuntimeConfig.NumConfigurations {
			owner, ok := owners[i%RuntimeConfig.NumConfigurations]
			if !ok {
				owner = Owner(i % RuntimeConfig.numOwners)
				owners[i%RuntimeConfig.NumConfigurations] = owner
			}
			deviceID := DeviceID(d)
			ac := &store.AppliedDeviceConfiguration{
				Id:       AppliedConfigurationID(i),
				DeviceId: deviceID,
				Owner:    owner,
				ConfigurationId: &pb.AppliedDeviceConfiguration_RelationTo{
					Id:      ConfigurationID(c),
					Version: uint64(i % RuntimeConfig.NumConfigurations),
				},
				Resources: AppliedConfigurationResource(t, deviceID, i%16, (i%5)+1),
				Timestamp: time.Now().UnixNano(),
			}
			SetAppliedConfigurationExecutedBy(ac, i)
			acs[ac.GetId()] = ac
			i++
		}
	}
	return acs
}

func AddAppliedConfigurationsToStore(ctx context.Context, t *testing.T, s store.Store) map[string]*store.AppliedDeviceConfiguration {
	acs := getAppliedConditions(t)
	acsToInsert := make([]*store.AppliedDeviceConfiguration, 0, len(acs))
	for _, c := range acs {
		acsToInsert = append(acsToInsert, c)
	}
	err := s.InsertAppliedConfigurations(ctx, acsToInsert...)
	require.NoError(t, err)
	return acs
}
