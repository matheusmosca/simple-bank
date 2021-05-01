package usecase

import (
	"context"
	"testing"

	account_domain "simple-bank/pkg/domain/account"
	"simple-bank/pkg/domain/entities"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Name                       string
	CPF                        string
	Secret                     string
	Balance                    int
	Want                       error
	RepositoryCreateResponse   error
	RepositoryGetByCPFResponse *entities.Account
	Message                    string
}

var (
	useCase        Account
	mockRepository *account_domain.RepositoryMock
)

func setupMock(createResponse error, getByCPFResponse *entities.Account) {
	mockRepository = &account_domain.RepositoryMock{}
	mockRepository.CreateFunc = func(in1 context.Context, acc *entities.Account) error {
		return createResponse
	}

	mockRepository.GetByCPFFunc = func(in1 context.Context, CPF string) (*entities.Account, error) {
		return getByCPFResponse, nil
	}

	useCase = Account{
		repository: mockRepository,
	}
}

func TestCreate(t *testing.T) {
	testCases := []testCase{
		{
			Name:                     "Maria",
			CPF:                      "042.954.750-17",
			Secret:                   "123456",
			Balance:                  0,
			Want:                     nil,
			RepositoryCreateResponse: nil,
			Message:                  "Should create an account successfully",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Message, func(t *testing.T) {
			setupMock(tc.RepositoryCreateResponse, tc.RepositoryGetByCPFResponse)

			_, err := useCase.Create(context.Background(), tc.Name, tc.CPF, tc.Secret, tc.Balance)
			assert.Equal(t, tc.Want, err)
		})
	}
}

// func TestCreateAccount(t *testing.T) {
// 	t.Run("Should create a account successfully", func(t *testing.T) {
// 		mockRepository := &account_domain.RepositoryMock{}
// 		testUseCase := Account{
// 			repository: mockRepository,
// 		}

// 		mockRepository.CreateFunc = func(in1 context.Context, acc *entities.Account) error {
// 			return nil
// 		}

// 		err := testUseCase.Create(context.Background(), "Maria", "042.954.750-17", "123456", 0)

// 		assert.Nil(t, err)
// 	})
// }
