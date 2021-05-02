package account

import (
	"context"
	"errors"
	"simple-bank/pkg/domain/entities"
)

var (
	ErrCPFAlreadyExists = errors.New(`the cpf is already in use`)
)

type UseCase interface {
	// List(ctx context.Context) ([]*entities.Account, error)
	Create(ctx context.Context, name, CPF, secret string) (*entities.Account, error)
}
