package account

import (
	"database/sql"

	"github.com/matheusmosca/simple-bank/pkg/domain/account"
)

var _ account.Repository = Repository{}

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
