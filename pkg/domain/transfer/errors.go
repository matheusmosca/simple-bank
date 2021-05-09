package transfer

import "simple-bank/pkg/domain/entities"

// Check if some error belongs to the transfer domain
func IsDomainError(err error) bool {
	for _, e := range entities.TransferDomainErrors {
		if e == err {
			return true
		}
	}

	return false
}
