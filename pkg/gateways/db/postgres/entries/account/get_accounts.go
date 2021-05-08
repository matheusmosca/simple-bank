package account

import (
	"context"
	"simple-bank/pkg/domain/entities"
)

func (r Repository) GetAccounts(ctx context.Context) ([]entities.Account, error) {
	query := `
		SELECT 
			id,
			name,
			cpf,
			balance,
			created_at
		FROM account
	`
	var accounts []entities.Account

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var acc entities.Account

		err := rows.Scan(
			&acc.ID,
			&acc.Name,
			&acc.CPF,
			&acc.Balance,
			&acc.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, acc)
	}

	return accounts, nil
}
