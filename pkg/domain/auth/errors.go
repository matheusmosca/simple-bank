package auth

import "errors"

var (
	ErrWrongCredentials = errors.New("wrong cpf or secret")
)
