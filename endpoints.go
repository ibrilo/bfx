package bfx

import (
	"net/url"
	"strings"
)

var defaultHTTPScheme = "https"

var (
	publicRESTDomain  = "api-pub.bitfinex.com"
	privateRESTDomain = "api.bitfinex.com"

	endpointRESTPublicPath  = defaultHTTPScheme + "://" + publicRESTDomain + "/v2"
	endpointRESTPrivatePath = defaultHTTPScheme + "://" + privateRESTDomain + "/v2"

	endpointPublicPlatformStatus = endpointRESTPublicPath + "/platform/status"
	endpointPublicTickers        = func(symbols []string) string {
		return endpointRESTPublicPath + "/tickers?symbols=" + strings.Join(symbols, ",")
	}
	endpointPublicTicker = func(symbol string) string {
		return endpointRESTPublicPath + "/ticker/" + symbol
	}
	endpointPublicTrades = func(symbol string, params url.Values) string {
		return endpointRESTPublicPath + "/trades/" + symbol + "/hist" + queryParamString(params)
	}
	endpointPublicBook = func(symbol, precision string, params url.Values) string {
		return endpointRESTPublicPath + "/book/" + symbol + "/" + precision + queryParamString(params)
	}
	endpointPublicStats = func(key, size, symbol, side, section string, params url.Values) string {
		return endpointRESTPublicPath + "/stats1/" + joinPathParams(key, size, symbol, side) + "/" + section + queryParamString(params)
	}
	endpointPublicCandles = func(timeframe, symbol, section string, params url.Values) string {
		return endpointRESTPublicPath + "/candles/" + joinPathParams("trade", timeframe, symbol) + "/" + section + queryParamString(params)
	}
	endpointPublicConfigs = func(action, object, detail string) string {
		return endpointRESTPublicPath + "/conf/" + joinPathParams("pub", action, object, detail)
	}

	endpointPublicStatus = "" // TODO: implementation

	endpointPublicLiquidationFeed = func(params url.Values) string {
		return endpointRESTPublicPath + "/liquidations/hist" + queryParamString(params)
	}
	endpointPublicRankings = func(key, timeframe, symbol, section string, params url.Values) string {
		return endpointRESTPublicPath + "/rankings/" + joinPathParams(key, timeframe, symbol) + "/" + section + queryParamString(params)
	}
	endpointPublicPulseHistory = func(params url.Values) string {
		return endpointRESTPublicPath + "/pulse/hist" + queryParamString(params)
	}
	endpointPublicPulseProfileDetails = func(nickname string) string {
		return endpointRESTPublicPath + "/pulse/profile/" + nickname
	}
	endpointPublicFundingStats = func(symbol string) string {
		return endpointRESTPublicPath + "/funding/stats/" + symbol + "/hist"
	}
)

func joinPathParams(params ...string) string {
	return strings.Join(params, ":")
}

func queryParamString(params url.Values) string {
	q := ""
	if len(params) != 0 {
		q = "?" + params.Encode()
	}
	return q
}
