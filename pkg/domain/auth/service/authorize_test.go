package service

import (
	"simple-bank/pkg/domain/auth"
	"simple-bank/pkg/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorize(t *testing.T) {
	secret := "12345678"
	cpf := "496.923.150-88"
	acc, _ := entities.NewAccount("Jorge", cpf, secret)

	t.Run("Should authorize an account without errors", func(t *testing.T) {
		setupAuthService()
		token, _ := CreateToken(*acc)
		id, err := Authorize(token)

		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, id, acc.ID)
		}
	})

	t.Run("Should return an error because the token is empty", func(t *testing.T) {
		setupAuthService()
		token := ""
		id, err := Authorize(token)

		assert.Equal(t, err, auth.ErrTokenNotProvided)
		assert.Nil(t, id)
	})

	t.Run("Should return an error because the token is not valid", func(t *testing.T) {
		setupAuthService()
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		id, err := Authorize(token)

		assert.Equal(t, err, auth.ErrInvalidToken)
		assert.Nil(t, id)
	})
}
