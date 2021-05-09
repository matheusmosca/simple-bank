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

type noBody struct {
	Message string `json:"message"`
}

func TestCreateAccount(t *testing.T) {
	path := "/api/v1/accounts"

	t.Run("Should return 201 and create an account", func(t *testing.T) {
		acc, _ := entities.NewAccount("jorge", "685.490.160-04", "12345678")

		handler := fakeHandler(mockResponse{
			CreateAccountErr:     nil,
			CreateAccountPayload: acc,
		})

		request := fakeRequest(http.MethodPost, path, CreateBodyRequest{
			Name:   acc.Name,
			CPF:    acc.CPF,
			Secret: "12345678",
		})

		wantResponseBody, _ := json.Marshal(ResponseBody{
			ID:        acc.ID,
			CPF:       acc.CPF,
			Name:      acc.Name,
			Balance:   acc.Balance,
			CreatedAt: acc.CreatedAt,
		})

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Create).ServeHTTP(response, request)

		assertResponseHelper(t, wantResponseBody, http.StatusCreated, response)
	})

	t.Run("Should return 400 due to a empty body", func(t *testing.T) {
		requestBody := ""

		wantResponseBody, _ := json.Marshal(noBody{Message: "invalid params"})

		handler := fakeHandler(mockResponse{})

		request := fakeRequest(http.MethodPost, path, requestBody)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Create).ServeHTTP(response, request)

		assertResponseHelper(t, wantResponseBody, http.StatusBadRequest, response)
	})

	t.Run("Should return 400 due to invalid fields", func(t *testing.T) {
		requestBody := CreateBodyRequest{
			Secret: "12345677",
		}

		wantResponseBody, _ := json.Marshal(ValidationErrorResponse{
			CPF:  "cpf is a required field",
			Name: "name is a required field",
		})

		handler := fakeHandler(mockResponse{})

		request := fakeRequest(http.MethodPost, path, requestBody)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Create).ServeHTTP(response, request)

		assertResponseHelper(t, wantResponseBody, http.StatusBadRequest, response)
	})

	t.Run("Should return 400 due to usecase validation error", func(t *testing.T) {
		requestBody := CreateBodyRequest{
			//? Invalid CPF
			CPF:    "123.456.123-12",
			Name:   "jorge",
			Secret: "12345677",
		}

		wantResponseBody, _ := json.Marshal(response.ErrorResponse{
			Message: entities.ErrInvalidCPF.Error(),
		})

		handler := fakeHandler(mockResponse{
			CreateAccountErr:     entities.ErrInvalidCPF,
			CreateAccountPayload: nil,
		})

		request := fakeRequest(http.MethodPost, path, requestBody)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Create).ServeHTTP(response, request)

		assertResponseHelper(t, wantResponseBody, http.StatusBadRequest, response)
	})

	t.Run("Should return 500 due to an unexpected behavior", func(t *testing.T) {
		// Valid request body
		requestBody := CreateBodyRequest{
			CPF:    "685.490.160-04",
			Name:   "jorge",
			Secret: "12345677",
		}

		wantResponseBody, _ := json.Marshal(response.ErrorResponse{
			Message: response.ErrIntervalServer,
		})

		handler := fakeHandler(mockResponse{
			CreateAccountErr:     errors.New("something went wrong"),
			CreateAccountPayload: nil,
		})

		request := fakeRequest(http.MethodPost, path, requestBody)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Create).ServeHTTP(response, request)

		assertResponseHelper(t, wantResponseBody, http.StatusInternalServerError, response)
	})
}
