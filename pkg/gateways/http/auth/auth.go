package auth

import (
	"github.com/gorilla/mux"

	"simple-bank/pkg/domain/auth"
	"simple-bank/pkg/gateways/http/util/validator"
)

type Handler struct {
	authService auth.Service
	validator   *validator.StructValidator
}

func NewHandler(r *mux.Router, auth auth.Service) *Handler {
	h := &Handler{
		authService: auth,
		validator:   validator.New(),
	}

	r.HandleFunc("/login", h.Login).Methods("POST")

	return h
}
