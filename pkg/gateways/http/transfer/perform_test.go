package transfer

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"simple-bank/pkg/domain/entities"
	"simple-bank/pkg/gateways/http/middlewares"
	"simple-bank/pkg/gateways/http/util/response"
	"testing"
)

func TestPerform(t *testing.T) {
	path := "/api/v1/transfers"
	// Fake data
	authContextID := middlewares.AuthContextKey("account_id")
	acc1, _ := entities.NewAccount("Jorge", "762.337.520-27", "123456")
	acc1.DepositMoney(10000)
	amount := 1000
	acc2, _ := entities.NewAccount("Jorge", "762.337.520-27", "123456")
	transfer, _ := entities.NewTransfer(acc1.ID, acc2.ID, amount)

	t.Run("Should return 201 and return a transfer", func(t *testing.T) {
		requestBody := PerformRequest{
			AccountDestinationID: acc2.ID,
			Amount:               amount,
		}

		wantResponseBody, _ := json.Marshal(ResponseBody{
			ID:                   transfer.ID,
			AccountOriginID:      acc1.ID,
			AccountDestinationID: acc2.ID,
			Amount:               amount,
			CreatedAt:            transfer.CreatedAt,
		})

		handler := fakeHandler(mockResponse{
			PerformErr:     nil,
			PerformPayload: transfer,
		})
		request := fakeRequest(http.MethodPost, path, requestBody)
		ctx := context.WithValue(request.Context(), authContextID, acc1.ID)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.PerformTransference).ServeHTTP(response, request.WithContext(ctx))

		assertResponseHelper(t, wantResponseBody, http.StatusCreated, response)
	})

	t.Run("Should return 401 due to an authorization error", func(t *testing.T) {
		requestBody := PerformRequest{
			AccountDestinationID: acc2.ID,
			Amount:               amount,
		}
		wantResponseBody, _ := json.Marshal(response.ErrorResponse{
			Message: response.ErrUnauthorized,
		})

		handler := fakeHandler(mockResponse{
			PerformErr:     nil,
			PerformPayload: nil,
		})
		request := fakeRequest(http.MethodPost, path, requestBody)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.PerformTransference).ServeHTTP(response, request)

		assertResponseHelper(t, wantResponseBody, http.StatusUnauthorized, response)
	})

	t.Run("Should return 400 due to empty request body", func(t *testing.T) {
		requestBody := ""
		wantResponseBody, _ := json.Marshal(response.ErrorResponse{
			Message: response.ErrDecode,
		})

		handler := fakeHandler(mockResponse{
			PerformErr:     nil,
			PerformPayload: nil,
		})
		request := fakeRequest(http.MethodPost, path, requestBody)
		ctx := context.WithValue(request.Context(), authContextID, acc1.ID)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.PerformTransference).ServeHTTP(response, request.WithContext(ctx))

		assertResponseHelper(t, wantResponseBody, http.StatusBadRequest, response)
	})

	t.Run("Should return 400 due to missing amount", func(t *testing.T) {
		requestBody := PerformRequest{
			AccountDestinationID: acc2.ID,
		}
		wantResponseBody, _ := json.Marshal(ValidationErrResponse{
			Amount: "amount is a required field",
		})

		handler := fakeHandler(mockResponse{
			PerformErr:     errors.New("destination account not provided"),
			PerformPayload: nil,
		})
		request := fakeRequest(http.MethodPost, path, requestBody)
		ctx := context.WithValue(request.Context(), authContextID, acc1.ID)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.PerformTransference).ServeHTTP(response, request.WithContext(ctx))

		assertResponseHelper(t, wantResponseBody, http.StatusBadRequest, response)
	})

	t.Run("Should return 500 due to an unknown error", func(t *testing.T) {
		// valid request
		requestBody := PerformRequest{
			AccountDestinationID: acc2.ID,
			Amount:               amount,
		}
		wantResponseBody, _ := json.Marshal(response.ErrorResponse{
			Message: response.ErrIntervalServer,
		})

		handler := fakeHandler(mockResponse{
			PerformErr:     errors.New("something went wrong"),
			PerformPayload: nil,
		})
		request := fakeRequest(http.MethodPost, path, requestBody)
		ctx := context.WithValue(request.Context(), authContextID, acc1.ID)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.PerformTransference).ServeHTTP(response, request.WithContext(ctx))

		assertResponseHelper(t, wantResponseBody, http.StatusInternalServerError, response)
	})

}
