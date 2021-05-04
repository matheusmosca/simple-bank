package account

import (
	"net/http"
	"simple-bank/pkg/domain/account"
	"simple-bank/pkg/gateways/http/util/response"

	"github.com/gorilla/mux"
)

func (h Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountID, ok := mux.Vars(r)["account_id"]
	if !ok {
		response.SendError(w, "account_id not provided", http.StatusNotFound)
		return
	}

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
	}

	response.Send(
		w,
		BalanceResponse{Balance: acc.DisplayBalance()},
		http.StatusOK,
	)
}
