package tradeqwik

import "time"

// The Ticker struct holds a map of CurrencyPair structs, the last time
// data was fetched successfully and in the case of an error, the error message
type Ticker struct {
	CurrencyPairs []*CurrencyPair
	LastUpdate    time.Time
	Error         error
}

// CurrencyPair holds the data pertaining to a specific pair of currencies
type CurrencyPair struct {
	Base    string  `json:"base"`
	Counter string  `json:"counter"`
	First   float64 `json:"first,string"`
	Low     float64 `json:"low,string"`
	High    float64 `json:"high,string"`
	Last    float64 `json:"last,string"`
	Volume  float64 `json:"volume,string"`
}

type OpenTrades struct {
	Base       string
	Counter    string
	Asks       []OpenTrade `json:"asks,string"`
	Bids       []OpenTrade `json:"bids,string"`
	LastUpdate time.Time
	Error      error
}

type OpenTrade struct {
	Amount float64 `json:"amount,string"`
	Price  float64 `json:"price,string"`
}

type RecentTrades struct {
	Base       string
	Counter    string
	Trades     []*RecentTrade
	LastUpdate time.Time
	Error      error
}

type RecentTrade struct {
	Type    string  `json:"typ"`
	Amount  string  `json:"amount,string"`
	Price   float64 `json:"price,string"`
	Fee     float64 `json:"fee,string"`
	Created int64   `json:"created"`
}

type AccountOpenTrades struct {
	Trades     []*AccountOpenTrade
	LastUpdate time.Time
	Error      error
}

type AccountOpenTrade struct {
	ID      int64   `json:"id"`
	Type    string  `json:"typ"`
	Amount  float64 `json:"amount,string"`
	Price   float64 `json:"price,string"`
	Created int64   `json:"created"`
}

type AccountHistory struct {
	APIKey string `json:"api_key"`
	Before int64  `json:"before"`
	Limit  int    `json:"limit"`
	LastId int64  `json:"last_id"`
}

var BaseURL = "https://www.tradeqwik.com"
