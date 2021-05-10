package account

import (
	"context"

	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
)

type Repository interface {
	Create(ctx context.Context, account *entities.Account) error
	GetByID(ctx context.Context, accountID string) (*entities.Account, error)
	GetByCPF(ctx context.Context, CPF string) (*entities.Account, error)
	GetAccounts(ctx context.Context) ([]entities.Account, error)
	UpdateBalance(ctx context.Context, ID string) error
}
