package auth

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

	"simple-bank/pkg/domain/auth"
	"simple-bank/pkg/gateways/http/util/validator"
)

type mockResponse struct {
	authenticateErr error
	token           string
}

func fakeHandler(res mockResponse) Handler {
	mockAuthService := auth.ServiceMock{
		AuthenticateFunc: func(ctx context.Context, cpf, secret string) (string, error) {
			return res.token, res.authenticateErr
		},
	}

	return Handler{
		authService: &mockAuthService,
		validator:   validator.New(),
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
