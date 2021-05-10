package usecase

import "github.com/matheusmosca/simple-bank/pkg/domain/account"

var _ account.UseCase = Account{}

type Account struct {
	repository account.Repository
}

func NewAccountUseCase(repo account.Repository) *Account {
	return &Account{
		repository: repo,
	}
}
