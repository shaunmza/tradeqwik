package main

import (
	"fmt"
	"strconv"

	"github.com/shaunmza/tradeqwik"
	"github.com/shaunmza/tradeqwik/trading"
)

func main() {

	trading.Init("YOUR KEY")

	/*success, err := trading.Sell("VIVA", "USD", 0.05, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(success)*/

	balances, err := trading.GetBalance()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(balances.Currencies["VIVA"])
	trades, err := trading.GetPending()
	if err != nil {
		fmt.Println(err)
	}

	for _, trade := range trades.Trades {
		fmt.Printf("Base: %s Counter: %s Id: %f Type: %s Price:%f \n", trade.Base, trade.Counter, trade.ID, trade.Type, trade.Price)
	}
	fmt.Println("=======================================================")
	hstruct := tradeqwik.AccountHistory{}

	history, err := trading.GetHistory(hstruct)
	if err != nil {
		fmt.Println(err)
	}

	for _, trade := range history.Trades {
		fmt.Println(" Type: " + trade.Type + " Price" + strconv.FormatFloat(trade.Price, 'f', 8, 64))
	}
	fmt.Println("=======================================================")

	/*trading.Cancel(4959)*/

}
