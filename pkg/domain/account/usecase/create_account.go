package usecase

import (
	"context"
	"simple-bank/pkg/domain/entities"
)

func (a Account) Create(ctx context.Context, input entities.CreateAccountInput) (*entities.Account, error) {
	checkCPF, _ := a.repository.GetByCPF(ctx, input.CPF)
	if checkCPF != nil {
		return nil, entities.ErrCPFAlreadyExists
	}

	acc, err := entities.NewAccount(input.Name, input.CPF, input.Secret)
	if err != nil {
		return nil, err
	}

	err = a.repository.Create(ctx, acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
