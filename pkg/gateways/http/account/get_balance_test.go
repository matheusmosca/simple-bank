package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"simple-bank/pkg/domain/entities"
	"simple-bank/pkg/gateways/http/util/response"
	"testing"
)

func TestGetBalance(t *testing.T) {
	acc, _ := entities.NewAccount("Jorge", "206.209.950-92", "kdlfsdfhd23")
	//? The repository returns the account without secret
	acc.Secret = ""
	path := fmt.Sprintf("/api/v1/accounts/%s/balance", acc.ID)
	requestBody := ""

	t.Run("should return 200 and a balance of an account", func(t *testing.T) {
		wantBody, _ := json.Marshal(BalanceResponse{
			Balance: acc.DisplayBalance(),
		})

		handler := fakeHandler(mockResponse{
			GetByIDAccountErr:     nil,
			GetByIDAccountPayload: acc,
		})

		res := httptest.NewRecorder()

		req := fakeRequest(http.MethodGet, path, requestBody)

		http.HandlerFunc(handler.GetBalance).ServeHTTP(res, req)

		assertResponseHelper(t, wantBody, http.StatusOK, res)
	})

	t.Run("should return 404 due to an known error on usecase", func(t *testing.T) {
		wantBody, _ := json.Marshal(response.ErrorResponse{
			Message: entities.ErrAccountDoesNotExist.Error(),
		})

		handler := fakeHandler(mockResponse{
			GetByIDAccountErr:     entities.ErrAccountDoesNotExist,
			GetByIDAccountPayload: nil,
		})

		res := httptest.NewRecorder()
		req := fakeRequest(http.MethodGet, path, requestBody)

		http.HandlerFunc(handler.GetBalance).ServeHTTP(res, req)

		assertResponseHelper(t, wantBody, http.StatusNotFound, res)
	})

	t.Run("should return 500 due to an unknow error on usecase", func(t *testing.T) {
		wantBody, _ := json.Marshal(response.ErrorResponse{
			Message: response.ErrIntervalServer,
		})

		handler := fakeHandler(mockResponse{
			GetByIDAccountErr:     errors.New("Something went wrong"),
			GetByIDAccountPayload: nil,
		})

		res := httptest.NewRecorder()
		req := fakeRequest(http.MethodGet, path, requestBody)

		http.HandlerFunc(handler.GetBalance).ServeHTTP(res, req)

		assertResponseHelper(t, wantBody, http.StatusInternalServerError, res)
	})
}
