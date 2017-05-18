package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/shaunmza/tradeqwik"
)

// Channel into which ticker data is sent
var rTradeRefreshPeriod int
var oTradeRefreshPeriod int
var rChan chan *tradeqwik.RecentTrades
var oChan chan *tradeqwik.OpenTrades

func main() {
	//var r *tradeqwik.RecentTrades
	var o *tradeqwik.OpenTrades

	// Initialse, so we can get the channel to receive updates from
	rChan = make(chan *tradeqwik.RecentTrades)
	oChan = make(chan *tradeqwik.OpenTrades)

	//trades.WatchRecentTrades("VIVA", "BTC")
	watchOpenTrades("VIVA", "BTC")
	watchOpenTrades("VIVA", "USD")
	watchOpenTrades("VIVA", "LTC")

	for {
		o = <-oChan

		if o.Error != nil {
			fmt.Printf("Open Trade Error! %s, Last Updated: %s\n", o.Error, o.LastUpdate)
		}

		// Just print it out for now
		for _, ask := range o.Asks {

			fmt.Println(o.Base + "/" + o.Counter + " " + strconv.FormatFloat(ask.Amount, 'f', 8, 64) + " " + strconv.FormatFloat(ask.Price, 'f', 8, 64))
		}
		fmt.Println("----------------------------------------")

		for _, bid := range o.Bids {

			fmt.Println(o.Base + "/" + o.Counter + " " + strconv.FormatFloat(bid.Amount, 'f', 8, 64) + " " + strconv.FormatFloat(bid.Price, 'f', 8, 64))
		}
		fmt.Println("----------------------------------------")
	}
	/*
		go func() {
			// Infinite loop so we keep getting ticker info
			for {

				// Get off of the channel
				r = <-rChan

				// If this is not nil then we encountered a problem, use this to determine
				// what to do next.
				// LastUpdate can be used to determine how stale the data is
				if r.Error != nil {
					fmt.Printf("Recent Trades Error! %s, Last Updated: %s\n", r.Error, r.LastUpdate)
				}

				// Just print it out for now
				for _, trade := range r.Trades {

					fmt.Println(r.Base + "/" + r.Counter + "/" + trade.Amount + " " + trade.Price + " " + trade.Type)
				}
				fmt.Println("=======================================================")

			}
		}()
	*/
}

func watchRecentTrades(base string, counter string) {
	// Don't call too often
	if rTradeRefreshPeriod < 10 {
		rTradeRefreshPeriod = 10
	}

	// Call now so we don't have to wait for the first batch of data
	go func() {
		s, err := tradeqwik.GetRecentTrades(base, counter)
		if err != nil {
			rChan <- s
		}
	}()

	ticker := time.NewTicker(time.Second * time.Duration(rTradeRefreshPeriod))
	go func() {
		for _ = range ticker.C {
			s, err := tradeqwik.GetRecentTrades(base, counter)
			if err != nil {
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
		fmt.Println(s)
		if err != nil {
			oChan <- s
		}
	}()

	ticker := time.NewTicker(time.Second * time.Duration(oTradeRefreshPeriod))
	go func() {
		for _ = range ticker.C {
			s, err := tradeqwik.GetOpenTrades(base, counter)
			if err != nil {
				oChan <- s
			}
		}
	}()
}
