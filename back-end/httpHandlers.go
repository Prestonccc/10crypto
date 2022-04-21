package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getProducts() []*Product {
	//res, err := http.Get("https://api.pro.coinbase.com/currencies")
	//resp, err := http.Get("https://api.coinbase.com/v2/currencies")
	resp, err := http.Get("https://api.exchange.coinbase.com/currencies")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var products []*Product
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&products)
	if err != nil {
		panic(err)
	}
	return products
}
