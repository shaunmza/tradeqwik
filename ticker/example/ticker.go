package main

import (
	"fmt"
	"strconv"

	"github.com/shaunmza/tradeqwik"
	"github.com/shaunmza/tradeqwik/ticker"
)

func main() {

	var r tradeqwik.Ticker

	// Initialse, so we can get the channel to receive updates from
	tChan := ticker.Init()

	// Infinite loop so we keep getting ticker info
	for {
		// Get off of the channel
		r = <-tChan

		// If this is not nil then we encountered a problem, use this to determine
		// what to do next.
		// LastUpdate can be used to determine how stale the data is
		if r.Error != nil {
			fmt.Printf("Error! %s, Last Updated: %s\n", r.Error, r.LastUpdate)
		}

		movement := " - "
		// Just print it out for now
		for _, pair := range r.CurrencyPairs {
			if pair.Last > pair.First {
				movement = " \033[32m^\033[0m "
			}
			if pair.Last < pair.First {
				movement = " \033[31mv\033[0m "
			}

			fmt.Println(movement + " " + strconv.FormatFloat(pair.Last, 'f', 8, 64) + " " + pair.Base + "/" + pair.Counter)
		}
		fmt.Println("=======================================================")
	}

}
