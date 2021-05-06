package http

import (
	"fmt"
	"log"
	"net/http"
	"simple-bank/pkg/domain/account"
	"simple-bank/pkg/domain/auth"

	account_handler "simple-bank/pkg/gateways/http/account"
	auth_handler "simple-bank/pkg/gateways/http/auth"

	"github.com/gorilla/mux"
)

type API struct {
	//TODO add more usecases
	AccountUseCase account.UseCase
	AuthService    auth.Service
}

func NewAPI(accUseCase account.UseCase, authService auth.Service) *API {
	return &API{
		AccountUseCase: accUseCase,
		AuthService:    authService,
	}
}

func (a API) Start() {
	router := mux.NewRouter()

	v1 := router.PathPrefix("/api/v1").Subrouter()

	account_handler.NewHandler(v1, a.AccountUseCase)
	auth_handler.NewHandler(v1, a.AuthService)

	fmt.Println("Starting api...")
	err := http.ListenAndServe(":3000", v1)
	log.Println(err)
}
