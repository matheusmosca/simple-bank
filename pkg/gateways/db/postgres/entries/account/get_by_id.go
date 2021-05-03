package account

import (
	"context"
	"simple-bank/pkg/domain/entities"
)

func (r Repository) GetByID(ctx context.Context, accountID string) (*entities.Account, error) {
	query := `
		SELECT
			id,
			name,
			cpf,
			balance,
			created_at
		FROM
			account
		WHERE
			id=$1`

	var acc entities.Account

	err := r.DB.QueryRowContext(ctx, query, accountID).Scan(
		&acc.ID,
		&acc.Name,
		&acc.CPF,
		&acc.Balance,
		&acc.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &acc, nil
}
