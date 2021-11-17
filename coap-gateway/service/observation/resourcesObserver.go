package observation

import (
	"context"
	"fmt"
	"sync"

	"github.com/plgd-dev/device/schema/interfaces"
	"github.com/plgd-dev/go-coap/v2/message"
	"github.com/plgd-dev/go-coap/v2/tcp"
	"github.com/plgd-dev/go-coap/v2/tcp/message/pool"
	"github.com/plgd-dev/hub/coap-gateway/resource"
	"github.com/plgd-dev/hub/pkg/log"
	"github.com/plgd-dev/hub/resource-aggregate/commands"
)

type OnObserveResource = func(ctx context.Context, deviceID, resourceHref string, notification *pool.Message) error
type OnGetResourceContent = func(ctx context.Context, deviceID, resourceHref string, notification *pool.Message) error

// ResourcesObserver is a thread-safe type that handles observation of resources belonging to
// a single device.
//
// The resource observer keeps track of observed resources to avoid multiple observation of the
// same resource. Each new unique observation fires an event:
//   - If the resource is observable then the connection to COAP-GW (coapConn) is used to
//   register for observations from COAP-GW. Observation notifications are handled by the
//   onObserveResource callback.
//   - If the resource is not observable then a GET request is sent to COAP-GW to receive
//   the content of the resource and the response is handled by the onGetResourceContent
//   callback.
type resourcesObserver struct {
	lock                 sync.Mutex
	deviceID             string
	resources            map[int64]*observedResource // [instanceID]
	coapConn             *tcp.ClientConn
	onObserveResource    OnObserveResource
	onGetResourceContent OnGetResourceContent
}

// Create new Resource Observer.
//
// All arguments (coapConn, onObserveResource and onGetResourceContent) must be non-nil,
// otherwise the function will panic.
func newResourcesObserver(deviceID string, coapConn *tcp.ClientConn, onObserveResource OnObserveResource, onGetResourceContent OnGetResourceContent) *resourcesObserver {
	fatalCannotCreate := func(err error) {
		log.Fatal("cannot create resource observer: %v", err)
	}
	if deviceID == "" {
		fatalCannotCreate(fmt.Errorf("empty deviceID"))
	}
	if coapConn == nil {
		fatalCannotCreate(fmt.Errorf("invalid coap-gateway connection"))
	}
	if onObserveResource == nil {
		fatalCannotCreate(fmt.Errorf("invalid onObserveResource callback"))
	}
	if onGetResourceContent == nil {
		fatalCannotCreate(fmt.Errorf("invalid onGetResourceContent callback"))
	}
	return &resourcesObserver{
		deviceID:             deviceID,
		resources:            make(map[int64]*observedResource),
		coapConn:             coapConn,
		onObserveResource:    onObserveResource,
		onGetResourceContent: onGetResourceContent,
	}
}

// Add resource to observer with given interface and wait for initialization message.
func (o *resourcesObserver) addAndWaitForResource(ctx context.Context, res *commands.Resource, obsInterface string) (bool, error) {
	o.lock.Lock()
	defer o.lock.Unlock()
	obs, err := o.addResourceLocked(ctx, res, obsInterface)
	if err != nil {
		return false, err
	}
	opened, err := obs.IsOpened(ctx)
	if err != nil {
		return false, err
	}
	return opened, nil
}

func (o *resourcesObserver) addResourceLocked(ctx context.Context, res *commands.Resource, obsInterface string) (*observedResource, error) {
	resID := res.GetResourceID()
	log.Debugf("observation of resource(%v) requested", resID)
	addObservationError := func(err error) error {
		return fmt.Errorf("cannot add resource observation: %w", err)
	}
	if o.deviceID == "" {
		return nil, addObservationError(fmt.Errorf("empty deviceID"))
	}
	instanceID := resource.GetInstanceID(resID.GetHref())
	if o.deviceID != resID.GetDeviceId() {
		return nil, addObservationError(fmt.Errorf("invalid deviceID(%v)", resID.GetDeviceId()))
	}
	href := resID.GetHref()
	if obsRes, ok := o.resources[instanceID]; ok {
		return obsRes, nil
	}
	obsRes := NewObservedResource(href, obsInterface)
	if err := o.handleResourceLocked(ctx, obsRes, res.IsObservable()); err != nil {
		return nil, addObservationError(err)
	}
	o.resources[instanceID] = obsRes
	return obsRes, nil
}

