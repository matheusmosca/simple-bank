package transfer

import (
	"github.com/gorilla/mux"

	"github.com/matheusmosca/simple-bank/pkg/domain/transfer"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http/util/validator"
)

type Handler struct {
	useCase   transfer.UseCase
	validator *validator.StructValidator
}

func NewHandler(r *mux.Router, useCase transfer.UseCase) *Handler {
	h := &Handler{
		useCase:   useCase,
		validator: validator.New(),
	}

	r.HandleFunc("/transfers", h.PerformTransference).Methods("POST")
	r.HandleFunc("/transfers", h.List).Methods("GET")
	return h
}
