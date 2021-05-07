package account

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"simple-bank/pkg/domain/entities"
	"simple-bank/pkg/gateways/http/util/response"
	"testing"
)

func TestListAccounts(t *testing.T) {
	path := "/api/v1/accounts"

	acc1, _ := entities.NewAccount("Jorge", "478.707.610-87", "12345678")
	acc2, _ := entities.NewAccount("Maria", "084.792.650-86", "1sd2345234")
	accountsResponse := []ResponseBody{
		{
			ID:        acc1.ID,
			CPF:       acc1.CPF,
			Name:      acc1.Name,
			Balance:   acc1.DisplayBalance(),
			CreatedAt: acc1.CreatedAt,
		},
		{
			ID:        acc2.ID,
			CPF:       acc2.CPF,
			Name:      acc2.Name,
			Balance:   acc2.DisplayBalance(),
			CreatedAt: acc2.CreatedAt,
		},
	}

	t.Run("Should return 200 and a empty slice", func(t *testing.T) {
		requestBody := ""

		wantBody, _ := json.Marshal(accountsResponse)

		handler := fakeHandler(mockResponse{
			ListErr:     nil,
			ListPayload: []entities.Account{*acc1, *acc2},
		})

		req := fakeRequest(http.MethodGet, path, requestBody)

		res := httptest.NewRecorder()

		http.HandlerFunc(handler.List).ServeHTTP(res, req)

		assertResponseHelper(t, wantBody, http.StatusOK, res)
	})

	t.Run("Should return 500", func(t *testing.T) {
		requestBody := ""

		wantBody, _ := json.Marshal(response.ErrorResponse{
			Message: response.ErrIntervalServer,
		})

		handler := fakeHandler(mockResponse{
			ListErr:     errors.New("something went wrong"),
			ListPayload: nil,
		})

		req := fakeRequest(http.MethodGet, path, requestBody)

		res := httptest.NewRecorder()

		http.HandlerFunc(handler.List).ServeHTTP(res, req)

		assertResponseHelper(t, wantBody, http.StatusInternalServerError, res)
	})
}
