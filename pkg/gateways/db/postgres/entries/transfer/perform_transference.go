package transfer

import (
	"context"
	"database/sql"
	"log"
	"simple-bank/pkg/domain/entities"
)

func (r Repository) PerformTransference(ctx context.Context, input entities.PerformTransferenceInput) error {
	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Println(err)
		return err
	}

	defer func() error {
		if err != nil {
			rollErr := tx.Rollback()
			if rollErr != nil {
				log.Println(err)
			}
			return err
		}
		return nil
	}()

	err = r.updateAccountBalance(ctx, tx, *input.OriginAcount)
	if err != nil {
		log.Println(err)
		return err
	}

	err = r.updateAccountBalance(ctx, tx, *input.DestinationAcount)
	if err != nil {
		log.Println(err)
		return err
	}

	err = r.saveTransfer(ctx, tx, input.Transfer)
	if err != nil {
		log.Panicln(err)
		return err
	}

	tx.Commit()

	return nil
}

func (r Repository) updateAccountBalance(ctx context.Context, tx *sql.Tx, acc entities.Account) error {
	statement := `
		UPDATE
			account
		SET
			balance=$1
		WHERE 
			id=$2`

	_, err := tx.ExecContext(ctx, statement, acc.Balance, acc.ID)

	return err
}

func (r Repository) saveTransfer(ctx context.Context, tx *sql.Tx, trans *entities.Transfer) error {
	statement := `
		INSERT INTO transfer
			(id,
				account_origin_id,
				account_destination_id,
				amount)
			Values ($1, $2, $3, $4)
		returning created_at`

	err := tx.QueryRowContext(
		ctx,
		statement,
		trans.ID,
		trans.AccountOriginID,
		trans.AccountDestinationID,
		trans.Amount,
	).Scan(&trans.CreatedAt)

	return err
}
