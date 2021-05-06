package auth

type (
	LoginRequest struct {
		CPF    string `json:"cpf" validate:"required"`
		Secret string `json:"secret" validate:"required"`
	}
	LoginResponse struct {
		Token string `json:"token"`
	}
	ValidationErrorResponse struct {
		CPF    string `json:"cpf,omitempty"`
		Secret string `json:"secret,omitempty"`
	}
)
