package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	account_usecase "simple-bank/pkg/domain/account/usecase"
	auth_service "simple-bank/pkg/domain/auth/service"
	"simple-bank/pkg/gateways/db/postgres"
	account_postgre "simple-bank/pkg/gateways/db/postgres/entries/account"
	"simple-bank/pkg/gateways/http"
)

const (
	host     = "simple_bank_db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "simple_bank_db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	psqlConfig := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=%s",
		user, password, host, port, dbname, "disable")

	db, err := sql.Open("postgres", psqlConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	err = postgres.RunMigrations(connectionString)
	if err != nil {
		log.Println(err)
	}

	accountUseCase := account_usecase.NewAccountUseCase(account_postgre.NewRepository(db))
	authService := auth_service.NewAuthService(accountUseCase)

	API := http.NewAPI(accountUseCase, authService)

	API.Start()
}
