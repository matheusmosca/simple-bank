package transfer

import (
	"context"
	"simple-bank/pkg/domain/entities"
)

type Repository interface {
	PerformTransference(context.Context, entities.PerformTransferenceInput) error
}
