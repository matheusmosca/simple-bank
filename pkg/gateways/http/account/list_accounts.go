package account

import (
	"net/http"
	"simple-bank/pkg/domain/account"
	"simple-bank/pkg/gateways/http/util/response"
)

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.usecase.List(r.Context())
	if err != nil {
		if account.IsDomainError(err) {
			response.SendError(w, err.Error(), http.StatusNotFound)
			return
		}
		response.SendError(w, response.ErrIntervalServer, http.StatusInternalServerError)
		return
	}

	res := make([]ResponseBody, len(accounts))

	for i, acc := range accounts {
		res[i] = ResponseBody{
			ID:        acc.ID,
			CPF:       acc.CPF,
			Balance:   acc.DisplayBalance(),
			CreatedAt: acc.CreatedAt,
		}
	}

	response.Send(w, res, http.StatusOK)
}
