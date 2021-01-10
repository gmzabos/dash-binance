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

	// Define PriceStruct input
	type PriceStruct struct {
		Symbol string  `json:"symbol"`
		Price  float64 `json:"price,string"`
	}

	var price PriceStruct
	json.Unmarshal([]byte(bodyString), &price)

	// Calculate percentages
	// to-do: range calculation instead of singles
	p101 := (price.Price) / 100 * 101
	p102 := (price.Price) / 100 * 102
	p103 := (price.Price) / 100 * 103
	p104 := (price.Price) / 100 * 104
	p105 := (price.Price) / 100 * 105

	// Output
	// to-do: range output
	// to-do: pretty output
	fmt.Printf("\nTrading pair: %s\n", price.Symbol)
	fmt.Println("-------------")
	fmt.Println("---> (5%)", p105)
	fmt.Println("---> (4%)", p104)
	fmt.Println("---> (3%)", p103)
	fmt.Println("---> (2%)", p102)
	fmt.Println("---> (1%)", p101)
	fmt.Println("---> ....", price.Price)
}
