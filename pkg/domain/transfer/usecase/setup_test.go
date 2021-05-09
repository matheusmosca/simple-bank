package usecase

import (
	"context"
	"simple-bank/pkg/domain/account"
	"simple-bank/pkg/domain/entities"
	"simple-bank/pkg/domain/transfer"
)

type mockResponse struct {
	PerformTransferenceErr          error
	GetAccountByIDErr               error
	GetAccountByIDPayload           *entities.Account
	ListTransfersByAccountIDErr     error
	ListTransfersByAccountIDPayload []entities.Transfer
}

func fakeUseCase(res mockResponse) Transfer {
	repo := &transfer.RepositoryMock{
		PerformTransferenceFunc: func(ctx context.Context, in entities.PerformTransferenceInput) error {
			return res.PerformTransferenceErr
		},
		ListTransfersByAccountIDFunc: func(contextMoqParam context.Context, s string) ([]entities.Transfer, error) {
			return res.ListTransfersByAccountIDPayload, res.ListTransfersByAccountIDErr
		},
	}

	accUseCase := &account.UseCaseMock{
		GetByIDFunc: func(ctx context.Context, accountID string) (*entities.Account, error) {
			return res.GetAccountByIDPayload, res.GetAccountByIDErr
		},
	}

	return Transfer{
		repository:     repo,
		accountUseCase: accUseCase,
	}
}
