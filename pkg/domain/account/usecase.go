package account

import (
	"context"
	"simple-bank/pkg/domain/entities"
)

type UseCase interface {
	List(ctx context.Context) ([]entities.Account, error)
	Create(ctx context.Context, input entities.CreateAccountInput) (*entities.Account, error)
	GetByID(ctx context.Context, accountID string) (*entities.Account, error)
}
