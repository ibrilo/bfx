package bfx

import (
	"errors"
)

// ErrParseTicker TOWRITE
var ErrParseTicker = errors.New("failed to parse ticker data. wrong data format")

const (
	tickerTypeTrade = iota
	tickerTypeFunding
)

var tickerTypes = map[int]string{
	tickerTypeTrade:   "trade",
	tickerTypeFunding: "funding",
}

// Ticker struct represent ticker data structure
type Ticker struct {
	tickerType          int
	Symbol              string  `json:"symbol"`
	Bid                 float64 `json:"bid"`
	BidSize             float64 `json:"bid_size"`
	Ask                 float64 `json:"ask"`
	AskSize             float64 `json:"ask_size"`
	DailyChange         float64 `json:"daily_change"`
	DailyChangeRelative float64 `json:"daily_change_relative"`
	LastPrice           float64 `json:"last_price"`
	Volume              float64 `json:"volume"`
	High                float64 `json:"high"`
	Low                 float64 `json:"low"`
	FRR                 float64 `json:"frr"`
	BidPeriod           int     `json:"bid_period"`
	AskPeriod           int     `json:"ask_period"`
	FRRAmountAvailable  float64 `json:"frr_amount_available"`
}

// Type return string representation of the ticker type
func (t *Ticker) Type() string {
	return tickerTypes[t.tickerType]
}

// Parse will parse data from responses
func (t *Ticker) parse(data interface{}) error {
	switch v := data.(type) {
	case []interface{}:
		switch len(v) {
		case 11:
			t.Symbol = v[0].(string)
			t.Bid = v[1].(float64)
			t.BidSize = v[2].(float64)
			t.Ask = v[3].(float64)
			t.AskSize = v[4].(float64)
			t.DailyChange = v[5].(float64)
			t.DailyChangeRelative = v[6].(float64)
			t.LastPrice = v[7].(float64)
			t.Volume = v[8].(float64)
			t.High = v[9].(float64)
			t.Low = v[10].(float64)
		case 17:
			t.Symbol = v[0].(string)
			t.FRR = v[1].(float64)
			t.Bid = v[2].(float64)
			t.BidPeriod = int(v[3].(float64))
			t.BidSize = v[4].(float64)
			t.Ask = v[5].(float64)
			t.AskPeriod = int(v[6].(float64))
			t.AskSize = v[7].(float64)
			t.DailyChange = v[8].(float64)
			t.DailyChangeRelative = v[9].(float64)
			t.LastPrice = v[10].(float64)
			t.Volume = v[11].(float64)
			t.High = v[12].(float64)
			t.Low = v[13].(float64)
			t.FRRAmountAvailable = v[16].(float64)
		default:
			return ErrParseTicker
		}
	default:
		return ErrParseTicker
	}

	return nil
}

// Tickers contain tickers set
type Tickers []Ticker

func parseTickers(data interface{}) (Tickers, error) {
	var set Tickers
	switch v := data.(type) {
	case []interface{}:
		for _, elem := range v {
			t := Ticker{}
			if err := t.parse(elem); err != nil {
				return nil, ErrParseTicker
			}
		}
	default:
		return nil, ErrParseTicker
	}
	return set, nil
}

// Trades return only trades tickers
func (tics Tickers) Trades() Tickers {
	var set []Ticker
	for _, tic := range tics {
		if tic.tickerType == tickerTypeTrade {
			set = append(set, tic)
		}
	}
	return set
}

// Fundings return only funding tickers
func (tics Tickers) Fundings() Tickers {
	var set []Ticker
	for _, tic := range tics {
		if tic.tickerType == tickerTypeFunding {
			set = append(set, tic)
		}
	}
	return set
}
