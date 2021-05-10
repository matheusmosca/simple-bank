package transfer

import (
	"context"

	"github.com/matheusmosca/simple-bank/pkg/domain/entities"
)

func (r Repository) ListTransfersByAccountID(ctx context.Context, ID string) ([]entities.Transfer, error) {
	statement := `
		SELECT 
			id,
			account_origin_id,
			account_destination_id,
			amount,
			created_at
		FROM 
			transfer
		WHERE account_origin_id=$1 OR account_destination_id=$1`

	rows, err := r.DB.QueryContext(ctx, statement, ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transfers := []entities.Transfer{}

	for rows.Next() {
		var t entities.Transfer
		rows.Scan(
			&t.ID,
			&t.AccountOriginID,
			&t.AccountDestinationID,
			&t.Amount,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, t)
	}

	return transfers, nil
}
