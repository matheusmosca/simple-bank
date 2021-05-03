package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name    string
	cpf     string
	secret  string
	want    error
	message string
}

func TestNewAccount(t *testing.T) {
	tCases := []testCase{
		{
			name:    "jorge",
			cpf:     "031.915.990-31",
			secret:  "123456789",
			want:    ErrInvalidCPF,
			message: "Should not create an account due to a invalid cpf",
		},
		{
			name:    "",
			cpf:     "031.915.990-61",
			secret:  "123456789",
			want:    ErrInvalidName,
			message: "Should not create an account due to a invalid name",
		},
		{
			name:    "Maria",
			cpf:     "031.915.990-61",
			secret:  "12345",
			want:    ErrInvalidSecret,
			message: "Should not create an account due to a short secret",
		},
		{
			name:    "Maria",
			cpf:     "031.915.990-61",
			secret:  "123451234512345123451234512345123451234512345123451",
			want:    ErrInvalidSecret,
			message: "Should not create an account due to a too long secret",
		},
	}

	for _, tc := range tCases {
		t.Run(tc.message, func(t *testing.T) {
			_, err := NewAccount(tc.name, tc.cpf, tc.secret)
			assert.Equal(t, tc.want, err)
		})
	}

	t.Run("Should create an account successfully", func(t *testing.T) {
		name := "Maria"
		cpf := "031.915.990-61"
		secret := "123456"

		acc, err := NewAccount(name, cpf, secret)

		assert.Nil(t, err)
		// avoid null pointer error
		if err != nil {
			return
		}

		assert.Equal(t, acc.Name, name)
		assert.Equal(t, acc.CPF, cpf)
		assert.Equal(t, acc.Balance, 0)

		assert.NotEqual(t, acc.Secret, secret)
	})
}
