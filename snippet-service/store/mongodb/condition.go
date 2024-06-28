package mongodb

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
	"github.com/plgd-dev/hub/v2/pkg/mongodb"
	"github.com/plgd-dev/hub/v2/pkg/strings"
	"github.com/plgd-dev/hub/v2/snippet-service/pb"
	"github.com/plgd-dev/hub/v2/snippet-service/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Store) InsertConditions(ctx context.Context, conds ...*store.Condition) error {
	documents := make([]interface{}, 0, len(conds))
	for _, cond := range conds {
		documents = append(documents, cond)
	}
	_, err := s.Collection(conditionsCol).InsertMany(ctx, documents)
	return err
}

func (s *Store) CreateCondition(ctx context.Context, cond *pb.Condition) (*pb.Condition, error) {
	newCond, err := store.ValidateAndNormalizeCondition(cond, false)
	if err != nil {
		return nil, err
	}
	if newCond.GetId() == "" {
		newCond.Id = uuid.NewString()
	}
	newCond.Timestamp = time.Now().UnixNano()
	storeCond := store.MakeFirstCondition(newCond)
	_, err = s.Collection(conditionsCol).InsertOne(ctx, storeCond)
	if err != nil {
		return nil, err
	}
	return storeCond.GetLatest()
}

func filterCondition(cond *pb.Condition) bson.M {
	filter := bson.M{
		store.IDKey:    cond.GetId(),
		store.OwnerKey: cond.GetOwner(),
	}
	if cond.GetConfigurationId() != "" {
		filter[store.ConfigurationIDKey] = cond.GetConfigurationId()
	}
	if cond.GetVersion() != 0 {
		// if is set -> it must be higher than the $latest.version
		filter[store.LatestKey+"."+store.VersionKey] = bson.M{"$lt": cond.GetVersion()}
	}
	return filter
}

func updateLatestCondition(cond *pb.Condition) (bson.M, bson.A) {
	ts := time.Now().UnixNano()
	unsetLatest := bson.A{}
	latest := bson.M{
		store.VersionKey:            cond.GetVersion(),
		store.EnabledKey:            cond.GetEnabled(),
		store.TimestampKey:          ts,
		store.JqExpressionFilterKey: cond.GetJqExpressionFilter(),
		store.ApiAccessTokenKey:     cond.GetApiAccessToken(),
	}
	if len(cond.GetDeviceIdFilter()) == 0 {
		unsetLatest = append(unsetLatest, store.DeviceIDFilterKey)
	} else {
		latest[store.DeviceIDFilterKey] = cond.GetDeviceIdFilter()
	}
	if len(cond.GetResourceTypeFilter()) == 0 {
		unsetLatest = append(unsetLatest, store.ResourceTypeFilterKey)
	} else {
		latest[store.ResourceTypeFilterKey] = cond.GetResourceTypeFilter()
	}
	if len(cond.GetResourceHrefFilter()) == 0 {
		unsetLatest = append(unsetLatest, store.ResourceHrefFilterKey)
	} else {
		latest[store.ResourceHrefFilterKey] = cond.GetResourceHrefFilter()
	}
	if cond.GetName() != "" {
		latest[store.NameKey] = cond.GetName()
	}
	if cond.GetVersion() == 0 {
		// if version is not set -> set it to $latest.version + 1
		latest[store.VersionKey] = incrementLatestVersion()
	}
	return latest, unsetLatest
}

func updateCondition(cond *pb.Condition) mongo.Pipeline {
	latest, unsetLatest := updateLatestCondition(cond)
	keys := make([]string, 0, len(latest))
	for k := range latest {
		keys = append(keys, k)
	}
	setVersions := appendLatestToVersions(keys)
	p := mongo.Pipeline{
		bson.D{{Key: mongodb.Set, Value: bson.M{
			store.LatestKey: latest,
		}}},
		bson.D{{Key: mongodb.Set, Value: bson.M{
			store.VersionsKey: setVersions,
		}}},
	}
	if len(unsetLatest) > 0 {
		p = append(p, bson.D{{Key: mongodb.Unset, Value: unsetLatest}})
	}
	return p
}

func (s *Store) UpdateCondition(ctx context.Context, cond *pb.Condition) (*pb.Condition, error) {
	newCond, err := store.ValidateAndNormalizeCondition(cond, true)
	if err != nil {
		return nil, err
	}

	filter := filterCondition(newCond)
	update := updateCondition(newCond)
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetProjection(bson.M{store.VersionsKey: false})
	result := s.Collection(conditionsCol).FindOneAndUpdate(ctx, filter, update, opts)
	if result.Err() != nil {
		return nil, result.Err()
	}

	updatedCond := &store.Condition{}
	err = result.Decode(&updatedCond)
	if err != nil {
		return nil, err
	}
	return updatedCond.GetLatest()
}

// getConditionsByID returns full condition documents matched by ID
func (s *Store) getConditionsByID(ctx context.Context, owner string, ids []string, p store.ProcessConditions) error {
	cur, err := s.Collection(conditionsCol).Find(ctx, toIdFilterQuery(owner, toIdQuery(ids), false))
	if err != nil {
		return err
	}
	return processCursor(ctx, cur, p)
}

