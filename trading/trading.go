package trading

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/shaunmza/tradeqwik"
)

type order struct {
	APIKey  string  `json:"api_key"`
	Base    string  `json:"base"`
	Counter string  `json:"counter"`
	Amount  float64 `json:"amount,string"`
	Price   float64 `json:"price,string"`
}

type cancel struct {
	APIKey  string `json:"api_key"`
	OrderID int64  `json:"order"`
}

type key struct {
	APIKey string `json:"api_key"`
}

type callSuccess struct {
	Message string  `json:"message"`
	OrderID int64   `json:"order"`
	Remains float64 `json:"remains"`
}

type Balances struct {
	Currencies map[string]*balance
	LastUpdate time.Time
}

type balance struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount,string"`
	Hold     float64 `json:"hold,string"`
}

var apiKey string

func Init(key string) {
	apiKey = key
}

func Buy(base string, counter string, amount float64, price float64) (*callSuccess, error) {
	endpoint := "/api/1/bid"
	data := order{APIKey: apiKey, Base: base, Counter: counter, Amount: amount, Price: price}

	resp, err := postData(data, endpoint)
	if err != nil {
		return &callSuccess{}, err
	}

	s := []byte(resp)

	var j *callSuccess
	if err := json.Unmarshal(s, &j); err != nil {
		return &callSuccess{}, err
	}

	if j.Message != "" {
		err = errors.New(j.Message)
		return &callSuccess{}, err
	}
	return j, err
}

func Sell(base string, counter string, amount float64, price float64) (*callSuccess, error) {
	endpoint := "/api/1/ask"
	data := order{APIKey: apiKey, Base: base, Counter: counter, Amount: amount, Price: price}

	resp, err := postData(data, endpoint)
	if err != nil {
		return &callSuccess{}, err
	}

	s := []byte(resp)

	var j *callSuccess
	if err := json.Unmarshal(s, &j); err != nil {
		return &callSuccess{}, err
	}

	if j.Message != "" {
		err = errors.New(j.Message)
		return &callSuccess{}, err
	}
	return j, err
}

func Cancel(orderId int64) (bool, error) {
	endpoint := "/api/1/cancel"
	data := cancel{APIKey: apiKey, OrderID: orderId}

	resp, err := postData(data, endpoint)
	if err != nil {
		return false, err
	}

	s := []byte(resp)

	var j *callSuccess
	if err := json.Unmarshal(s, &j); err != nil {
		return false, err
	}

	if j.Message != "" {
		err = errors.New(j.Message)
		return false, err
	}
	return true, err
}

func GetPending() (tradeqwik.AccountOpenTrades, error) {
	endpoint := "/api/1/pending_trades"
	data := key{APIKey: apiKey}

	resp, err := postData(data, endpoint)
	if err != nil {
		return tradeqwik.AccountOpenTrades{}, err
	}

	s := []byte(resp)

	var j []*tradeqwik.AccountOpenTrade
	if err := json.Unmarshal(s, &j); err != nil {
		return tradeqwik.AccountOpenTrades{}, err
	}

	ret := tradeqwik.AccountOpenTrades{Trades: j, LastUpdate: time.Now()}
	return ret, err
}

func GetHistory(data tradeqwik.AccountHistory) (tradeqwik.RecentTrades, error) {

	endpoint := "/api/1/trade_history"
	data.APIKey = apiKey

	resp, err := postData(data, endpoint)
	if err != nil {
		return tradeqwik.RecentTrades{}, err
	}

	s := []byte(resp)

	var j []*tradeqwik.RecentTrade
	if err := json.Unmarshal(s, &j); err != nil {
		return tradeqwik.RecentTrades{}, err
	}

	ret := tradeqwik.RecentTrades{Trades: j, LastUpdate: time.Now()}
	return ret, err
}

func postData(data interface{}, endpoint string) (string, error) {
	url := tradeqwik.BaseURL + endpoint

	j, err := json.Marshal(data)
	if err != nil {
		return "{}", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return "{}", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "{}", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), err
}

func GetBalance() (Balances, error) {
	endpoint := "/api/1/balance"
	data := key{APIKey: apiKey}

	resp, err := postData(data, endpoint)
	if err != nil {
		return Balances{}, err
	}

	s := []byte(resp)

	var j []*balance
	if err := json.Unmarshal(s, &j); err != nil {
		return Balances{}, err
	}

	var m = make(map[string]*balance, len(j))
	for _, b := range j {
		m[b.Currency] = b
	}

	ret := Balances{Currencies: m, LastUpdate: time.Now()}
	return ret, err
}
