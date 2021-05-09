package service

import (
	"log"
	"os"
	"simple-bank/pkg/domain/auth"

	"github.com/dgrijalva/jwt-go"
)

// Validates the token and returns the account_id or and error
// if the validation goes wrong
func Authorize(token string) (interface{}, error) {
	claims := jwt.MapClaims{}

	if token == "" {
		return nil, auth.ErrTokenNotProvided
	}

	j, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("AUTH_SECRET")), nil
	})

	if err != nil || !j.Valid {
		log.Println(err)
		return nil, auth.ErrInvalidToken
	}

	if id, ok := claims["Id"]; ok {
		return id, nil
	}

	return nil, auth.ErrInvalidToken
}
