package account

import (
	"context"
	"simple-bank/pkg/domain/entities"
)

type UseCase interface {
	// List(ctx context.Context) ([]*entities.Account, error)
	Create(ctx context.Context, name, CPF, secret string) (*entities.Account, error)
}
