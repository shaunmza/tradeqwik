package main

import (
	"fmt"
	"time"

	"github.com/shaunmza/tradeqwik"
)

// Channel into which ticker data is sent
var rTradeRefreshPeriod int
var oTradeRefreshPeriod int
var rChan chan *tradeqwik.RecentTrades
var oChan chan *tradeqwik.OpenTrades

func main() {
	var r *tradeqwik.RecentTrades
	var o *tradeqwik.OpenTrades

	// Initialse, so we can get the channel to receive updates from
	rChan = make(chan *tradeqwik.RecentTrades)
	oChan = make(chan *tradeqwik.OpenTrades)

	watchRecentTrades("VIVA", "BTC")
	watchOpenTrades("VIVA", "BTC")
	watchOpenTrades("VIVA", "USD")
	watchOpenTrades("VIVA", "LTC")

	for {
		select {
		case o = <-oChan:
			fmt.Println("Asks:")
			if o.Error != nil {
				fmt.Printf("Open Trade Error! %s, Last Updated: %s\n", o.Error, o.LastUpdate)
			}

			// Just print it out for now
			for _, ask := range o.Asks {
				fmt.Printf("%s/%s Amount: %f Price: %f\n", o.Base, o.Counter, ask.Amount, ask.Price)
			}
			fmt.Println("----------------------------------------")

			for _, bid := range o.Bids {
				fmt.Printf("%s/%s Amount: %f Price: %f\n", o.Base, o.Counter, bid.Amount, bid.Price)
			}
			fmt.Println("----------------------------------------")

		case r = <-rChan:
			// If this is not nil then we encountered a problem, use this to determine
			// what to do next.
			// LastUpdate can be used to determine how stale the data is
			if r.Error != nil {
				fmt.Printf("Recent Trades Error! %s, Last Updated: %s\n", r.Error, r.LastUpdate)
			}

			// Just print it out for now
			for _, trade := range r.Trades {
				fmt.Printf("%s/%s Type: %s Amount: %f Price: %f\n", r.Base, r.Counter, trade.Type, trade.Amount, trade.Price)
			}
			fmt.Println("=======================================================")
		}
	}
}

func watchRecentTrades(base string, counter string) {
	// Don't call too often
	if rTradeRefreshPeriod < 10 {
		rTradeRefreshPeriod = 10
	}

	// Call now so we don't have to wait for the first batch of data
	go func() {
		s, err := tradeqwik.GetRecentTrades(base, counter)
		if err == nil {
			rChan <- s
		}
	}()

	ticker := time.NewTicker(time.Second * time.Duration(rTradeRefreshPeriod))
	go func() {
		for _ = range ticker.C {
			s, err := tradeqwik.GetRecentTrades(base, counter)
			if err == nil {
				rChan <- s
			}
		}
	}()
}

func watchOpenTrades(base string, counter string) {
	// Don't call too often
	if oTradeRefreshPeriod < 10 {
		oTradeRefreshPeriod = 10
	}

	// Call now so we don't have to wait for the first batch of data
	go func() {
		s, err := tradeqwik.GetOpenTrades(base, counter)
		if err == nil {
			oChan <- s
		}
	}()

	ticker := time.NewTicker(time.Second * time.Duration(oTradeRefreshPeriod))
	go func() {
		for _ = range ticker.C {
			s, err := tradeqwik.GetOpenTrades(base, counter)
			if err == nil {
				oChan <- s
			}
		}
	}()
}
