package usecase

import (
	"context"
	"log"

	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
)

func (a Account) List(ctx context.Context) ([]entities.Account, error) {
	accounts, err := a.repository.GetAccounts(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return accounts, nil
}
