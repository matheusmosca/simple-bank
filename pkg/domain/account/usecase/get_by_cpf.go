package usecase

import (
	"context"
	"log"

	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
)

func (a Account) GetByCPF(ctx context.Context, CPF string) (*entities.Account, error) {
	acc, err := a.repository.GetByCPF(ctx, CPF)
	if err != nil {
		log.Println(err)
		return nil, entities.ErrAccountDoesNotExist
	}

	return acc, nil
}
