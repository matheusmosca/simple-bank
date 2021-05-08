package usecase

import (
	"context"
	"simple-bank/pkg/domain/entities"
)

func (t Transfer) Perform(ctx context.Context, input entities.CreateTransferInput) (*entities.Transfer, error) {
	return nil, nil
}
