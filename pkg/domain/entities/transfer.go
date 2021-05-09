package entities

import (
	"errors"
	"time"
)

var (
	ErrInvalidAmount           = errors.New("the amount must be greater than 0")
	ErrOrigAccEqualDestAcc     = errors.New("the destination account can't be equal the origin account")
	ErrInsufficientFunds       = errors.New("the account has insufficient funds")
	ErrOrigAccountDoesNotExist = errors.New("the origin account does not exist")
	ErrDestAccountDoesNotExist = errors.New("the destination account does not exist")
	TransferDomainErrors       = []error{
		ErrInvalidAmount,
		ErrOrigAccEqualDestAcc,
		ErrInsufficientFunds,
		ErrOrigAccountDoesNotExist,
		ErrDestAccountDoesNotExist,
		ErrAccountDoesNotExist,
	}
)

type CreateTransferInput struct {
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
}

type PerformTransferenceInput struct {
	OriginAcount      *Account
	DestinationAcount *Account
	Transfer          *Transfer
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
	if t.AccountDestinationID == t.AccountOriginID {
		return ErrOrigAccEqualDestAcc
	}

	return nil
}
