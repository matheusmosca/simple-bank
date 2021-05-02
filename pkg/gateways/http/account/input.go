package account

import "time"

type (
	CreateBodyRequest struct {
		Name   string `json:"name" validate:"required"`
		CPF    string `json:"cpf" validate:"required"`
		Secret string `json:"secret" validate:"required"`
	}

	ResponseBody struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		CPF       string    `json:"cpf"`
		Balance   int       `json:"balance"`
		CreatedAt time.Time `json:"created_at"`
	}

	ValidationErrorResponse struct {
		Name   string `json:"name,omitempty"`
		CPF    string `json:"cpf,omitempty"`
		Secret string `json:"secret,omitempty"`
	}
)
