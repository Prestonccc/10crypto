package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// type Info struct {
// 	ID       string `json:"id"`
// 	Name     string `json:"name"`
// 	Min_Size string `json:"min_size"`
// }

// type Product struct {
// 	Data []Info `json:"data"`
// }

type Product struct {
	ID            string   `json:"id"`
	NAME          string   `json:"name"`
	Min_size      string   `json:"min_size"`
	Status        string   `json:"status"`
	Message       string   `json:"message"`
	Max_precision string   `json:"max_precision"`
	Convertible   []string `json:"convertible_to"`
	Details       struct {
		Type               string `json:"crypto"`
		Symbol             string `json:"symbol"`
		Network            int    `json:"network_confirmations"`
		Sort               int    `json:"sort_order"`
		Crypto_address     string `json:"crypto_address_link"`
		Crypto_transaction string `json:"crypto_transaction_link"`
		Payment_method     struct {
			Method string `json:"crypto"`
		}
		Group_types     []string `json:"group_types"`
		Display_name    string   `json:"display_name"`
		Processing_time int      `json:"processing_time_seconds"`
		Min_withdrawal  float32  `json:"min_withdrawal_amount"`
		Max_withdrawal  float32  `json:"max_withdrawal_amount"`
	}
}

func getProducts() []*Product {
	//res, err := http.Get("https://api.pro.coinbase.com/currencies")
	//resp, err := http.Get("https://api.coinbase.com/v2/currencies")
	resp, err := http.Get("https://api.exchange.coinbase.com/currencies")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer resp.Body.Close()

	// coinData, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// coinMap := Product{}

	// // Decode JSON into our map
	// err = json.Unmarshal([]byte(coinData), &coinMap)

	// if err != nil {
	// 	println(err)
	// }

	// return coinMap
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
