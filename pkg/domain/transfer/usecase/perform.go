package usecase

import (
	"context"

	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
)

func (t Transfer) Perform(ctx context.Context, input entities.CreateTransferInput) (*entities.Transfer, error) {
	transfer, err := entities.NewTransfer(
		input.AccountOriginID,
		input.AccountDestinationID,
		input.Amount,
	)
	if err != nil {
		return nil, err
	}

	origAcc, err := t.accountUseCase.GetByID(ctx, input.AccountOriginID)
	if err != nil {
		return nil, entities.ErrOrigAccountDoesNotExist
	}

	destAcc, err := t.accountUseCase.GetByID(ctx, input.AccountDestinationID)
	if err != nil {
		return nil, entities.ErrDestAccountDoesNotExist
	}

	err = origAcc.WithdrawMoney(input.Amount)
	if err != nil {
		return nil, err
	}

	err = destAcc.DepositMoney(input.Amount)
	if err != nil {
		return nil, err
	}

	err = t.repository.PerformTransference(ctx, entities.PerformTransferenceInput{
		OriginAcount:      origAcc,
		DestinationAcount: destAcc,
		Transfer:          transfer,
	})
	if err != nil {
		return nil, err
	}

	return transfer, nil
}
