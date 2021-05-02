package usecase

import (
	"context"
	"simple-bank/pkg/domain/entities"
)

func (a Account) Create(ctx context.Context, name, CPF, secret string) (*entities.Account, error) {
	checkCPF, _ := a.repository.GetByCPF(ctx, CPF)
	if checkCPF != nil {
		return nil, entities.ErrCPFAlreadyExists
	}

	acc, err := entities.NewAccount(name, CPF, secret)
	if err != nil {
		return nil, err
	}

	err = a.repository.Create(ctx, acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
