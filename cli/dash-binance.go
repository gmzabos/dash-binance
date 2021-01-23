package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// All functions
func main() {
	GetSymbol()
}

// API Query for trade pair symbol
func GetSymbol() {

	// Set and explain flag(s)
	tpPtr := flag.String("tp", "", "Trade pair symbol to query. (Required)")
	flag.Parse()

	// Show usage
	// Exit if no flag given
	if *tpPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Construct URL for API call
	tpSymbol := "symbol=" + url.QueryEscape(*tpPtr)
	queryURL := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?%s", tpSymbol)

	// Make the API call
	response, err := http.Get(queryURL)
	if err != nil {
		log.Fatal(err)
	}

	// Close when finished
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
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
	fmt.Println("-----------")
	fmt.Printf("Trade pair: %s\n", price.Symbol)
	fmt.Println("-----------")
	fmt.Println("---> (...5%)", p105)
	fmt.Println("---> (...4%)", p104)
	fmt.Println("---> (...3%)", p103)
	fmt.Println("---> (...2%)", p102)
	fmt.Println("---> (...1%)", p101)
	fmt.Println("---> (PRICE)", price.Price)

}
