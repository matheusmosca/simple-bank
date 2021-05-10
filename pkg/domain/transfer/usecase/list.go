package usecase

import (
	"context"
	"log"

	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
)

func (t Transfer) List(ctx context.Context, origID string) ([]entities.Transfer, error) {
	_, err := t.accountUseCase.GetByID(ctx, origID)
	if err != nil {
		return nil, entities.ErrAccountDoesNotExist
	}

	transfers, err := t.repository.ListTransfersByAccountID(ctx, origID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return transfers, nil
}
