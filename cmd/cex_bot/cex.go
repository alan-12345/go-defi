package main

import (
	"encoding/json"
	"fmt"
	"go_defi/networks/binance"
	"go_defi/utils"
	"io/ioutil"
	"log"
	"net/http"
)

type CexPrice struct {
	Symbol   string
	Base     string
	Quote    string
	BidPrice string
	BidQty   string
	AskPrice string
	AskQty   string
	Time     int
}

func fetch_cex_prices(prices *[]CexPrice) error {
	resp, err := http.Get("https://fapi.binance.com/fapi/v1/ticker/bookTicker")
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var raw_prices []CexPrice
	json.Unmarshal(body, &raw_prices)

	var fmt_prices []CexPrice
	for _, pair := range raw_prices {
		if val, ok := binance_tokens.TOKENS[pair.Symbol]; ok {
			pair.Base = val.Base
			pair.Quote = val.Quote
			fmt_prices = append(fmt_prices, pair)
		}
	}

	*prices = fmt_prices

	return nil
}

func run_bot() error {
	var prices []CexPrice
	if err := fetch_cex_prices(&prices); err != nil {
		return err
	}

	fmt.Printf("%+v", prices)
	return nil
}

func start_bot() {
	utils.PrintDashed()
	log.Println("Running cex_bot")
	utils.PrintDashed()

	if err := run_bot(); err != nil {
		log.Println("something wrong with running the bot", err)
	}
}
