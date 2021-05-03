package account

import "simple-bank/pkg/domain/entities"

// Check if some error belongs to the account domain
func IsDomainError(err error) bool {
	for _, e := range entities.AccountDomainErrors {
		if e == err {
			return true
		}
	}

	return false
}
