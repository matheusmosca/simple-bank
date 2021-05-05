package usecase

import (
	"context"
	"errors"
	"simple-bank/pkg/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAccounts(t *testing.T) {
	acc1, _ := entities.NewAccount("Maria", "929.654.080-10", "12345678")
	acc2, _ := entities.NewAccount("Jorge", "360.112.530-75", "123456789")
	acc3, _ := entities.NewAccount("Joana", "997.755.060-34", "1234d678")

	ctx := context.Background()

	t.Run("Should return a slice of accounts without errors", func(t *testing.T) {
		setupAccountUseCase()

		expectedAccounts := []entities.Account{*acc1, *acc2, *acc3}

		mockRepository.GetAccountsFunc = func(ctx context.Context) ([]entities.Account, error) {
			return expectedAccounts, nil
		}

		accounts, err := useCase.List(ctx)
		assert.Nil(t, err)

		if err != nil {
			assert.Equal(t, expectedAccounts, accounts)
		}
	})

	t.Run("Should return an error due to a repository error", func(t *testing.T) {
		setupAccountUseCase()

		mockRepository.GetAccountsFunc = func(ctx context.Context) ([]entities.Account, error) {
			return nil, errors.New("Could not fetch accounts")
		}

		accounts, err := useCase.List(ctx)

		assert.NotNil(t, err)
		assert.Nil(t, accounts)
	})
}
