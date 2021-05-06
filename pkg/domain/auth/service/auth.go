package service

import (
	"context"
	"simple-bank/pkg/domain/account"
	"simple-bank/pkg/domain/auth"
)

var _ auth.Service = Auth{}

type Auth struct {
	accountUseCase account.UseCase
}

func NewAuthService(accUseCase account.UseCase) auth.Service {
	return &Auth{
		accountUseCase: accUseCase,
	}
}

func (a Auth) Authenticate(ctx context.Context, CPF, secret string) (string, error) {
	return "", nil
}
