package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/matheusmosca/simple-bank/pkg/common/cpf"
	"github.com/matheusmosca/simple-bank/pkg/common/hash"
)

func newID() string {
	return uuid.NewString()
}

const DefaultBalanceValue = 1000

var (
	ErrCPFAlreadyExists    = errors.New("the cpf is already in use")
	ErrInvalidCPF          = errors.New("invalid cpf format. example of cpf: 601.647.540-83")
	ErrInvalidName         = errors.New("the name can't have more than 255 characters")
	ErrInvalidSecret       = errors.New("the secret must have a length between 6 and 50")
	ErrAccountDoesNotExist = errors.New("this account does not exist")

	AccountDomainErrors = []error{
		ErrCPFAlreadyExists,
		ErrInvalidCPF,
		ErrInvalidName,
		ErrInvalidSecret,
		ErrAccountDoesNotExist,
	}
)

type CreateAccountInput struct {
	Name   string
	CPF    string
	Secret string
}

type Account struct {
	ID        string
	Name      string
	CPF       string
	Secret    string
	Balance   int
	CreatedAt time.Time
}

func NewAccount(name, CPF, secret string) (*Account, error) {
	acc := Account{
		ID:      newID(),
		Name:    name,
		CPF:     CPF,
		Secret:  secret,
		Balance: DefaultBalanceValue,
	}
	err := acc.Validate()
	hash, _ := hash.New(acc.Secret)

	acc.Secret = hash

	if err != nil {
		return nil, err
	}

	return &acc, nil
}

func (a Account) Validate() error {
	if len(a.Secret) < 6 || len(a.Secret) > 50 {
		return ErrInvalidSecret
	}
	if len(a.Name) == 0 || len(a.Name) > 255 {
		return ErrInvalidName
	}
	if !cpf.Validate(a.CPF) {
		return ErrInvalidCPF
	}

	return nil
}

// Checks if an accounts has sufficient balance
// to perform an transaction with the provided amount
func (a Account) CheckWalletFunds(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if (a.Balance - amount) < 0 {
		return ErrInsufficientFunds
	}

	return nil
}

func (a *Account) DepositMoney(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	a.Balance += amount
	return nil
}

func (a *Account) WithdrawMoney(amount int) error {
	err := a.CheckWalletFunds(amount)
	if err != nil {
		return err
	}

	a.Balance -= amount
	return nil
}
