package bfx

var defaultHTTPScheme = "https"

var (
	publicRESTDomain  = "api-pub.bitfinex.com"
	privateRESTDomain = "api.bitfinex.com"

	endpointRESTPublicPath  = defaultHTTPScheme + "://" + publicRESTDomain + "/v2"
	endpointRESTPrivatePath = defaultHTTPScheme + "://" + privateRESTDomain + "/v2"
)
