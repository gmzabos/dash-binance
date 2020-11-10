package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Define price structure input
type Price []struct {
	Symbol 	string `json:"symbol"`
	Price 	string `json:"price"`
}

// This is all of our functions
func main() {
	getBnbEur()
	getEthEur()
	getBnbEth()
	getXlmEth()
}

// GET function, simple API Query for BNBEUR
func getBnbEur() {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=BNBEUR")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var PriceStruct Price
	json.Unmarshal(bodyBytes, &PriceStruct)
}

// GET function, simple API Query for ETHEUR
func getEthEur() {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=ETHEUR")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var PriceStruct Price
	json.Unmarshal(bodyBytes, &PriceStruct)
}


// GET function, simple API Query for BNBETH
func getBnbEth() {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=BNBETH")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var PriceStruct Price
	json.Unmarshal(bodyBytes, &PriceStruct)
}

// GET function, simple API Query for BNBEUR
func getXlmEth() {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=XLMETH")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var PriceStruct Price
	json.Unmarshal(bodyBytes, &PriceStruct)
}