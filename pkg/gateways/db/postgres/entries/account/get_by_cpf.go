package account

import (
	"context"
	"simple-bank/pkg/domain/entities"
)

func (r Repository) GetByCPF(ctx context.Context, CPF string) (*entities.Account, error) {
	query := `
		SELECT
			id,
			name,
			cpf,
			balance,
			secret,
			created_at
		FROM
			account
		WHERE
			cpf=$1`

	var acc entities.Account

	err := r.DB.QueryRowContext(ctx, query, CPF).Scan(
		&acc.ID,
		&acc.Name,
		&acc.CPF,
		&acc.Balance,
		&acc.Secret,
		&acc.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &acc, nil
}
