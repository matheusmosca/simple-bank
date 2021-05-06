package auth

import (
	"net/http"
	"simple-bank/pkg/gateways/http/util/response"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqBody LoginRequest
	err := response.Decode(r, &reqBody)
	if err != nil {
		response.SendError(w, response.ErrDecode, http.StatusBadRequest)
		return
	}

	var validationErr ValidationErrorResponse
	err = h.validator.Validate(reqBody, &validationErr)
	if err != nil {
		response.Send(w, validationErr, http.StatusBadRequest)
		return
	}

	token, err := h.authService.Authenticate(r.Context(), reqBody.CPF, reqBody.Secret)
	if err != nil {
		response.SendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response.Send(
		w,
		LoginResponse{Token: token},
		http.StatusOK,
	)
}
