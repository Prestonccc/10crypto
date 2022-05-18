package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"

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

// type data struct {
// 	Amount   string `json:"amount"`
// 	Currency string `json:"currency"`
// }

// type cryptoValue struct {
// 	Data data `json:"data"`
// }

// type crypto struct {
// 	ID    int     `json:"id"`
// 	Code  string  `json:"code"`
// 	Name  string  `json:"name"`
// 	Price float32 `json:"price"`
// }

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
	db := dbConnection()
	// _, drop_err := db.Exec(`Drop table if exists cryptoinfo`)
	// if drop_err != nil {
	// 	panic(drop_err)
	// }
	// _, create_err := db.Exec(`Create table cryptoinfo (id serial primary key, code varchar(50) not null, name varchar(30) not null, price varchar(50) not null)`)
	// if create_err != nil {
	// 	panic(create_err)
	// }
	sqlStatement := `INSERT INTO cryptoinfo (code, name, price) VALUES ($1, $2, $3)`
	products := getProducts()
	_, clear_err := db.Exec(`TRUNCATE cryptoinfo`)
	if clear_err != nil {
		panic(clear_err)
	}

	for i := 0; i < len(products); i++ {
		exchange := Getcoinprice(products[i].ID, "AUD")
		if exchange.Data.Amount != "" {
			// fmt.Println(products[i].ID, products[i].NAME, exchange.Data.Amount)
			_, err := db.Exec(sqlStatement, products[i].NAME, products[i].ID, exchange.Data.Amount)
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

func cryptoDetail(w http.ResponseWriter, r *http.Request) {
	code := path.Base(r.URL.String())
	db := dbConnection()

	row, er := db.Query(`SELECT* FROM cryptoinfo where name=$1`, code)
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

func signupHandler(w http.ResponseWriter, r *http.Request) {
	var user UserInfo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	// fmt.Println(user.Username)
	// Create table userinfo (user_id serial primary key, username varchar(50) not null, email varchar(255) not null, password varchar(50) not null, balance float(8) not null)
	db := dbConnection()
	err = db.QueryRow(`SELECT user_id FROM userinfo WHERE username = $1 AND email = $2 `, user.Username, user.Email).Scan(&user.ID)
	if err != nil {
		fmt.Println("User does not exist!")
		sqlStatement := `INSERT INTO userinfo (username, password, email, balance) VALUES($1, $2, $3, $4);`
		_, err = db.Exec(sqlStatement, user.Username, user.Password, user.Email, user.Balance)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Registered")
	}

	row, er := db.Query("SELECT * FROM userinfo where username = $1 AND email = $2", user.Username, user.Email)
	if er != nil {
		panic(er)
	}
	defer row.Close()
	users := make([]UserInfo, 0)
	for row.Next() {
		user := UserInfo{}

		if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Balance); err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	if err := row.Err(); err != nil {
		panic(err)
	}
	jsonData, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
	fmt.Println("hello")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("r.body", r.Body)
		// decoder := json.NewDecoder(r.Body)
		// fmt.Println(decoder)
		var user LoginInfo
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(body, &user)
		if err != nil {
			panic(err)
		}
		fmt.Printf("userinfo : %+v\n", user)
		// ----------------------
		// decoder, err := ioutil.ReadAll(r.Body)
		// defer r.Body.Close()
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println(len(decoder) == 0)
		// err = json.Unmarshal(decoder, &user)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println(decoder.Decode(&user))
		db := dbConnection()
		// row, err := db.Query(`SELECT * FROM userinfo WHERE email = $1 AND password = $2`, email, password)
		row, err := db.Query(`SELECT * FROM userinfo WHERE email = $1 AND password = $2`, user.Email, user.Password)
		if err != nil {
			fmt.Println("Unauthorised!")
			json.NewEncoder(w).Encode("Unauthorised Login!")
			return
		}
		fmt.Println("Logged In!")

		// u := UserInfo{ID: user.ID, Username: user.Username, Password: user.Password, Email: user.Email}
		// fmt.Println(u)
		// w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(u)

		users := make([]UserInfo, 0)
		for row.Next() {
			u := UserInfo{}

			if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.Balance); err != nil {
				panic(err)
			}
			users = append(users, u)
		}
		if err := row.Err(); err != nil {
			panic(err)
		}
		jsonData, err := json.Marshal(users)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(jsonData))
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(users)

		db.Close()
	}

}

func topupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var topup Topup
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(body, &topup)
		if err != nil {
			panic(err)
		}
		fmt.Printf("topupinfo : %+v\n", topup)
		fmt.Printf("topupinfo : %+v\n", topup.Topup)

		db := dbConnection()
		row, err := db.Query(`SELECT balance FROM userinfo WHERE username = $1 AND email = $2`, topup.Username, topup.Email)
		if err != nil {
			panic(err)
		}
		var tempnumber float32
		for row.Next() {
			if err := row.Scan(&tempnumber); err != nil {
				panic(err)
			}
		}
		inputnumber, err := strconv.ParseFloat(topup.Topup, 32)
		if err != nil {
			panic(err)
		}
		fmt.Println(tempnumber)
		if err != nil {
			fmt.Println("Unauthorised!")
			json.NewEncoder(w).Encode("Unauthorised topup!")
			return
		} else {
			sqlStatement := `UPDATE userinfo SET balance = $1 WHERE username = $2 AND email = $3`
			_, err = db.Exec(sqlStatement, tempnumber+float32(inputnumber), topup.Username, topup.Email)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		row, err = db.Query(`SELECT username, email, balance FROM userinfo WHERE username = $1 AND email = $2`, topup.Username, topup.Email)
		if err != nil {
			panic(err)
		}

		tp := make([]Topup, 0)
		for row.Next() {
			t := Topup{}

			if err := row.Scan(&t.Username, &t.Email, &t.Topup); err != nil {
				panic(err)
			}
			tp = append(tp, t)
		}
		if err := row.Err(); err != nil {
			panic(err)
		}
		jsonData, err := json.Marshal(tp)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(jsonData))
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(tp)

		db.Close()

	}
}

func main() {
	router := chi.NewRouter()
	router.Use(cors.AllowAll().Handler)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/api/crypto", cryptoApi)
	router.Get("/api/crypto/{code}", cryptoDetail)
	router.Post("/api/crypto/signup", signupHandler)
	router.Post("/api/crypto/signin", signinHandler)
	router.Post("/api/crypto/topup", topupHandler)
	fmt.Println("Server is listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
