package wallet

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func NewWalletRepository(DB *sql.DB) *Repository {
	return &Repository{
		DB: DB,
	}
}
