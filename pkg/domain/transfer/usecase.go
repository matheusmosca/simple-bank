package transfer

import (
	"context"
	"simple-bank/pkg/domain/entities"
)

type UseCase interface {
	//? see the param origID
	List(ctx context.Context, origID string) ([]entities.Transfer, error)
	Perform(context.Context, entities.CreateTransferInput) (*entities.Transfer, error)
}
