package auth

import "context"

type Service interface {
	Authenticate(ctx context.Context, CPF, secret string) (string, error)
}
