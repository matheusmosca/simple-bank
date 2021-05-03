package account

import (
	"net/http"

	"simple-bank/pkg/domain/account"
	"simple-bank/pkg/gateways/http/util/response"
)

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateBodyRequest

	err := response.Decode(r, &reqBody)
	if err != nil {
		response.SendError(w, response.ErrDecode, http.StatusBadRequest)
		return
	}

	var validationErrPayload ValidationErrorResponse
	err = h.Validator.Validate(reqBody, &validationErrPayload)

	if err != nil {
		response.Send(
			w,
			validationErrPayload,
			http.StatusBadRequest,
		)
		return
	}

	acc, err := h.UseCase.Create(r.Context(), reqBody.Name, reqBody.CPF, reqBody.Secret)
	if err != nil {
		if account.IsDomainError(err) {
			response.SendError(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.SendError(w, response.ErrIntervalServer, http.StatusInternalServerError)
		return
	}

	response.Send(w, ResponseBody{
		ID:        acc.ID,
		Name:      acc.Name,
		CPF:       acc.CPF,
		Balance:   acc.Balance,
		CreatedAt: acc.CreatedAt,
	}, http.StatusCreated)
}