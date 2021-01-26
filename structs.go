package bfx

import (
	"errors"
	"fmt"
	"time"
)

// ErrParseTicker TOWRITE
var errParseTicker = errors.New("failed to parse ticker data. wrong data format")
var errParseTickers = errors.New("failed to parse tickers data. wrong format")

// ErrParseTrade TOWRITE
var errParseTrade = errors.New("failed to parse trade data. wrong data format")
var errParseTrades = errors.New("failed to parse trades data. wrong data format")

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
func (t *Ticker) parse(data interface{}, symbol string) error {
	switch v := data.(type) {
	case []interface{}:
		switch len(v) {
		case 10:
			t.tickerType = tickerTypeTrade
			t.Symbol = symbol
			t.Bid = v[0].(float64)
			t.BidSize = v[1].(float64)
			t.Ask = v[2].(float64)
			t.AskSize = v[3].(float64)
			t.DailyChange = v[4].(float64)
			t.DailyChangeRelative = v[5].(float64)
			t.LastPrice = v[6].(float64)
			t.Volume = v[7].(float64)
			t.High = v[8].(float64)
			t.Low = v[9].(float64)
		case 11:
			t.tickerType = tickerTypeTrade
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
		case 16:
			t.tickerType = tickerTypeFunding
			t.Symbol = symbol
			t.FRR = v[0].(float64)
			t.Bid = v[1].(float64)
			t.BidPeriod = int(v[2].(float64))
			t.BidSize = v[3].(float64)
			t.Ask = v[4].(float64)
			t.AskPeriod = int(v[5].(float64))
			t.AskSize = v[6].(float64)
			t.DailyChange = v[7].(float64)
			t.DailyChangeRelative = v[8].(float64)
			t.LastPrice = v[9].(float64)
			t.Volume = v[10].(float64)
			t.High = v[11].(float64)
			t.Low = v[12].(float64)
			t.FRRAmountAvailable = v[15].(float64)
		case 17:
			t.tickerType = tickerTypeFunding
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
			return errParseTicker
		}
	default:
		return errParseTicker
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
			if err := t.parse(elem, ""); err != nil {
				return nil, err
			}
			set = append(set, t)
		}
	default:
		return nil, errParseTickers
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

const (
	tradeTypeTrade = iota
	tradeTypeFunding
)

var tradeTypes = map[int]string{
	tradeTypeTrade:   "trade",
	tradeTypeFunding: "funding",
}

// Trade TOWRITE
type Trade struct {
	tradeType int
	ID        int
	MTS       time.Time
	Amount    float64
	Price     float64
	Rate      float64
	Period    int
}

// Type TOWRITE
func (t *Trade) Type() string {
	return tradeTypes[t.tradeType]
}

func (t *Trade) parse(data interface{}) error {
	v, ok := data.([]interface{})
	if !ok {
		if Debug {
			fmt.Printf("Passed type: %T\nActually: %v\n", data, data)
		}
		return errParseTrade
	}

	if len(v) == 4 {
		t.ID = v[0].(int)
		t.MTS = time.Unix(v[1].(int64), 0)
		t.Amount = v[2].(float64)
		t.Price = v[3].(float64)
		return nil
	}
	if len(v) == 5 {
		t.ID = v[0].(int)
		t.MTS = time.Unix(v[1].(int64), 0)
		t.Amount = v[2].(float64)
		t.Rate = v[3].(float64)
		t.Period = v[4].(int)
		return nil
	}

	return errParseTrade
}

type trades []Trade

func parseTrades(data interface{}) (trades, error) {
	var set trades
	v, ok := data.([]interface{})
	if !ok {
		return nil, errParseTrades
	}

	for _, trade := range v {
		t := Trade{}
		if err := t.parse(trade); err != nil {
			return nil, err
		}
		set = append(set, t)
	}

	return set, nil
}
