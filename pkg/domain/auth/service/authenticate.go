package service

import (
	"context"
	"log"
	"simple-bank/pkg/common/cpf"
	"simple-bank/pkg/common/hash"
	"simple-bank/pkg/domain/auth"
	"simple-bank/pkg/domain/entities"
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
		log.Println(err)
		//? Does not notify the client that there isn't an account with
		//? the provided cpf
		return "", auth.ErrWrongCredentials
	}

	if hash.CompareSecrets(secret, acc.Secret) {
		return createToken(*acc)
	}

	return "", auth.ErrWrongCredentials
}
