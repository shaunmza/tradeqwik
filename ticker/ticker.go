package ticker

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/shaunmza/tradeqwik"
)

// Channel into which ticker data is sent
var tickerCh chan tradeqwik.Ticker
var TickerRefreshPeriod int

// Init is called so we can return the channels.
// This allows messages to be read off of it,
// and we can close it from outside too
func Init() chan tradeqwik.Ticker {
	tickerCh = make(chan tradeqwik.Ticker)

	startTicker()

	return tickerCh
}

// startTicker runs a ticker to fetch the data periodically
func startTicker() {
	// Don't call too often
	if TickerRefreshPeriod < 10 {
		TickerRefreshPeriod = 10
	}

	// Call now so we don't have to wait for the first batch of data
	go func() {
		s := getTickerData()
		tickerCh <- s
	}()

	ticker := time.NewTicker(time.Second * time.Duration(TickerRefreshPeriod))
	go func() {
		for _ = range ticker.C {
			s := getTickerData()
			tickerCh <- s
		}
	}()
}

// getData makes the actual http request, parses the JSON and returns
// the data in a struct
func getTickerData() tradeqwik.Ticker {

	resp, err := http.Get(tradeqwik.BaseURL + "/api/1/ticker")
	if err != nil {
		return tradeqwik.Ticker{Error: err}
	}

	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return tradeqwik.Ticker{Error: err}
	}

	var j []*tradeqwik.CurrencyPair
	if err := json.Unmarshal(ret, &j); err != nil {
		log.Fatal(err)
		return tradeqwik.Ticker{Error: err}
	}

	r := tradeqwik.Ticker{CurrencyPairs: j, LastUpdate: time.Now()}
	return r
}

/*
// WatchCoins runs a ticker to fetch the data periodically
func WatchCoins(coins []string, period int) {
	// Don't call too often
	if TickerRefreshPeriod == nil || TickerRefreshPeriod < 10 {
		TickerRefreshPeriod = 10
	}

	// Call now so we don't have to wait for the first batch of data
	go func() {
		s := getData(coins)
		coinCh <- s
	}()

	ticker := time.NewTicker(time.Second * time.Duration(period))
	go func() {
		for _ = range ticker.C {
			s := getData(coins)
			coinCh <- s
		}
	}()
}

// getData makes the actual http request, parses the JSON and returns
// the data in a struct
func getData(coins []string) data.Ticker {

	resp, err := http.Get(tickerURL)
	if err != nil {
		return data.Ticker{Error: err}
	}

	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return data.Ticker{Error: err}
	}

	var j []*data.Coin
	var res []*data.Coin
	if err := json.Unmarshal(ret, &j); err != nil {
		log.Fatal(err)
		return data.Ticker{Error: err}
	}

	for _, c1 := range j {
		for _, c2 := range coins {
			if c1.ID == c2 {
				res = append(res, c1)
			}
		}

	}

	r := data.Ticker{Coins: res, LastUpdate: time.Now()}
	return r
}
*/
