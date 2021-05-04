package usecase

import (
	"context"
	"simple-bank/pkg/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByCPF(t *testing.T) {
	input := entities.CreateAccountInput{
		Name:   "jorge",
		CPF:    "984.157.750-08",
		Secret: "123456",
	}
	ctx := context.Background()

	t.Run("Should return an account without errors", func(t *testing.T) {
		setupAccountUseCase()

		expectedAccount, _ := entities.NewAccount(input.Name, input.CPF, input.Secret)

		// Returns a valid account
		mockRepository.GetByCPFFunc = func(ctx context.Context, cpf string) (*entities.Account, error) {
			return expectedAccount, nil
		}

		a, err := useCase.GetByCPF(ctx, input.CPF)

		assert.Nil(t, err)

		if err != nil {
			assert.Equal(t, a, expectedAccount)
		}
	})

	t.Run("Should not return an account, Account does not exist", func(t *testing.T) {
		setupAccountUseCase()

		acc, _ := entities.NewAccount(input.Name, input.CPF, input.Secret)

		// Does not return an account
		mockRepository.GetByCPFFunc = func(ctx context.Context, cpf string) (*entities.Account, error) {
			return nil, entities.ErrAccountDoesNotExist
		}

		a, err := useCase.GetByCPF(ctx, acc.CPF)

		assert.Equal(t, err, entities.ErrAccountDoesNotExist)
		assert.Nil(t, a)
	})
}
