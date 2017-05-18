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
	Base    string `json:"base"`
	Counter string `json:"counter"`
	First   string `json:"first"`
	Low     string `json:"low"`
	High    string `json:"high"`
	Last    string `json:"last"`
	Volume  string `json:"volume"`
}

type OpenTrades struct {
	Base    string
	Counter string
	Asks    []OpenTrade `json:"asks"`
	Bids    []OpenTrade `json:"bids"`
}

type OpenTrade struct {
	Amount string `json:"amount"`
	Price  string `json:"price"`
}

type RecentTrades struct {
	Base    string
	Counter string
	Type    string `json:"typ"`
	Amount  string `json:"amount"`
	Price   string `json:"price"`
	Fee     string `json:"fee"`
	Created string `json:"created"`
}
