package entities

import (
	"errors"
	"time"
)

var (
	ErrInvalidAmount     = errors.New("the amount must be greater than 0")
	TransferDomainErrors = []error{ErrInvalidAmount}
)

type CreateTransferInput struct {
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
}

type Transfer struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
	CreatedAt            time.Time
}

func NewTransfer(origID, destID string, amount int) (*Transfer, error) {
	trans := &Transfer{
		ID:                   newID(),
		AccountOriginID:      origID,
		AccountDestinationID: destID,
		Amount:               amount,
		CreatedAt:            time.Now(),
	}

	err := trans.Validate()
	if err != nil {
		return nil, err
	}

	return trans, nil
}

func (t Transfer) Validate() error {
	if t.Amount <= 0 {
		return ErrInvalidAmount
	}

	return nil
}
