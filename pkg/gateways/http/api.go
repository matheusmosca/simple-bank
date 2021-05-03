package http

import (
	"fmt"
	"net/http"
	"simple-bank/pkg/domain/account"

	account_handler "simple-bank/pkg/gateways/http/account"

	"github.com/gorilla/mux"
)

type API struct {
	//TODO add more usecases
	AccountUseCase account.UseCase
}

func NewAPI(accUseCase account.UseCase) *API {
	return &API{
		AccountUseCase: accUseCase,
	}
}

func (a API) Start() {
	router := mux.NewRouter()

	v1 := router.PathPrefix("/api/v1").Subrouter()

	account_handler.NewHandler(v1, a.AccountUseCase)

	fmt.Println("Starting api...")
	http.ListenAndServe(":3000", v1)
}