// Handle given resource.
//
// For observable resources subscribe to observations, for unobservable resources retrieve
// their content.
func (o *resourcesObserver) handleResourceLocked(ctx context.Context, obsRes *observedResource, isObservable bool) error {
	if obsRes.Href() == commands.StatusHref {
		log.Debugf("observation of resource /%v%v skipped", o.deviceID, obsRes.Href())
		return nil
	}

	if isObservable {
		obs, err := o.observeResourceLocked(ctx, obsRes)
		if err != nil {
			return err
		}
		obsRes.SetObservation(obs)
		return nil
	}
	return o.getResourceContentLocked(ctx, obsRes.Href())
}

// Register to COAP-GW resource observation for given resource
func (o *resourcesObserver) observeResourceLocked(ctx context.Context, obsRes *observedResource) (*tcp.Observation, error) {
	cannotObserveResourceError := func(deviceID, href string, err error) error {
		return fmt.Errorf("cannot observe resource /%v%v: %w", deviceID, href, err)
	}
	if o.deviceID == "" {
		return nil, cannotObserveResourceError(o.deviceID, obsRes.Href(), fmt.Errorf("empty deviceID"))
	}
	initialized := false
	obs, err := o.coapConn.Observe(ctx, obsRes.Href(), func(msg *pool.Message) {
		if !initialized {
			initialized = true
			observable := true
			if _, errObs := msg.Observe(); errObs != nil {
				log.Debugf("href: %v not observable err: %v", obsRes.Href(), errObs)
				observable = false
			}
			obsRes.SetOpened(observable)
			if !observable {
				// TODO: not necessary when go-coap is updated
				o.cancelResourcesObservations(ctx, []string{obsRes.href})

				// fallback to get for non-observable resource
				if errGet := o.onGetResourceContent(ctx, o.deviceID, obsRes.Href(), msg); errGet != nil {
					log.Errorf("cannot get resource /%v%v content: %w", o.deviceID, obsRes.Href(), errGet)
				}
				return
			}
		}
		if err2 := o.onObserveResource(ctx, o.deviceID, obsRes.Href(), msg); err2 != nil {
			log.Error(cannotObserveResourceError(o.deviceID, obsRes.Href(), err2))
			return
		}
	},
		// TODO: remove for OC_IF_BASELINE
		message.Option{
			ID:    message.URIQuery,
			Value: []byte("if=" + obsRes.Interface()),
		})
	if err != nil {
		return nil, cannotObserveResourceError(o.deviceID, obsRes.Href(), err)
	}
	return obs, nil
}

// Request resource content form COAP-GW
func (o *resourcesObserver) getResourceContentLocked(ctx context.Context, href string) error {
	cannotGetResourceError := func(deviceID, href string, err error) error {
		return fmt.Errorf("cannot get resource /%v%v content: %w", deviceID, href, err)
	}
	if o.deviceID == "" {
		return cannotGetResourceError(o.deviceID, href, fmt.Errorf("empty deviceID"))
	}
	resp, err := o.coapConn.Get(ctx, href,
		// TODO: remove for OC_IF_BASELINE
		message.Option{
			ID:    message.URIQuery,
			Value: []byte("if=" + interfaces.OC_IF_BASELINE),
		})
	defer pool.ReleaseMessage(resp)
	if err != nil {
		return cannotGetResourceError(o.deviceID, href, err)
	}
	if err := o.onGetResourceContent(ctx, o.deviceID, href, resp); err != nil {
		return cannotGetResourceError(o.deviceID, href, err)
	}
	return nil
}

// Add multiple resources to observer.
func (o *resourcesObserver) addResources(ctx context.Context, resources []*commands.Resource) error {
	o.lock.Lock()
	defer o.lock.Unlock()
	return o.addResourcesLocked(ctx, resources)
}

func (o *resourcesObserver) addResourcesLocked(ctx context.Context, resources []*commands.Resource) error {
	var errors []error
	for _, resource := range resources {
		_, err := o.addResourceLocked(ctx, resource, interfaces.OC_IF_BASELINE)
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("cannot add resources to observe: %v", errors)
	}
	return nil
}

