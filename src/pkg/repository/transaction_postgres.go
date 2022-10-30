package repository

import (
	"avito-test/dto"
	"avito-test/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db}
}

func (r *TransactionPostgres) CreateTransaction(transaction dto.Transaction, trType model.TransactionType) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (producer_id, consumer_id, transaction_type, amount, description) "+
		"VALUES ($1, $2, $3, $4, $5) RETURNING id", usersTransactionTable)

	row := r.db.QueryRow(
		query,
		transaction.ProducerId,
		transaction.ConsumerId,
		trType,
		transaction.Amount,
		transaction.Description,
	)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
