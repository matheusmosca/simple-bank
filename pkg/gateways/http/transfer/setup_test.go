package transfer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
	"github.com/matheusmosca/simple-bank/pkg/domain/transfer"
	"github.com/matheusmosca/simple-bank/pkg/gateways/http/util/validator"
)

type mockResponse struct {
	ListErr        error
	ListPayload    []entities.Transfer
	PerformErr     error
	PerformPayload *entities.Transfer
}

func fakeHandler(res mockResponse) Handler {
	transferMock := &transfer.UseCaseMock{
		ListFunc: func(ctx context.Context, origID string) ([]entities.Transfer, error) {
			return res.ListPayload, res.ListErr
		},
		PerformFunc: func(contextMoqParam context.Context, createTransferInput entities.CreateTransferInput) (*entities.Transfer, error) {
			return res.PerformPayload, res.PerformErr
		},
	}

	return Handler{
		useCase:   transferMock,
		validator: validator.New(),
	}
}

func fakeRequest(method, path string, body interface{}) *http.Request {
	reqBody, _ := json.Marshal(body)
	req := bytes.NewReader(reqBody)
	return httptest.NewRequest(method, path, req)
}

// test helper to assert a response
func assertResponseHelper(t *testing.T, wantBody []byte, wantHTTPMethod int, res *httptest.ResponseRecorder) {
	assert.Equal(t, string(wantBody), strings.TrimSpace(res.Body.String()))
	assert.Equal(t, wantHTTPMethod, res.Code)
	assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
}
