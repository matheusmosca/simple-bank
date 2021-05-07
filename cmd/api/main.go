package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"simple-bank/pkg/common/configuration"
	account_usecase "simple-bank/pkg/domain/account/usecase"
	auth_service "simple-bank/pkg/domain/auth/service"
	"simple-bank/pkg/gateways/db/postgres"
	account_postgre "simple-bank/pkg/gateways/db/postgres/entries/account"
	"simple-bank/pkg/gateways/http"
)

func main() {
	// Load Config
	cfg, err := configuration.LoadConfig()
	if err != nil {
		log.Fatal("Unable to load credit-fs-api configuration")
	}

	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	err = postgres.RunMigrations(cfg.URL())
	if err != nil {
		log.Fatal(err)
	}

	accountUseCase := account_usecase.NewAccountUseCase(account_postgre.NewRepository(db))
	authService := auth_service.NewAuthService(accountUseCase)

	API := http.NewAPI(accountUseCase, authService)

	API.Start("0.0.0.0",cfg.API.Port)
}
