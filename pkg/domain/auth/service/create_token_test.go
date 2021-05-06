package service

import (
	"os"
	"simple-bank/pkg/domain/entities"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {

	t.Run("Should create a valid JWT token", func(t *testing.T) {
		acc, _ := entities.NewAccount("Jorge", "830.088.320-75", "12345678")

		claims := jwt.MapClaims{}
		token, err := createToken(*acc)

		assert.NotNil(t, token)
		assert.Nil(t, err)
		if err != nil {
			return
		}

		j, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("AUTH_SECRET")), nil
		})

		assert.Nil(t, err)
		if err != nil {
			assert.True(t, j.Valid)
			assert.Equal(t, acc.ID, claims["Id"])
		}
	})
}
