package http

import (
	"fmt"
	"log"
	"net/http"
	"simple-bank/pkg/domain/account"
	"simple-bank/pkg/domain/auth"
	"simple-bank/pkg/domain/transfer"
	"time"

	account_handler "simple-bank/pkg/gateways/http/account"
	auth_handler "simple-bank/pkg/gateways/http/auth"
	"simple-bank/pkg/gateways/http/middlewares"
	transfer_handler "simple-bank/pkg/gateways/http/transfer"

	"github.com/gorilla/mux"
)

type API struct {
	AccountUseCase  account.UseCase
	AuthService     auth.Service
	TransferUseCase transfer.UseCase
}

func NewAPI(
	accUseCase account.UseCase,
	authService auth.Service,
	transferUseCase transfer.UseCase) *API {
	return &API{
		AccountUseCase:  accUseCase,
		AuthService:     authService,
		TransferUseCase: transferUseCase,
	}
}

func (a API) Start(host string, port string) {
	router := mux.NewRouter()

	publicV1 := router.PathPrefix("/api/v1").Subrouter()
	AuthV1 := router.PathPrefix("/api/v1").Subrouter()

	account_handler.NewHandler(publicV1, a.AccountUseCase)
	auth_handler.NewHandler(publicV1, a.AuthService)
	transfer_handler.NewHandler(AuthV1, a.TransferUseCase)

	AuthV1.Use(middlewares.Authorize)

	endpoint := fmt.Sprintf("%s:%s", host, port)

	srv := &http.Server{
		Handler:      router,
		Addr:         endpoint,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting api...")
	log.Fatal(srv.ListenAndServe())
}
