package main

import (
	"fmt"
	"strconv"

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
	trades := trading.GetPending()

	for _, trade := range trades.Trades {
		fmt.Println(strconv.FormatInt(trade.ID, 10) + " Type: " + trade.Type + " Price" + strconv.FormatFloat(trade.Price, 'f', 8, 64))
	}
	fmt.Println("=======================================================")
	/*hstruct := tradeqwik.AccountHistory{}

	history := trading.GetHistory(hstruct)

	for _, trade := range history.Trades {
		fmt.Println(" Type: " + trade.Type + " Price" + strconv.FormatFloat(trade.Price, 'f', 8, 64))
	}
	fmt.Println("=======================================================")

	trading.Cancel(4959)*/

}
