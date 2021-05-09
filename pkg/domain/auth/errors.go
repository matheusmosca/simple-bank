package auth

import "errors"

var (
	ErrWrongCredentials = errors.New("wrong cpf or secret")
	ErrInvalidToken     = errors.New("invalid token")
	ErrTokenNotProvided = errors.New("token not provided")
)