func toIdOrConfIdQuery(ids, confIds []string) bson.M {
	filter := make([]bson.M, 0, 2)
	idFilter := inArrayQuery(store.IDKey, ids)
	if len(idFilter) > 0 {
		filter = append(filter, idFilter)
	}
	confIdFilter := inArrayQuery(store.ConfigurationIDKey, confIds)
	if len(confIdFilter) > 0 {
		filter = append(filter, confIdFilter)
	}
	if len(filter) == 0 {
		return nil
	}
	if len(filter) == 1 {
		return filter[0]
	}
	return bson.M{mongodb.Or: filter}
}

// getLatestConditions returns the latest condition from document matched by condition ID or configuration ID
func (s *Store) getLatestConditions(ctx context.Context, owner string, ids, confIds []string, p store.ProcessConditions) error {
	opt := options.Find().SetProjection(bson.M{store.VersionsKey: false})
	cur, err := s.Collection(conditionsCol).Find(ctx, toIdFilterQuery(owner, toIdOrConfIdQuery(ids, confIds), false), opt)
	if err != nil {
		return err
	}
	return processCursor(ctx, cur, p)
}

// getConditionsByAggregation returns conditions matched by ID and versions
func (s *Store) getConditionsByAggregation(ctx context.Context, owner, id string, versions []uint64, p store.ProcessConditions) error {
	cur, err := s.Collection(conditionsCol).Aggregate(ctx, getPipeline(owner, id, versions))
	if err != nil {
		return err
	}
	return processCursor(ctx, cur, p)
}

func (s *Store) GetConditions(ctx context.Context, owner string, query *pb.GetConditionsRequest, p store.Process[store.Condition]) error {
	vf := pb.PartitionIDFilter(query.GetIdFilter())
	confIdLatestFilter := strings.Unique(query.GetConfigurationIdFilter())
	var errors *multierror.Error
	if len(vf.All) > 0 || vf.IsEmpty() && len(confIdLatestFilter) == 0 {
		err := s.getConditionsByID(ctx, owner, vf.All, p)
		errors = multierror.Append(errors, err)
	}

	if len(vf.Latest) > 0 || len(confIdLatestFilter) > 0 {
		err := s.getLatestConditions(ctx, owner, vf.Latest, query.GetConfigurationIdFilter(), p)
		errors = multierror.Append(errors, err)
	}

	for id, vf := range vf.Versions {
		err := s.getConditionsByAggregation(ctx, owner, id, vf, p)
		errors = multierror.Append(errors, err)
	}
	return errors.ErrorOrNil()
}

func (s *Store) DeleteConditions(ctx context.Context, owner string, query *pb.DeleteConditionsRequest) (int64, error) {
	return s.delete(ctx, conditionsCol, owner, query.GetIdFilter())
}

func toLatestEnabledQueryFilter() bson.D {
	key := store.LatestKey + "." + store.EnabledKey
	return bson.D{{Key: key, Value: true}}
}

func toLatestDeviceIDQueryFilter(deviceID string) bson.M {
	key := store.LatestKey + "." + store.DeviceIDFilterKey
	return bson.M{mongodb.Or: bson.A{
		bson.M{key: bson.M{mongodb.Exists: false}},
		bson.M{key: deviceID},
	}}
}

func toLatestResourceHrefQueryFilter(resourceHref string) bson.M {
	key := store.LatestKey + "." + store.ResourceHrefFilterKey
	return bson.M{mongodb.Or: bson.A{
		bson.M{key: bson.M{mongodb.Exists: false}},
		bson.M{key: resourceHref},
	}}
}

func toLatestResouceTypeQueryFilter(resourceTypeFilter []string) bson.M {
	key := store.LatestKey + "." + store.ResourceTypeFilterKey
	return bson.M{mongodb.Or: bson.A{
		bson.M{key: bson.M{mongodb.Exists: false}},
		bson.M{key: bson.M{mongodb.All: resourceTypeFilter}},
	}}
}

func toLatestConditionsQuery(owner string, queries *store.GetLatestConditionsQuery) interface{} {
	filters := make([]interface{}, 0, 5)
	filters = append(filters, toLatestEnabledQueryFilter())
	if owner != "" {
		filters = append(filters, bson.D{{Key: store.OwnerKey, Value: owner}})
	}
	if queries.DeviceID != "" {
		filters = append(filters, toLatestDeviceIDQueryFilter(queries.DeviceID))
	}
	if queries.ResourceHref != "" {
		filters = append(filters, toLatestResourceHrefQueryFilter(queries.ResourceHref))
	}
	if len(queries.ResourceTypeFilter) > 0 {
		filters = append(filters, toLatestResouceTypeQueryFilter(queries.ResourceTypeFilter))
	}
	return toFilterQuery(mongodb.And, filters)
}

func (s *Store) GetLatestEnabledConditions(ctx context.Context, owner string, query *store.GetLatestConditionsQuery, p store.ProcessConditions) error {
	if err := store.ValidateAndNormalizeConditionsQuery(query); err != nil {
		return err
	}
	opt := options.Find().SetProjection(bson.M{store.VersionsKey: false})
	cur, err := s.Collection(conditionsCol).Find(ctx, toLatestConditionsQuery(owner, query), opt)
	if err != nil {
		return err
	}
	return processCursor(ctx, cur, p)
}