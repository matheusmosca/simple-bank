package main

import (
	"database/sql"
	"fmt"

	account_usecase "simple-bank/pkg/domain/account/usecase"
	account_postgre "simple-bank/pkg/gateways/db/postgres/entries/account"
	"simple-bank/pkg/gateways/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
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

	accountUseCase := account_usecase.NewAccountUseCase(account_postgre.NewRepository(db))

	API := http.NewAPI(accountUseCase)

	API.Start()
}
