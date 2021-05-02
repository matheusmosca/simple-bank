package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "simple_bank_db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "simple_bank_db"
)

func main() {
	psqlConfig := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	// accountUseCase := account_usecase.NewAccountUseCase(account_postgree.NewRepository(db))

	// acc, err := accountUseCase.Create(context.Background(), "maria", "738.333.222-12", "2222222222", 0)
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Println(acc)

	r := mux.NewRouter()

	r.HandleFunc("/api/v1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	http.ListenAndServe(":3000", r)
}
