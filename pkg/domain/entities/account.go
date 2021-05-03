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

var (
	ErrCPFAlreadyExists = errors.New("the cpf is already in use")
	ErrInvalidCPF       = errors.New("invalid cpf format. example of cpf: 601.647.540-83")
	ErrInvalidName      = errors.New("the name can't have more than 255 characters")
	ErrInvalidSecret    = errors.New("the secret must have a length between 6 and 50")
	AccountDomainErrors = []error{
		ErrCPFAlreadyExists,
		ErrInvalidCPF,
		ErrInvalidName,
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
	if len(a.Name) == 0 || len(a.Name) > 255 {
		return ErrInvalidName
	}
	if !cpf.Validate(a.CPF) {
		return ErrInvalidCPF
	}

	return nil
}
