// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package transfer

import (
	"context"
	"sync"

	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
)

// Ensure, that UseCaseMock does implement UseCase.
// If this is not the case, regenerate this file with moq.
var _ UseCase = &UseCaseMock{}

// UseCaseMock is a mock implementation of UseCase.
//
// 	func TestSomethingThatUsesUseCase(t *testing.T) {
//
// 		// make and configure a mocked UseCase
// 		mockedUseCase := &UseCaseMock{
// 			ListFunc: func(ctx context.Context, origID string) ([]entities.Transfer, error) {
// 				panic("mock out the List method")
// 			},
// 			PerformFunc: func(contextMoqParam context.Context, createTransferInput entities.CreateTransferInput) (*entities.Transfer, error) {
// 				panic("mock out the Perform method")
// 			},
// 		}
//
// 		// use mockedUseCase in code that requires UseCase
// 		// and then make assertions.
//
// 	}
type UseCaseMock struct {
	// ListFunc mocks the List method.
	ListFunc func(ctx context.Context, origID string) ([]entities.Transfer, error)

	// PerformFunc mocks the Perform method.
	PerformFunc func(contextMoqParam context.Context, createTransferInput entities.CreateTransferInput) (*entities.Transfer, error)

	// calls tracks calls to the methods.
	calls struct {
		// List holds details about calls to the List method.
		List []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// OrigID is the origID argument value.
			OrigID string
		}
		// Perform holds details about calls to the Perform method.
		Perform []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// CreateTransferInput is the createTransferInput argument value.
			CreateTransferInput entities.CreateTransferInput
		}
	}
	lockList    sync.RWMutex
	lockPerform sync.RWMutex
}

// List calls ListFunc.
func (mock *UseCaseMock) List(ctx context.Context, origID string) ([]entities.Transfer, error) {
	if mock.ListFunc == nil {
		panic("UseCaseMock.ListFunc: method is nil but UseCase.List was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		OrigID string
	}{
		Ctx:    ctx,
		OrigID: origID,
	}
	mock.lockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	mock.lockList.Unlock()
	return mock.ListFunc(ctx, origID)
}

// ListCalls gets all the calls that were made to List.
// Check the length with:
//     len(mockedUseCase.ListCalls())
func (mock *UseCaseMock) ListCalls() []struct {
	Ctx    context.Context
	OrigID string
} {
	var calls []struct {
		Ctx    context.Context
		OrigID string
	}
	mock.lockList.RLock()
	calls = mock.calls.List
	mock.lockList.RUnlock()
	return calls
}

// Perform calls PerformFunc.
func (mock *UseCaseMock) Perform(contextMoqParam context.Context, createTransferInput entities.CreateTransferInput) (*entities.Transfer, error) {
	if mock.PerformFunc == nil {
		panic("UseCaseMock.PerformFunc: method is nil but UseCase.Perform was just called")
	}
	callInfo := struct {
		ContextMoqParam     context.Context
		CreateTransferInput entities.CreateTransferInput
	}{
		ContextMoqParam:     contextMoqParam,
		CreateTransferInput: createTransferInput,
	}
	mock.lockPerform.Lock()
	mock.calls.Perform = append(mock.calls.Perform, callInfo)
	mock.lockPerform.Unlock()
	return mock.PerformFunc(contextMoqParam, createTransferInput)
}

// PerformCalls gets all the calls that were made to Perform.
// Check the length with:
//     len(mockedUseCase.PerformCalls())
func (mock *UseCaseMock) PerformCalls() []struct {
	ContextMoqParam     context.Context
	CreateTransferInput entities.CreateTransferInput
} {
	var calls []struct {
		ContextMoqParam     context.Context
		CreateTransferInput entities.CreateTransferInput
	}
	mock.lockPerform.RLock()
	calls = mock.calls.Perform
	mock.lockPerform.RUnlock()
	return calls
}
