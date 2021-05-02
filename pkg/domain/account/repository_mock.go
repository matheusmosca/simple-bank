// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package account

import (
	"context"
	"simple-bank/pkg/domain/entities"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked Repository
// 		mockedRepository := &RepositoryMock{
// 			CreateFunc: func(ctx context.Context, account *entities.Account) error {
// 				panic("mock out the Create method")
// 			},
// 			GetAccountsFunc: func(ctx context.Context) ([]*entities.Account, error) {
// 				panic("mock out the GetAccounts method")
// 			},
// 			GetByCPFFunc: func(ctx context.Context, CPF string) (*entities.Account, error) {
// 				panic("mock out the GetByCPF method")
// 			},
// 			GetByIDFunc: func(ctx context.Context, accountID string) (*entities.Account, error) {
// 				panic("mock out the GetByID method")
// 			},
// 			UpdateBalanceFunc: func(ctx context.Context, ID string) error {
// 				panic("mock out the UpdateBalance method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, account *entities.Account) error

	// GetAccountsFunc mocks the GetAccounts method.
	GetAccountsFunc func(ctx context.Context) ([]*entities.Account, error)

	// GetByCPFFunc mocks the GetByCPF method.
	GetByCPFFunc func(ctx context.Context, CPF string) (*entities.Account, error)

	// GetByIDFunc mocks the GetByID method.
	GetByIDFunc func(ctx context.Context, accountID string) (*entities.Account, error)

	// UpdateBalanceFunc mocks the UpdateBalance method.
	UpdateBalanceFunc func(ctx context.Context, ID string) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Account is the account argument value.
			Account *entities.Account
		}
		// GetAccounts holds details about calls to the GetAccounts method.
		GetAccounts []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetByCPF holds details about calls to the GetByCPF method.
		GetByCPF []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CPF is the CPF argument value.
			CPF string
		}
		// GetByID holds details about calls to the GetByID method.
		GetByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AccountID is the accountID argument value.
			AccountID string
		}
		// UpdateBalance holds details about calls to the UpdateBalance method.
		UpdateBalance []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the ID argument value.
			ID string
		}
	}
	lockCreate        sync.RWMutex
	lockGetAccounts   sync.RWMutex
	lockGetByCPF      sync.RWMutex
	lockGetByID       sync.RWMutex
	lockUpdateBalance sync.RWMutex
}

// Create calls CreateFunc.
func (mock *RepositoryMock) Create(ctx context.Context, account *entities.Account) error {
	if mock.CreateFunc == nil {
		panic("RepositoryMock.CreateFunc: method is nil but Repository.Create was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Account *entities.Account
	}{
		Ctx:     ctx,
		Account: account,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, account)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedRepository.CreateCalls())
func (mock *RepositoryMock) CreateCalls() []struct {
	Ctx     context.Context
	Account *entities.Account
} {
	var calls []struct {
		Ctx     context.Context
		Account *entities.Account
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// GetAccounts calls GetAccountsFunc.
func (mock *RepositoryMock) GetAccounts(ctx context.Context) ([]*entities.Account, error) {
	if mock.GetAccountsFunc == nil {
		panic("RepositoryMock.GetAccountsFunc: method is nil but Repository.GetAccounts was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAccounts.Lock()
	mock.calls.GetAccounts = append(mock.calls.GetAccounts, callInfo)
	mock.lockGetAccounts.Unlock()
	return mock.GetAccountsFunc(ctx)
}

// GetAccountsCalls gets all the calls that were made to GetAccounts.
// Check the length with:
//     len(mockedRepository.GetAccountsCalls())
func (mock *RepositoryMock) GetAccountsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAccounts.RLock()
	calls = mock.calls.GetAccounts
	mock.lockGetAccounts.RUnlock()
	return calls
}

// GetByCPF calls GetByCPFFunc.
func (mock *RepositoryMock) GetByCPF(ctx context.Context, CPF string) (*entities.Account, error) {
	if mock.GetByCPFFunc == nil {
		panic("RepositoryMock.GetByCPFFunc: method is nil but Repository.GetByCPF was just called")
	}
	callInfo := struct {
		Ctx context.Context
		CPF string
	}{
		Ctx: ctx,
		CPF: CPF,
	}
	mock.lockGetByCPF.Lock()
	mock.calls.GetByCPF = append(mock.calls.GetByCPF, callInfo)
	mock.lockGetByCPF.Unlock()
	return mock.GetByCPFFunc(ctx, CPF)
}

// GetByCPFCalls gets all the calls that were made to GetByCPF.
// Check the length with:
//     len(mockedRepository.GetByCPFCalls())
func (mock *RepositoryMock) GetByCPFCalls() []struct {
	Ctx context.Context
	CPF string
} {
	var calls []struct {
		Ctx context.Context
		CPF string
	}
	mock.lockGetByCPF.RLock()
	calls = mock.calls.GetByCPF
	mock.lockGetByCPF.RUnlock()
	return calls
}

// GetByID calls GetByIDFunc.
func (mock *RepositoryMock) GetByID(ctx context.Context, accountID string) (*entities.Account, error) {
	if mock.GetByIDFunc == nil {
		panic("RepositoryMock.GetByIDFunc: method is nil but Repository.GetByID was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		AccountID string
	}{
		Ctx:       ctx,
		AccountID: accountID,
	}
	mock.lockGetByID.Lock()
	mock.calls.GetByID = append(mock.calls.GetByID, callInfo)
	mock.lockGetByID.Unlock()
	return mock.GetByIDFunc(ctx, accountID)
}

// GetByIDCalls gets all the calls that were made to GetByID.
// Check the length with:
//     len(mockedRepository.GetByIDCalls())
func (mock *RepositoryMock) GetByIDCalls() []struct {
	Ctx       context.Context
	AccountID string
} {
	var calls []struct {
		Ctx       context.Context
		AccountID string
	}
	mock.lockGetByID.RLock()
	calls = mock.calls.GetByID
	mock.lockGetByID.RUnlock()
	return calls
}

// UpdateBalance calls UpdateBalanceFunc.
func (mock *RepositoryMock) UpdateBalance(ctx context.Context, ID string) error {
	if mock.UpdateBalanceFunc == nil {
		panic("RepositoryMock.UpdateBalanceFunc: method is nil but Repository.UpdateBalance was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  ID,
	}
	mock.lockUpdateBalance.Lock()
	mock.calls.UpdateBalance = append(mock.calls.UpdateBalance, callInfo)
	mock.lockUpdateBalance.Unlock()
	return mock.UpdateBalanceFunc(ctx, ID)
}

// UpdateBalanceCalls gets all the calls that were made to UpdateBalance.
// Check the length with:
//     len(mockedRepository.UpdateBalanceCalls())
func (mock *RepositoryMock) UpdateBalanceCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockUpdateBalance.RLock()
	calls = mock.calls.UpdateBalance
	mock.lockUpdateBalance.RUnlock()
	return calls
}