package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
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

type crypto struct {
	ID    int     `json:"id"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func Getcoinprice(coin string, currency string) cryptoValue {
	coinPrice := fmt.Sprintf(coinbaseAPI, coin, currency)
	// fmt.Println(coinPrice)
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
func exchange() {
	sqlStatement := `INSERT INTO cryptoinfo (code, name, price) VALUES ($1, $2, $3)`
	db := dbConnection()
	products := getProducts()
	_, clear_err := db.Exec(`TRUNCATE cryptoinfo`)
	if clear_err != nil {
		panic(clear_err)
	}

	for i := 0; i < len(products); i++ {
		exchange := Getcoinprice(products[i].ID, "AUD")
		if exchange.Data.Amount != "" {
			fmt.Println(products[i].ID, products[i].NAME, exchange.Data.Amount)
			_, err := db.Exec(sqlStatement, products[i].ID, products[i].NAME, exchange.Data.Amount)
			if err != nil {
				panic(err)
			}
		}
	}
}
func cryptoApi(w http.ResponseWriter, r *http.Request) {
	exchange()
	db := dbConnection()
	row, er := db.Query("SELECT* FROM cryptoinfo order by price::numeric desc limit 10")
	if er != nil {
		panic(er)
	}
	defer row.Close()

	cpts := make([]crypto, 0)

	for row.Next() {
		cpt := crypto{}

		if err := row.Scan(&cpt.ID, &cpt.Name, &cpt.Code, &cpt.Price); err != nil {
			panic(err)
		}
		cpts = append(cpts, cpt)
	}
	if err := row.Err(); err != nil {
		panic(err)
	}
	jsonData, err := json.Marshal(cpts)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	db.Close()
}

func main() {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Origin, Contect-Type, Accept"},
	}))

	router.Get("/api/crypto", cryptoApi)

	fmt.Println("Server is listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
