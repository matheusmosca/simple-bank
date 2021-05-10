package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
)

func CreateToken(u entities.Account) (string, error) {
	j := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":        u.ID,
		"IssuedAt":  time.Now().Unix(),
		"ExpiresAt": 1500,
	})

	token, err := j.SignedString([]byte(os.Getenv("AUTH_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}
