package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matheusmosca/simple-bank/pkg/common/hash"
	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
)

func TestCreate(t *testing.T) {
	t.Run("Should create an account without errors", func(t *testing.T) {
		setupAccountUseCase()

		input := entities.CreateAccountInput{
			Name:   "jose",
			CPF:    "984.157.750-08",
			Secret: "123456",
		}

		hash, _ := hash.New("123456")
		expectedAcount := &entities.Account{
			Name:    "jose",
			CPF:     "984.157.750-08",
			Secret:  hash,
			Balance: entities.DefaultBalanceValue,
		}

		mockRepository.CreateFunc = func(ctx context.Context, in *entities.Account) error {
			return nil
		}

		// There isn't an account with this CPF
		mockRepository.GetByCPFFunc = func(ctx context.Context, cpf string) (*entities.Account, error) {
			return nil, nil
		}

		acc, err := useCase.Create(context.Background(), input)
		assert.Nil(t, err)
		if err != nil {
			assert.Equal(t, expectedAcount.Name, acc.Name)
			assert.Equal(t, expectedAcount.CPF, acc.CPF)
			assert.Equal(t, expectedAcount.Secret, acc.Secret)
			assert.Equal(t, expectedAcount.Balance, acc.Balance)
		}
	})

	t.Run("Should not create an account, the provided cpf is already used", func(t *testing.T) {
		setupAccountUseCase()

		input := entities.CreateAccountInput{
			Name:   "jose",
			CPF:    "984.157.750-08",
			Secret: "123456",
		}

		mockRepository.CreateFunc = func(ctx context.Context, in *entities.Account) error {
			return entities.ErrCPFAlreadyExists
		}

		// Found an account with the provided CPF
		mockRepository.GetByCPFFunc = func(ctx context.Context, cpf string) (*entities.Account, error) {
			return &entities.Account{}, nil
		}

		acc, err := useCase.Create(context.Background(), input)

		assert.Nil(t, acc)
		assert.Equal(t, err, entities.ErrCPFAlreadyExists)
	})
}
