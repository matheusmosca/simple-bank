package entities

import (
	"errors"
	"simple-bank/pkg/common/cpf"
	"simple-bank/pkg/common/hash"
	"time"

	"github.com/google/uuid"
)

func newID() string {
	return uuid.NewString()
}

const DefaultBalanceValue = 0

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

func (a Account) DisplayBalance() float64 {
	return float64(a.Balance) / 100
}
