package entities

import (
	"time"

	"github.com/google/uuid"
)

func newID() string {
	return uuid.NewString()
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
		Balance: 0,
	}
	err := acc.Validate()

	if err != nil {
		return nil, err
	}

	return &acc, nil
}

// TODO add proper validation
// TODO validate cpf
// TODO hash secret with bcrypt
func (a Account) Validate() error {
	return nil
}
