package tradeqwik

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// GetRecentTrades makes the actual http request, parses the JSON and returns
// the data in a struct
func GetRecentTrades(base string, counter string) (*RecentTrades, error) {
	resp, err := http.Get(BaseURL + "/api/1/recent_trades/" + base + "/" + counter)
	if err != nil {

		return &RecentTrades{Base: base, Counter: counter}, err
	}

	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return &RecentTrades{Base: base, Counter: counter}, err
	}

	var j []*RecentTrade
	if err := json.Unmarshal(ret, &j); err != nil {
		log.Fatal(err)
		return &RecentTrades{Base: base, Counter: counter}, err
	}

	r := RecentTrades{Trades: j, Base: base, Counter: counter, LastUpdate: time.Now()}

	return &r, err
}

// GetOpenTrades makes the actual http request, parses the JSON and returns
// the data in a struct
func GetOpenTrades(base string, counter string) (*OpenTrades, error) {
	resp, err := http.Get(BaseURL + "/api/1/open_trades/" + base + "/" + counter)
	if err != nil {
		return &OpenTrades{Base: base, Counter: counter}, err
	}

	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return &OpenTrades{Base: base, Counter: counter}, err
	}

	var j *OpenTrades
	if err := json.Unmarshal(ret, &j); err != nil {
		log.Fatal(err)
		return &OpenTrades{Base: base, Counter: counter}, err
	}

	j.Base = base
	j.Counter = counter
	j.LastUpdate = time.Now()
	return j, err
}
