package account

import (
	"simple-bank/pkg/domain/account"

	"github.com/gorilla/mux"
)

type Handler struct {
	UseCase account.UseCase
}

func NewHandler(r *mux.Router, useCase account.UseCase) *Handler {
	h := &Handler{
		UseCase: useCase,
	}

	// TODO implement handler methods
	// r.HandleFunc("/accounts", h.Create).Methods("POST")
	// r.HandleFunc("/accounts/{id}/balance", h.GetBalance).Methods("GET")
	// r.HandleFunc("/accounts", h.GetAccounts).Methods("GET")

	return h
}
