package usecase

import (
	"context"
	"log"
	"simple-bank/pkg/domain/entities"
)

func (a Account) GetByID(ctx context.Context, accountID string) (*entities.Account, error) {
	acc, err := a.repository.GetByID(ctx, accountID)
	if err != nil {
		log.Println(err)
		return nil, entities.ErrAccountDoesNotExist
	}

	return acc, nil
}
