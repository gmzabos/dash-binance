package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// bookTicker structure
type BookTicker []struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

func main() {
	get()
}

func get() {
	fmt.Println("...API Query...")
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/bookTicker")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	var bookTickerStruct BookTicker
	json.Unmarshal(bodyBytes, &bookTickerStruct)
}