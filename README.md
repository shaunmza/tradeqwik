# Trade Qwik API client library

Simple API client library for the Trade Qwik API, written in Go

## Installation

This is not a standalone app, use it in your project as a package. To get it simply run;
`go get github.com/shaunmza/tradeqwik`

## Usage

Examples are provided in the following directories;
`examples/orders` - Get open trades and recent trades
`ticker/example` - Periodically fetch ticker data
`trading/example` - Buy, sell or cancel orders, get your recent trade history, get your open trades - much of this is implemented in the marketmaker example too

`examples/marketmaker` - a bot that attempts to close orders within a 10% range of the target VIVA USD value of 5.5 (currently)
This depends on another library, run `go get github.com/shaunmza/coinmarketcap` to get it.
You will also need to provide your API key, edit `examples/marketmaker/bot.go` and put your API key where you see this line;
```golang
trading.Init("YOUR API KEY")
```

## Authors

* **Shaun Morrow** - *Initial work* - [shaunmza](https://github.com/shaunmza)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
