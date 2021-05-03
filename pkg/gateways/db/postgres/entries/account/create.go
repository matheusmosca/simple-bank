package entries

import (
	"context"
	"simple-bank/pkg/domain/entities"
)

func (r Repository) Create(ctx context.Context, account *entities.Account) error {
	statement := `INSERT INTO account
							(id,
							 name,
							 cpf,
							 secret,
							 balance)
							VALUES ($1, $2, $3, $4, $5)
							returning created_at
							`

	err := r.DB.QueryRowContext(ctx, statement,
		account.ID,
		account.Name,
		account.CPF,
		account.Secret,
		account.Balance,
	).Scan(&account.CreatedAt)

	return err
}

func (r Repository) GetByID(ctx context.Context, accountID string) (*entities.Account, error) {
	return nil, nil
}

func (r Repository) GetByCPF(ctx context.Context, CPF string) (*entities.Account, error) {
	return nil, nil
}

func (r Repository) GetAccounts(ctx context.Context) ([]*entities.Account, error) {
	return nil, nil
}

func (r Repository) UpdateBalance(ctx context.Context, ID string) error {
	return nil
}
