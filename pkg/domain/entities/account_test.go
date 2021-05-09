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
		assert.Equal(t, acc.Balance, DefaultBalanceValue)

		assert.NotEqual(t, acc.Secret, secret)
	})
}
func TestDepositMoney(t *testing.T) {
	acc, _ := NewAccount("jorge", "031.915.990-61", "123784397")

	t.Run("the account should recieve money normally", func(t *testing.T) {
		acc.Balance = 10

		want := acc.Balance + 300
		acc.DepositMoney(300)

		assert.Equal(t, acc.Balance, want)
	})

	t.Run("should return a error due to an invalid amount provided", func(t *testing.T) {
		err := acc.DepositMoney(0)
		assert.Equal(t, err, ErrInvalidAmount)
	})
}

func TestWithdrawMoney(t *testing.T) {
	t.Run("the account's wallet should withdraw normally", func(t *testing.T) {
		acc, _ := NewAccount("jorge", "031.915.990-61", "123784397")

		err := acc.WithdrawMoney(DefaultBalanceValue)

		assert.Nil(t, err)
		assert.Equal(t, acc.Balance, 0)
	})

	t.Run("should return a error due to an invalid amount provided", func(t *testing.T) {
		acc, _ := NewAccount("jorge", "031.915.990-61", "123784397")
		err := acc.WithdrawMoney(0)
		assert.Equal(t, err, ErrInvalidAmount)
		// The balance should not be modified
		assert.Equal(t, acc.Balance, DefaultBalanceValue)
	})

	t.Run("should return a error due to insuficient funds", func(t *testing.T) {
		// The account has the DefaultBalanceValue
		acc, _ := NewAccount("jorge", "031.915.990-61", "123784397")

		err := acc.WithdrawMoney(DefaultBalanceValue + 1)
		assert.Equal(t, err, ErrInsufficientFunds)
		// The balance should not be modified
		assert.Equal(t, acc.Balance, DefaultBalanceValue)
	})
}

func TestAccountWalletFunds(t *testing.T) {
	acc, _ := NewAccount("jorge", "031.915.990-61", "123784397")

	t.Run("the account should have sufficient funds", func(t *testing.T) {
		err := acc.CheckWalletFunds(DefaultBalanceValue)
		assert.Nil(t, err)
	})

	t.Run("the account should not have sufficient funds", func(t *testing.T) {
		err := acc.CheckWalletFunds(DefaultBalanceValue + 1)
		assert.Equal(t, err, ErrInsufficientFunds)
	})

	t.Run("should return a error due to an invalid amount provided", func(t *testing.T) {
		err := acc.CheckWalletFunds(0)
		assert.Equal(t, err, ErrInvalidAmount)
	})
}
