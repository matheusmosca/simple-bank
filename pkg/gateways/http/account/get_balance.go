package account

import (
	"net/http"
	"simple-bank/pkg/domain/account"
	"simple-bank/pkg/gateways/http/util/response"

	"github.com/gorilla/mux"
)

func (h Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountID := mux.Vars(r)["account_id"]

	acc, err := h.usecase.GetByID(r.Context(), accountID)
	if err != nil {
		if account.IsDomainError(err) {
			response.Send(
				w,
				response.ErrorResponse{Message: err.Error()},
				http.StatusNotFound,
			)
			return
		}

		response.SendError(w, response.ErrIntervalServer, http.StatusInternalServerError)
		return
	}

	response.Send(
		w,
		BalanceResponse{Balance: acc.Balance},
		http.StatusOK,
	)
}
