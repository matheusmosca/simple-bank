package usecase

import (
	"context"
	"simple-bank/pkg/domain/account"
	"simple-bank/pkg/domain/entities"
)

func (a Account) Create(ctx context.Context, name, CPF, secret string, balance int) (*entities.Account, error) {
	checkCPF, _ := a.repository.GetByCPF(ctx, CPF)
	if checkCPF != nil {
		return nil, account.ErrCPFAlreadyExists
	}

	acc, err := entities.NewAccount(name, CPF, secret, balance)
	if err != nil {
		return nil, err
	}

	err = a.repository.Create(ctx, acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
