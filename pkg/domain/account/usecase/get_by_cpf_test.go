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

		mockRepository.CreateFunc = func(ctx context.Context, a *entities.Account) error {
			return nil
		}

		mockRepository.GetByCPFFunc = func(ctx context.Context, cpf string) (*entities.Account, error) {
			return nil, nil
		}

		expectedAccount, err := useCase.Create(ctx, input)
		assert.Nil(t, err)

		// Returns a valid account
		mockRepository.GetByCPFFunc = func(ctx context.Context, cpf string) (*entities.Account, error) {
			return expectedAccount, nil
		}

		if err != nil {
			a, err := useCase.GetByCPF(ctx, input.CPF)

			assert.Nil(t, err)

			if err != nil {
				assert.Equal(t, expectedAccount.Name, a.Name)
				assert.Equal(t, expectedAccount.CPF, a.CPF)
				assert.Equal(t, expectedAccount.Secret, a.Secret)
				assert.Equal(t, expectedAccount.Balance, a.Balance)
			}
		}
	})

	t.Run("Should not return an account, Account does not exist", func(t *testing.T) {
		setupAccountUseCase()

		mockRepository.CreateFunc = func(ctx context.Context, a *entities.Account) error {
			return nil
		}

		mockRepository.GetByCPFFunc = func(ctx context.Context, cpf string) (*entities.Account, error) {
			return nil, nil
		}

		_, err := useCase.Create(ctx, input)
		assert.Nil(t, err)

		// Returns a valid account
		mockRepository.GetByCPFFunc = func(ctx context.Context, cpf string) (*entities.Account, error) {
			return nil, entities.ErrAccountDoesNotExist
		}

		if err != nil {
			a, err := useCase.GetByCPF(ctx, input.CPF)

			assert.Equal(t, err, entities.ErrAccountDoesNotExist)
			assert.Nil(t, a)

		}
	})
}
