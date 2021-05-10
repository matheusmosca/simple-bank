package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matheusmosca/simple-bank/pkg/domain/auth"
	"github.com/matheusmosca/simple-bank/pkg/domain/auth/service"
	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http/util/response"
)

type noBody struct {
	Message string `json:"message"`
}

func TestLogin(t *testing.T) {
	path := "/api/v1/login"

	t.Run("Should return 200 and a jwt token", func(t *testing.T) {
		requestBody := LoginRequest{
			CPF:    "830.088.320-75",
			Secret: "12345678",
		}

		acc, _ := entities.NewAccount("Jorge", requestBody.CPF, requestBody.Secret)
		token, _ := service.CreateToken(*acc)

		wantResponseBody, _ := json.Marshal(LoginResponse{Token: token})

		handler := fakeHandler(mockResponse{
			authenticateErr: nil,
			token:           token,
		})

		request := fakeRequest(http.MethodPost, path, requestBody)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Login).ServeHTTP(response, request)

		assertResponseHelper(t, wantResponseBody, http.StatusOK, response)
	})

	t.Run("Should return 401 due to wrong credentials", func(t *testing.T) {
		requestBody := LoginRequest{
			CPF:    "830.088.320-75",
			Secret: "12345678",
		}

		wantResponseBody, _ := json.Marshal(response.ErrorResponse{
			Message: auth.ErrWrongCredentials.Error(),
		})

		handler := fakeHandler(mockResponse{
			authenticateErr: auth.ErrWrongCredentials,
			token:           "",
		})

		request := fakeRequest(http.MethodPost, path, requestBody)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Login).ServeHTTP(response, request)

		assertResponseHelper(t, wantResponseBody, http.StatusUnauthorized, response)
	})

	t.Run("Should return 400 due to a empty request body", func(t *testing.T) {
		requestBody := ""

		wantResponseBody, _ := json.Marshal(noBody{Message: "invalid params"})

		handler := fakeHandler(mockResponse{})

		request := fakeRequest(http.MethodPost, path, requestBody)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Login).ServeHTTP(response, request)

		assertResponseHelper(t, wantResponseBody, http.StatusBadRequest, response)
	})

	t.Run("Should return 400 due to invalid fields", func(t *testing.T) {
		requestBody := LoginRequest{}

		wantResponseBody, _ := json.Marshal(ValidationErrorResponse{
			CPF:    "cpf is a required field",
			Secret: "secret is a required field",
		})

		handler := fakeHandler(mockResponse{})

		request := fakeRequest(http.MethodPost, path, requestBody)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Login).ServeHTTP(response, request)

		assertResponseHelper(t, wantResponseBody, http.StatusBadRequest, response)
	})
}
