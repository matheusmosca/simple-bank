package transfer

import (
	"context"

	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
)

type Repository interface {
	PerformTransference(context.Context, entities.PerformTransferenceInput) error
	ListTransfersByAccountID(context.Context, string) ([]entities.Transfer, error)
}
