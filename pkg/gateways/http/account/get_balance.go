package account

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/matheusmosca/simple-bank/pkg/domain/account"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http/util/response"
)

func (h Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountID := mux.Vars(r)["account_id"]

	acc, err := h.usecase.GetByID(r.Context(), accountID)
	if err != nil {
		if account.IsDomainError(err) {
			_ = response.Send(
				w,
				response.ErrorResponse{Message: err.Error()},
				http.StatusNotFound,
			)
			return
		}

		_ = response.SendError(w, response.ErrIntervalServer, http.StatusInternalServerError)
		return
	}

	_ = response.Send(
		w,
		BalanceResponse{Balance: acc.Balance},
		http.StatusOK,
	)
}
