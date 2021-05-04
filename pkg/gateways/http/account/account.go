package account

import (
	"simple-bank/pkg/domain/account"
	"simple-bank/pkg/gateways/http/util/validator"

	"github.com/gorilla/mux"
)

type Handler struct {
	usecase   account.UseCase
	validator *validator.StructValidator
}

func NewHandler(r *mux.Router, usecase account.UseCase) *Handler {
	h := &Handler{
		usecase:   usecase,
		validator: validator.New(),
	}

	// TODO implement handler methods
	r.HandleFunc("/accounts", h.Create).Methods("POST")
	r.HandleFunc("/accounts", h.List).Methods("GET")
	r.HandleFunc("/accounts/{account_id}/balance", h.GetBalance).Methods("GET")

	return h
}