// Get list of observable and non-observable resources added to resourcesObserver.
//
// Empty instanceIDs parameter is ignored and function will return all resources. Otherwise only
// resources with instanceID value contained in the instanceIDs array are returned.
func (o *resourcesObserver) getResources(instanceIDs []int64) []*commands.ResourceId {
	getAllDeviceIDMatches := len(instanceIDs) == 0

	uniqueInstanceIDs := make(map[int64]struct{})
	for _, v := range instanceIDs {
		uniqueInstanceIDs[v] = struct{}{}
	}

	matches := make([]*commands.ResourceId, 0, 16)
	o.lock.Lock()
	defer o.lock.Unlock()
	if getAllDeviceIDMatches {
		for _, value := range o.resources {
			matches = append(matches, &commands.ResourceId{
				DeviceId: o.deviceID,
				Href:     value.Href(),
			})
		}
		return matches
	}

	for instanceID := range uniqueInstanceIDs {
		if resource, ok := o.resources[instanceID]; ok {
			matches = append(matches, &commands.ResourceId{
				DeviceId: o.deviceID,
				Href:     resource.Href(),
			})
		}
	}
	return matches
}

// Cancel observation of given resources.
func (o *resourcesObserver) cancelResourcesObservations(ctx context.Context, hrefs []string) {
	observations := o.popTrackedObservations(ctx, hrefs)
	for _, obs := range observations {
		if err := obs.Cancel(ctx); err != nil {
			log.Debugf("cannot cancel resource observation: %w", err)
		}
	}
}

func (o *resourcesObserver) popTrackedObservations(ctx context.Context, hrefs []string) []*tcp.Observation {
	observations := make([]*tcp.Observation, 0, 32)
	o.lock.Lock()
	defer o.lock.Unlock()
	for _, href := range hrefs {
		var instanceID int64
		found := false
		for insID, r := range o.resources {
			if r.Href() == href {
				found = true
				instanceID = insID
				break
			}
		}
		if !found {
			log.Errorf("no observed resource with given href(%v) found", href)
			continue
		}
		obs := o.popObservationLocked(ctx, instanceID)
		if obs != nil {
			observations = append(observations, obs)
			log.Debugf("canceling observation on resource(/%v%v)", o.deviceID, href)
		}
		o.removeResourceLocked(instanceID)
	}
	return observations
}

func (o *resourcesObserver) popObservationLocked(ctx context.Context, instanceID int64) *tcp.Observation {
	log.Debugf("remove published resource ocf://%v/%v", o.deviceID, instanceID)
	var r *observedResource
	if res, ok := o.resources[instanceID]; ok {
		r = res
	}
	if r != nil {
		opened, err := r.IsOpened(ctx)
		if err != nil {
			log.Errorf("cannot get resource('/%v%v') observation state: %w", o.deviceID, r.Href(), err)
			opened = true
		}
		obs := r.PopObservation()
		if !opened {
			log.Debugf("skipping canceling of closed resource('/%v%v') observation", o.deviceID, r.Href())
			return nil
		}
		return obs
	}

	return nil
}

func (o *resourcesObserver) removeResourceLocked(instanceID int64) {
	if obs, ok := o.resources[instanceID]; ok {
		obs.SetOpened(false)
	}
	delete(o.resources, instanceID)
}

// Remove all observations.
func (o *resourcesObserver) CleanObservedResources(ctx context.Context) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.cleanObservedResourcesLocked(ctx)
}

func (o *resourcesObserver) cleanObservedResourcesLocked(ctx context.Context) {
	observedResources := o.popObservedResourcesLocked()
	for _, obs := range observedResources {
		opened, err := obs.IsOpened(ctx)
		if err != nil {
			log.Errorf("cannot get resource('/%v%v') observation state: %w", o.deviceID, obs.Href(), err)
			opened = true
		}
		if v := obs.PopObservation(); v != nil {
			if !opened {
				log.Debugf("skipping canceling of closed resource('/%v%v') observation", o.deviceID, obs.Href())
				continue
			}
			if err := v.Cancel(ctx); err != nil {
				log.Errorf("cannot cancel resource('/%v%v') observation: %w", o.deviceID, obs.Href(), err)
			}
		}
	}
}

func (o *resourcesObserver) popObservedResourcesLocked() map[int64]*observedResource {
	observations := o.resources
	o.resources = make(map[int64]*observedResource)
	return observations
}