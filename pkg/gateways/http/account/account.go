package account

import (
	"github.com/gorilla/mux"

	"github.com/matheusmosca/simple-bank/pkg/domain/account"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http/util/validator"
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
