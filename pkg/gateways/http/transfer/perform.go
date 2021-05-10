package transfer

import (
	"net/http"

	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
	"github.com/matheusmosca/simple-bank/pkg/domain/transfer"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http/middlewares"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http/util/response"
)

func (h Handler) PerformTransference(w http.ResponseWriter, r *http.Request) {
	var reqBody PerformRequest

	authAccountID, ok := middlewares.GetAccountID(r.Context())
	if !ok || authAccountID == "" {
		response.SendError(w, response.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	err := response.Decode(r, &reqBody)
	if err != nil {
		response.SendError(w, response.ErrDecode, http.StatusBadRequest)
		return
	}

	var valErr ValidationErrResponse
	err = h.validator.Validate(reqBody, &valErr)
	if err != nil {
		response.Send(w, valErr, http.StatusBadRequest)
		return
	}

	trans, err := h.useCase.Perform(r.Context(), entities.CreateTransferInput{
		AccountOriginID:      authAccountID,
		AccountDestinationID: reqBody.AccountDestinationID,
		Amount:               reqBody.Amount,
	})

	if err != nil {
		if transfer.IsDomainError(err) {
			response.Send(w, response.ErrorResponse{Message: err.Error()}, http.StatusBadRequest)
			return
		}
		response.SendError(w, response.ErrIntervalServer, http.StatusInternalServerError)
		return
	}

	response.Send(w, ResponseBody{
		ID:                   trans.ID,
		AccountOriginID:      trans.AccountOriginID,
		AccountDestinationID: trans.AccountDestinationID,
		Amount:               trans.Amount,
		CreatedAt:            trans.CreatedAt,
	}, http.StatusCreated)
}
