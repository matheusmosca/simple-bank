package transfer

import (
	"database/sql"

	"github.com/matheusmosca/simple-bank/pkg/domain/transfer"
)

var _ transfer.Repository = Repository{}

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
