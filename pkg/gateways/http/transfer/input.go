package transfer

import (
	"simple-bank/pkg/domain/entities"
	"time"
)

type (
	PerformRequest struct {
		AccountDestinationID string `json:"account_destination_id" validate:"required"`
		Amount               int    `json:"amount" validate:"required"`
	}
	ValidationErrResponse struct {
		AccountDestinationID string `json:"account_destination_id,omitempty"`
		Amount               string `json:"amount,omitempty"`
	}
	ResponseBody struct {
		ID                   string    `json:"id"`
		AccountDestinationID string    `json:"account_destination_id"`
		AccountOriginID      string    `json:"account_origin_id"`
		Amount               int       `json:"amount"`
		CreatedAt            time.Time `json:"created_at"`
	}
)

func formatSliceResponse(transfers []entities.Transfer) []ResponseBody {
	transfersResponse := make([]ResponseBody, len(transfers))

	for i, trans := range transfers {
		transfersResponse[i] = ResponseBody{
			ID:                   trans.ID,
			AccountOriginID:      trans.AccountOriginID,
			AccountDestinationID: trans.AccountDestinationID,
			Amount:               trans.Amount,
			CreatedAt:            trans.CreatedAt,
		}
	}

	return transfersResponse
}
