package service

import (
	"github.com/matheusmosca/simple-bank/pkg/domain/account"
	"github.com/matheusmosca/simple-bank/pkg/domain/auth"
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
