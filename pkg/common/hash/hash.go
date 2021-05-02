package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// Creates a new hash
func New(secret string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
