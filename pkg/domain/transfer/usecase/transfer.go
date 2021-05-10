package usecase

import (
	"github.com/matheusmosca/simple-bank/pkg/domain/account"
	"github.com/matheusmosca/simple-bank/pkg/domain/transfer"
)

var _ transfer.UseCase = Transfer{}

type Transfer struct {
	repository     transfer.Repository
	accountUseCase account.UseCase
}

func NewTransfer(repo transfer.Repository, accUseCase account.UseCase) *Transfer {
	return &Transfer{
		repository:     repo,
		accountUseCase: accUseCase,
	}
}
