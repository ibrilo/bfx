package bfx

import (
	"net/url"
	"strings"
)

var defaultHTTPScheme = "https"

func queryParamString(params url.Values) string {
	q := ""
	if len(params) != 0 {
		q = "?" + params.Encode()
	}
	return q
}

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
		return endpointRESTPublicPath + "/trades/" + symbol + "/hyst" + queryParamString(params)
	}
	endpointPublicBook = func(symbol, precision string, params url.Values) string {
		return endpointRESTPublicPath + "/book/" + symbol + "/" + precision + queryParamString(params)
	}
	endpointPublicStats = func(key, size, symbol, side, section string, params url.Values) string {
		return endpointRESTPublicPath + "/stats1/" + strings.Join([]string{key, size, symbol, side}, ":") + "/" + section + queryParamString(params)
	}
	endpointPublicCandles = func(timeframe, symbol, section string, params url.Values) string {
		return endpointRESTPublicPath + "/candles/" + strings.Join([]string{"trade", timeframe, symbol}, ":") + "/" + section + queryParamString(params)
	}
	endpointPublicConfigs = func(action, object, detail string) string {
		return endpointRESTPublicPath + "/conf/" + strings.Join([]string{"pub", action, object, detail}, ":")
	}

	endpointPublicStatus = "" // TODO: implementation

	endpointPublicLiquidationFeed = func(params url.Values) string {
		return endpointRESTPublicPath + "/liquidations/hist" + queryParamString(params)
	}
	endpointPublicRankings = func(key, timeframe, symbol, section string, params url.Values) string {
		return endpointRESTPublicPath + "/rankings/" + strings.Join([]string{key, timeframe, symbol}, ":") + "/" + section + queryParamString(params)
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
