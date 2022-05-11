package main

//Coin id source models, used only to get coin id and name
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

type data struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type cryptoValue struct {
	Data data `json:"data"`
}

type crypto struct {
	ID    int     `json:"id"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type UserInfo struct {
	ID       string  `json:"user_id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Balance  float32 `json:"balance"`
}

type LoginInfo struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Balance  float32 `json:"balance"`
}
