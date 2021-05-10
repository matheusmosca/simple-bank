package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
)

func TestGetByID(t *testing.T) {
	input := entities.CreateAccountInput{
		Name:   "Maria",
		CPF:    "984.157.750-08",
		Secret: "123456",
	}
	ctx := context.Background()

	t.Run("Should return an account without errors", func(t *testing.T) {
		setupAccountUseCase()

		expectedAccount, _ := entities.NewAccount(input.Name, input.CPF, input.Secret)

		// Returns a valid account
		mockRepository.GetByIDFunc = func(ctx context.Context, id string) (*entities.Account, error) {
			return expectedAccount, nil
		}

		a, err := useCase.GetByID(ctx, expectedAccount.ID)

		assert.Nil(t, err)

		if err != nil {
			assert.Equal(t, a, expectedAccount)
		}
	})

	t.Run("Should not return an account, Account does not exist", func(t *testing.T) {
		setupAccountUseCase()

		acc, _ := entities.NewAccount(input.Name, input.CPF, input.Secret)

		// Does not return an account
		mockRepository.GetByIDFunc = func(ctx context.Context, id string) (*entities.Account, error) {
			return nil, entities.ErrAccountDoesNotExist
		}

		a, err := useCase.GetByID(ctx, acc.ID)

		assert.Equal(t, err, entities.ErrAccountDoesNotExist)
		assert.Nil(t, a)
	})
}
