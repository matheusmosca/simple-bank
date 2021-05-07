package http

import (
	"fmt"
	"log"
	"net/http"
	"simple-bank/pkg/domain/account"
	"simple-bank/pkg/domain/auth"
	"time"

	account_handler "simple-bank/pkg/gateways/http/account"
	auth_handler "simple-bank/pkg/gateways/http/auth"

	"github.com/gorilla/mux"
)

type API struct {
	AccountUseCase account.UseCase
	AuthService    auth.Service
}

func NewAPI(accUseCase account.UseCase, authService auth.Service) *API {
	return &API{
		AccountUseCase: accUseCase,
		AuthService:    authService,
	}
}

func (a API) Start(host string, port string) {
	router := mux.NewRouter()

	v1 := router.PathPrefix("/api/v1").Subrouter()

	account_handler.NewHandler(v1, a.AccountUseCase)
	auth_handler.NewHandler(v1, a.AuthService)

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
