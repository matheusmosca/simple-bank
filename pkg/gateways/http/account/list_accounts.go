package account

import (
	"net/http"

	"github.com/matheusmosca/simple-bank/pkg/gateways/http/util/response"
)

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.usecase.List(r.Context())
	if err != nil {
		_ = response.SendError(w, response.ErrIntervalServer, http.StatusInternalServerError)
		return
	}

	res := make([]ResponseBody, len(accounts))

	for i, acc := range accounts {
		res[i] = ResponseBody{
			ID:        acc.ID,
			Name:      acc.Name,
			CPF:       acc.CPF,
			Balance:   acc.Balance,
			CreatedAt: acc.CreatedAt,
		}
	}

	_ = response.Send(w, res, http.StatusOK)
}
