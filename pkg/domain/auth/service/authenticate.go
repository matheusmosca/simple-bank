package service

import (
	"context"

	"github.com/matheusmosca/simple-bank/pkg/common/cpf"
	"github.com/matheusmosca/simple-bank/pkg/common/hash"
	"github.com/matheusmosca/simple-bank/pkg/domain/auth"
	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
)

func (a Auth) Authenticate(ctx context.Context, CPF, secret string) (string, error) {
	if !cpf.Validate(CPF) {
		return "", cpf.ErrInvalidCPF
	}
	if len(secret) < 6 || len(secret) > 50 {
		return "", entities.ErrInvalidSecret
	}

	acc, err := a.accountUseCase.GetByCPF(ctx, CPF)
	if err != nil {
		return "", auth.ErrWrongCredentials
	}

	isAccount, _ := hash.CompareSecrets(secret, acc.Secret)
	if !isAccount {
		return "", auth.ErrWrongCredentials
	}

	token, err := CreateToken(*acc)
	if err != nil {
		return "", err
	}

	return token, nil
}
