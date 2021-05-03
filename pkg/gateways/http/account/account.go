package account

import (
	"simple-bank/pkg/domain/account"
	"simple-bank/pkg/gateways/http/util/validator"

	"github.com/gorilla/mux"
)

type Handler struct {
	UseCase   account.UseCase
	Validator *validator.StructValidator
}

func NewHandler(r *mux.Router, useCase account.UseCase) *Handler {
	h := &Handler{
		UseCase:   useCase,
		Validator: validator.New(),
	}

	// TODO implement handler methods
	r.HandleFunc("/accounts", h.Create).Methods("POST")
	r.HandleFunc("/accounts", h.List).Methods("GET")
	// r.HandleFunc("/accounts/{id}/balance", h.GetBalance).Methods("GET")

	return h
}
