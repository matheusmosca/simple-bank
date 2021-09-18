package transfer

import (
	"net/http"

	"github.com/matheusmosca/simple-bank/pkg/domain/transfer"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http/middlewares"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http/util/response"
)

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	accountID, ok := middlewares.GetAccountID(r.Context())
	if !ok || accountID == "" {
		_ = response.SendError(w, response.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	transfers, err := h.useCase.List(r.Context(), accountID)
	if err != nil {
		if transfer.IsDomainError(err) {
			_ = response.SendError(w, err.Error(), http.StatusBadRequest)
			return
		}
		_ = response.SendError(w, response.ErrIntervalServer, http.StatusInternalServerError)
		return
	}

	transfersResponse := formatSliceResponse(transfers)
	_ = response.Send(w, transfersResponse, http.StatusOK)
}
