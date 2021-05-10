package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matheusmosca/simple-bank/pkg/domain/auth"
	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
)

func TestAuthenticate(t *testing.T) {

	t.Run("Should authorize an account without errors", func(t *testing.T) {
		setupAuthService()

		ctx := context.Background()
		secret := "12345678"
		cpf := "496.923.150-88"

		acc, _ := entities.NewAccount("Jorge", cpf, secret)

		mockAccountUseCase.GetByCPFFunc = func(ctx context.Context, cpf string) (*entities.Account, error) {
			return acc, nil
		}

		token, err := service.Authenticate(ctx, cpf, secret)

		assert.Nil(t, err)
		assert.True(t, validToken(token, acc.ID, t))
	})

	t.Run("Should not authorize due to wrong cpf", func(t *testing.T) {
		setupAuthService()
		ctx := context.Background()

		//? some wrong cpf
		wrongCPF := "116.587.650-79"
		secret := "12345678"

		// Repository doesn't found an account with the provided cpf
		mockAccountUseCase.GetByCPFFunc = func(ctx context.Context, cpf string) (*entities.Account, error) {
			return nil, auth.ErrWrongCredentials
		}

		token, err := service.Authenticate(ctx, wrongCPF, secret)

		assert.Equal(t, auth.ErrWrongCredentials, err)
		assert.Empty(t, token)
	})

	t.Run("Should not authorize due to an incorrect secret", func(t *testing.T) {
		setupAuthService()

		ctx := context.Background()
		cpf := "496.923.150-88"
		secret := "12345678"

		acc, _ := entities.NewAccount("Jorge", cpf, secret)

		// Repository found an account
		mockAccountUseCase.GetByCPFFunc = func(ctx context.Context, cpf string) (*entities.Account, error) {
			return acc, nil
		}

		incorrectSecret := "someOtherSecret"

		// call with a wrong secret
		token, err := service.Authenticate(ctx, cpf, incorrectSecret)

		assert.Equal(t, auth.ErrWrongCredentials, err)
		assert.Empty(t, token)
	})
}
