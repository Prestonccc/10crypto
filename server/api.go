package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)

// coinbase api get coin price
const coinbaseAPI = "https://api.coinbase.com/v2/prices/%s-%s/spot"

type data struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type cryptoValue struct {
	Data data `json:"data"`
}

func Getcoinprice(coin string, currency string) cryptoValue {
	coinPrice := fmt.Sprintf(coinbaseAPI, coin, currency)
	fmt.Println(coinPrice)
	resp, err := http.Get(coinPrice)
	if err != nil {
		log.Fatalln(err)
	}
	var cResp cryptoValue
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}
	return cResp
}

func homepage(w http.ResponseWriter, r *http.Request) {

	//sqlStatement := `INSERT INTO cryptoinfo (name, base_currency, quote_currency) VALUES ($1, $2, $3)`
	sqlStatement := `INSERT INTO cryptoinfo (code, name, price) VALUES ($1, $2, $3)`
	db := dbConnection()
	products := getProducts()
	_, clear_err := db.Exec(`TRUNCATE cryptoinfo`)
	if clear_err != nil {
		panic(clear_err)
	}

	for i := 0; i < len(products); i++ {
		//_, err := db.Exec(sqlStatement, products[i].ID, products[i].BaseCurrency, products[i].QuoteCurrency)

		exchange := Getcoinprice(products[i].ID, "AUD")
		if exchange.Data.Amount != "" {
			fmt.Println(products[i].ID, products[i].NAME, exchange.Data.Amount)
			_, err := db.Exec(sqlStatement, products[i].ID, products[i].NAME, exchange.Data.Amount)
			if err != nil {
				panic(err)
			}
		}

	}

	json.NewEncoder(w).Encode("Top 10 Crypto List in AUD")

	row, er := db.Query("SELECT* FROM cryptoinfo order by price::numeric desc limit 10")
	if er != nil {
		panic(er)
	}
	for row.Next() {
		var (
			name  string
			code  string
			price float32
		)
		if err := row.Scan(&name, &code, &price); err != nil {
			panic(err)
		}
		fmt.Printf("%s, %s", name, code)
		//fmt.Sprintf("%f",price)
		json.NewEncoder(w).Encode(name + ", " + code + ", " + fmt.Sprintf("%f", price))
	}
	if err := row.Err(); err != nil {
		panic(err)
	}
}

func main() {
	// bitcoin := Getcoinprice("BTC", "AUD")
	// fmt.Println(reflect.TypeOf(bitcoin.Data.Amount))
	// etheur := Getcoinprice("ETH", "AUD")
	// fmt.Printf("%s: %s\n", etheur.Data.Currency, etheur.Data.Amount)
	router := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	// http.ListenAndServe(":3000", r)
	router.Get("/home", homepage)
	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
