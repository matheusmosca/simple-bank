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

func CompareSecrets(secret string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	if err != nil {
		return false, err
	}

	return true, nil
}
