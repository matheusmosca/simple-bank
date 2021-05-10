package transfer

import (
	"log"
	"net/http"

	"github.com/matheusmosca/simple-bank/pkg/domain/transfer"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http/middlewares"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http/util/response"
)

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	accountID, ok := middlewares.GetAccountID(r.Context())
	if !ok || accountID == "" {
		response.SendError(w, response.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	transfers, err := h.useCase.List(r.Context(), accountID)
	log.Println(transfers)
	if err != nil {
		if transfer.IsDomainError(err) {
			response.SendError(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.SendError(w, response.ErrIntervalServer, http.StatusInternalServerError)
		return
	}

	transfersResponse := formatSliceResponse(transfers)
	response.Send(w, transfersResponse, http.StatusOK)
}
