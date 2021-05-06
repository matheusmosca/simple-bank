package service

import (
	"os"
	"simple-bank/pkg/domain/account"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

var (
	service            Auth
	mockAccountUseCase *account.UseCaseMock
)

func setupAuthService() {
	mockAccountUseCase = &account.UseCaseMock{}

	service = Auth{
		accountUseCase: mockAccountUseCase,
	}
}

// validToken is a test helper function that
// verify if a token is valid and belongs to an provided accountID
func validToken(token string, accountID string, t *testing.T) bool {
	claims := jwt.MapClaims{}

	j, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("AUTH_SECRET")), nil
	})
	if err != nil {
		t.Log(err)
		return false
	}
	if j.Valid && claims["Id"] == accountID {
		return true
	}

	return false
}
