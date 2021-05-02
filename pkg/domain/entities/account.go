package entities

import (
	"errors"
	"simple-bank/pkg/common/hash"
	"time"

	"github.com/google/uuid"
)

func newID() string {
	return uuid.NewString()
}

var (
	ErrCPFAlreadyExists = errors.New("the cpf is already in use")
	ErrInvalidCPF       = errors.New("invalid cpf")
	ErrInvalidSecret    = errors.New("the secret must have a length between 6 and 50")
	DomainErrors        = []error{
		ErrCPFAlreadyExists,
		ErrInvalidCPF,
		ErrInvalidSecret,
	}
)

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
		Balance: 0,
	}
	err := acc.Validate()
	hash, _ := hash.New(acc.Secret)

	acc.Secret = hash

	if err != nil {
		return nil, err
	}

	return &acc, nil
}

// TODO add proper validation
// TODO validate cpf
// TODO hash secret with bcrypt
func (a Account) Validate() error {
	if len(a.Secret) < 6 || len(a.Secret) > 50 {
		return ErrInvalidSecret
	}
	return nil
}
