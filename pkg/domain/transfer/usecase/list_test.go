package usecase

import (
	"context"
	"errors"
	"simple-bank/pkg/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	ctx := context.Background()
	acc1, _ := entities.NewAccount("Jorge", "762.337.520-27", "123456")
	acc1.DepositMoney(10000)
	acc2, _ := entities.NewAccount("Jorge", "762.337.520-27", "123456")
	transfer1, _ := entities.NewTransfer(acc1.ID, acc2.ID, 100)
	transfer2, _ := entities.NewTransfer(acc2.ID, acc1.ID, 50)
	transfers := []entities.Transfer{*transfer1, *transfer2}

	t.Run("should return an list of transfers without errors", func(t *testing.T) {
		want := transfers
		useCase := fakeUseCase(mockResponse{
			GetAccountByIDErr:               nil,
			GetAccountByIDPayload:           acc1,
			ListTransfersByAccountIDErr:     nil,
			ListTransfersByAccountIDPayload: transfers,
		})

		got, err := useCase.List(ctx, acc1.ID)

		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, got, want)
		}
	})

	t.Run("should throw an error because the account_id is invalid", func(t *testing.T) {
		useCase := fakeUseCase(mockResponse{
			GetAccountByIDErr:               entities.ErrAccountDoesNotExist,
			GetAccountByIDPayload:           nil,
			ListTransfersByAccountIDErr:     errors.New("this method was not reached"),
			ListTransfersByAccountIDPayload: nil,
		})

		got, err := useCase.List(ctx, acc1.ID)

		assert.Equal(t, entities.ErrAccountDoesNotExist, err)
		assert.Nil(t, got)
	})

	t.Run("should throw an error due to something wrong on repository", func(t *testing.T) {
		useCase := fakeUseCase(mockResponse{
			GetAccountByIDErr:               nil,
			GetAccountByIDPayload:           acc1,
			ListTransfersByAccountIDErr:     errors.New("something went wrong on repository"),
			ListTransfersByAccountIDPayload: nil,
		})

		got, err := useCase.List(ctx, acc1.ID)

		assert.Equal(t, errors.New("something went wrong on repository"), err)
		assert.Nil(t, got)
	})
}
