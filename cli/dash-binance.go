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
	getSymbol()
}

// GET function
// API Query for BNBEUR
func getSymbol() {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=BNBEUR")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	// Define price structure input
	type BnbPrice struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	}

	var price BnbPrice
	json.Unmarshal([]byte(bodyString), &price)
	fmt.Printf("Symbol: %s, Price: %s\n", price.Symbol, price.Price)
}
