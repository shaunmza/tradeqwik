package main

import (
	"fmt"
	"time"

	"github.com/shaunmza/coinmarketcap"
	"github.com/shaunmza/tradeqwik"
	"github.com/shaunmza/tradeqwik/trading"
)

type price struct {
	Base    string
	Counter string
	Price   float64
}

var c *coinmarketcap.Ticker

var oTradeRefreshPeriod int
var cRefreshPeriod int

var oChan chan *tradeqwik.OpenTrades
var cChan chan *coinmarketcap.Ticker

var vivaTargetPrice float64
var priceTargets map[string]price
var openTrades map[string]*tradeqwik.OpenTrades
var targetSpread float64

func main() {
	// 1 VIVA ~ 5.5 USD
	vivaTargetPrice = 5.5
	// withing a 10% range
	targetSpread = 0.1
	// look for new trades every 60 seconds
	oTradeRefreshPeriod = 60

	// Put your API key in here
	trading.Init("YOUR API KEY")

	// Initialise these, if you add more currencies, change the 1 to whatever
	priceTargets = make(map[string]price, 1)
	openTrades = make(map[string]*tradeqwik.OpenTrades, 1)

	t := make([]string, 0)
	t = append(t, "bitcoin")
	// Only btc for now, but you can have more
	/*t = append(t, "litecoin")
	t = append(t, "steem")
	t = append(t, "steem-dollars")*/

	// Coinmarketcap endpoints are updated every 5 minutes, se we use that here
	period := 60 * 5
	ticker := time.NewTicker(time.Second * time.Duration(period))

	// Because we are impatient, call it now
	r, err := coinmarketcap.GetData(t)

	// If this is not nil then we encountered a problem, use this to determine
	// what to do next.
	// LastUpdate can be used to determine how stale the data is
	if err != nil {
		fmt.Printf("Error! %s, Last Updated: %s\n", err, r.LastUpdate)
	}

	mapCoins(r)

	// We can do this for many different pairs, only BTC for now
	watchOpenTrades("VIVA", "BTC")

	// Infinite loop so we keep getting prices
	for _ = range ticker.C {
		// Get latest prices
		r, err = coinmarketcap.GetData(t)

		// If this is not nil then we encountered a problem, use this to determine
		// what to do next.
		// LastUpdate can be used to determine how stale the data is
		if err != nil {
			fmt.Printf("Error! %s, Last Updated: %s\n", err, r.LastUpdate)
		}

		// Set our prices
		mapCoins(r)

		// Gets called from watchOpenTrades too
		matchOpenTrades()

	}
}

func matchOpenTrades() {
	fmt.Println("Going to match open trades")

	for s, priceT := range priceTargets {
		minAskPrice := priceT.Price + (priceT.Price * targetSpread)
		maxBidPrice := priceT.Price - (priceT.Price * targetSpread)
		fmt.Printf("Price is: %f\nMin ask is: %f\nMax bid is: %f\n", priceT.Price, minAskPrice, maxBidPrice)

		for _, ask := range openTrades[s].Asks {
			if ask.Price < maxBidPrice {
				//create bid to match the ask
				fmt.Printf("\nCreating a bid for %f at %f\n", ask.Amount, ask.Price)
				r, err := trading.Buy(priceT.Base, priceT.Counter, ask.Amount, ask.Price)
				if err != nil {
					// Do something here
					fmt.Println(err)
				} else {
					fmt.Println(r)
				}
			}
		}

		for _, bid := range openTrades[s].Bids {
			if bid.Price > minAskPrice {
				//create ask to match open bid
				fmt.Printf("\nCreating a ask for %f at %f\n", bid.Amount, bid.Price)
				r, err := trading.Buy(priceT.Base, priceT.Counter, bid.Amount, bid.Price)
				if err != nil {
					// Do something here
					fmt.Println(err)
				} else {
					fmt.Println(r)
				}
			}
		}
	}
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
			mapOpenTrades(base, counter, s)
			// React to this now
			matchOpenTrades()
		}
	}()

	ticker := time.NewTicker(time.Second * time.Duration(oTradeRefreshPeriod))
	go func() {
		for _ = range ticker.C {
			s, err := tradeqwik.GetOpenTrades(base, counter)
			if err == nil {
				mapOpenTrades(base, counter, s)
				// React to this now
				matchOpenTrades()
			}
		}
	}()
}

func mapCoins(ticker coinmarketcap.Ticker) {
	fmt.Println("Mapping coins")
	for _, coin := range ticker.Coins {
		p := 1 / coin.PriceUsd * vivaTargetPrice
		switch coin.ID {
		case "bitcoin":
			fmt.Printf("Bitcoin price: %f, Target VIVA price: %f\n", coin.PriceUsd, p)
			priceTargets["BTC"] = price{Base: "VIVA", Counter: "BTC", Price: p}
		case "litecoin":
			fmt.Printf("Litecoin price: %f, Target VIVA price: %f\n", coin.PriceUsd, p)
			priceTargets["LTC"] = price{Base: "VIVA", Counter: "LTC", Price: p}
		case "steem":
			fmt.Printf("Steem price: %f, Target VIVA price: %f\n", coin.PriceUsd, p)
			priceTargets["STEEM"] = price{Base: "VIVA", Counter: "STEEM", Price: p}

		}
	}

}

func mapOpenTrades(base string, counter string, trades *tradeqwik.OpenTrades) {
	fmt.Println("Mapping trades")
	switch counter {
	case "BTC":
		openTrades["BTC"] = trades
	case "litecoin":
		openTrades["LTC"] = trades
	case "steem":
		openTrades["STEEM"] = trades

	}

}
