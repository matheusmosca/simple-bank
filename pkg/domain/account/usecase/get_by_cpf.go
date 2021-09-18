package usecase

import (
	"context"

	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
)

func (a Account) GetByCPF(ctx context.Context, CPF string) (*entities.Account, error) {
	acc, err := a.repository.GetByCPF(ctx, CPF)
	if err != nil {
		return nil, entities.ErrAccountDoesNotExist
	}

	return acc, nil
}
