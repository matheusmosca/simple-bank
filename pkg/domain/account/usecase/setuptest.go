package usecase

import "github.com/matheusmosca/simple-bank/pkg/domain/account"

var (
	useCase        Account
	mockRepository *account.RepositoryMock
)

func setupAccountUseCase() {
	mockRepository = &account.RepositoryMock{}

	useCase = Account{
		repository: mockRepository,
	}
}
