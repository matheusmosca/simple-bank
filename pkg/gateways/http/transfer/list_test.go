package transfer

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http/middlewares"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http/util/response"
)

func TestList(t *testing.T) {
	path := "/api/v1/transfers"
	requestBody := ""
	// Fake data
	authContextID := middlewares.AuthContextKey("account_id")
	acc1, _ := entities.NewAccount("Jorge", "762.337.520-27", "123456")
	_ = acc1.DepositMoney(10000)
	acc2, _ := entities.NewAccount("Jorge", "762.337.520-27", "123456")
	transfer1, _ := entities.NewTransfer(acc1.ID, acc2.ID, 100)
	transfer2, _ := entities.NewTransfer(acc2.ID, acc1.ID, 50)
	transfers := []entities.Transfer{*transfer1, *transfer2}

	t.Run("Should return 200 and a slice of transfers", func(t *testing.T) {
		wantResponseBody, _ := json.Marshal(formatSliceResponse(transfers))

		handler := fakeHandler(mockResponse{
			ListErr:     nil,
			ListPayload: transfers,
		})
		request := fakeRequest(http.MethodGet, path, requestBody)
		ctx := context.WithValue(request.Context(), authContextID, acc1.ID)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.List).ServeHTTP(response, request.WithContext(ctx))

		assertResponseHelper(t, wantResponseBody, http.StatusOK, response)
	})

	t.Run("Should return 401 due to account_id not found in context", func(t *testing.T) {
		wantResponseBody, _ := json.Marshal(response.ErrorResponse{
			Message: response.ErrUnauthorized,
		})

		handler := fakeHandler(mockResponse{
			ListErr:     errors.New("account didn't found in context"),
			ListPayload: nil,
		})
		request := fakeRequest(http.MethodGet, path, requestBody)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.List).ServeHTTP(response, request)

		assertResponseHelper(t, wantResponseBody, http.StatusUnauthorized, response)
	})

	t.Run("Should return 400 due to an usecase known error", func(t *testing.T) {
		wantResponseBody, _ := json.Marshal(response.ErrorResponse{
			Message: entities.ErrAccountDoesNotExist.Error(),
		})
		handler := fakeHandler(mockResponse{
			ListErr:     entities.ErrAccountDoesNotExist,
			ListPayload: nil,
		})
		request := fakeRequest(http.MethodGet, path, requestBody)
		ctx := context.WithValue(request.Context(), authContextID, acc1.ID)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.List).ServeHTTP(response, request.WithContext(ctx))

		assertResponseHelper(t, wantResponseBody, http.StatusBadRequest, response)
	})

	t.Run("Should return 500 due to unexpected error", func(t *testing.T) {
		wantResponseBody, _ := json.Marshal(response.ErrorResponse{
			Message: response.ErrIntervalServer,
		})
		handler := fakeHandler(mockResponse{
			ListErr:     errors.New("something went wrong"),
			ListPayload: nil,
		})
		request := fakeRequest(http.MethodGet, path, requestBody)
		ctx := context.WithValue(request.Context(), authContextID, acc1.ID)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.List).ServeHTTP(response, request.WithContext(ctx))

		assertResponseHelper(t, wantResponseBody, http.StatusInternalServerError, response)
	})
}
