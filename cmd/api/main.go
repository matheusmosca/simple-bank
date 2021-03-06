package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/matheusmosca/simple-bank/pkg/common/configuration"
	account_usecase "github.com/matheusmosca/simple-bank/pkg/domain/account/usecase"
	auth_service "github.com/matheusmosca/simple-bank/pkg/domain/auth/service"
	transfer_usecase "github.com/matheusmosca/simple-bank/pkg/domain/transfer/usecase"
	"github.com/matheusmosca/simple-bank/pkg/gateways/db/postgres"
	account_postgre "github.com/matheusmosca/simple-bank/pkg/gateways/db/postgres/entries/account"
	transfer_postgre "github.com/matheusmosca/simple-bank/pkg/gateways/db/postgres/entries/transfer"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http"
)

func main() {
	// Load Config
	cfg, err := configuration.LoadConfig()
	if err != nil {
		log.Fatal("Unable to load configuration")
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

	accountRepo := account_postgre.NewRepository(db)
	transferRepo := transfer_postgre.NewRepository(db)
	accountUseCase := account_usecase.NewAccountUseCase(accountRepo)
	authService := auth_service.NewAuthService(accountUseCase)
	transferUseCase := transfer_usecase.NewTransfer(transferRepo, accountUseCase)

	API := http.NewAPI(accountUseCase, authService, transferUseCase)

	API.Start("0.0.0.0", cfg.API.Port)
}
