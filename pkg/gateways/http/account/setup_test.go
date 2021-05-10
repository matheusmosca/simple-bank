package account

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matheusmosca/simple-bank/pkg/domain/account"
	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http/util/validator"
)

type mockResponse struct {
	CreateAccountErr      error
	CreateAccountPayload  *entities.Account
	GetByIDAccountErr     error
	GetByIDAccountPayload *entities.Account
	ListErr               error
	ListPayload           []entities.Account
}

func fakeHandler(res mockResponse) Handler {
	mockAccountUseCase := &account.UseCaseMock{
		CreateFunc: func(ctx context.Context, input entities.CreateAccountInput) (*entities.Account, error) {
			return res.CreateAccountPayload, res.CreateAccountErr
		},
		GetByIDFunc: func(ctx context.Context, accountID string) (*entities.Account, error) {
			return res.GetByIDAccountPayload, res.GetByIDAccountErr
		},
		ListFunc: func(ctx context.Context) ([]entities.Account, error) {
			return res.ListPayload, res.ListErr
		},
	}

	return Handler{
		usecase:   mockAccountUseCase,
		validator: validator.New(),
	}
}

func fakeRequest(method, path string, body interface{}) *http.Request {
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
	}
	req := bytes.NewReader(reqBody)
	return httptest.NewRequest(method, path, req)
}

// test helper to assert a response
func assertResponseHelper(t *testing.T, wantBody []byte, wantHTTPMethod int, res *httptest.ResponseRecorder) {
	assert.Equal(t, string(wantBody), strings.TrimSpace(res.Body.String()))
	assert.Equal(t, wantHTTPMethod, res.Code)
	assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
}
