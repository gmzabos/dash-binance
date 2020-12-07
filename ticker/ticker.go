package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// This is all of our functions
func main() {
	getBnbEur()
	getXlmBnb()
	getUnfiBnb()
}

// GET function
// API Query for BNBEUR
func getBnbEur() {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=BNBEUR")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

// Define price structure input
type BnbPrice struct {
	Symbol 	string `json:"symbol"`
	Price 	string `json:"price"`
}

var price BnbPrice
json.Unmarshal([]byte(bodyString), &price)
fmt.Printf("Symbol: %s, Price: %s\n", price.Symbol, price.Price)
}

// GET function
// API Query for XLMBNB
func getXlmBnb() {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=XLMBNB")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

// Define price structure input
type XlmPrice struct {
	Symbol 	string `json:"symbol"`
	Price 	string `json:"price"`
}

var price XlmPrice
json.Unmarshal([]byte(bodyString), &price)
fmt.Printf("Symbol: %s, Price: %s\n", price.Symbol, price.Price)
}


// GET function
// API Query for UNFIBNB
func getUnfiBnb() {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=UNFIBNB")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

// Define price structure input
type UnfiPrice struct {
	Symbol 	string `json:"symbol"`
	Price 	string `json:"price"`
}

var price UnfiPrice
json.Unmarshal([]byte(bodyString), &price)
fmt.Printf("Symbol: %s, Price: %s\n", price.Symbol, price.Price)
}