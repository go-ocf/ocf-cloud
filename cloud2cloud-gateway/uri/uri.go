package uri

// Resource Service URIs.
const (
	API     string = "/api"
	Version string = API + "/v1"
	Devices string = Version + "/devices"
	Device  string = Devices + "/{deviceID}"

	ResourceValues string = Devices + "/{deviceID}/{resourceHref}"

	DevicesSubscriptions string = Devices + "/subscriptions"
	DevicesSubscription  string = Devices + "/subscriptions/{subscriptionID}"

	DeviceSubscriptions string = Device + "/subscriptions"
	DeviceSubscription  string = Device + "/subscriptions/{subscriptionID}"

	ResourceSubscriptions string = Devices + "/{deviceID}/{resourceHref}/subscriptions"
	ResourceSubscription  string = Devices + "/{deviceID}/{resourceHref}/subscriptions/{subscriptionID}"
)
