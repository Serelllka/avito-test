package repository

import (
	"avito-test/dto"
	"avito-test/model"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AccountBalanceModel struct {
	Id      uint `db:"id"`
	Income  uint `db:"income"`
	Outcome uint `db:"outcome"`
}

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db}
}

func (r *TransactionPostgres) CreateTransaction(transaction dto.Transaction, trType model.TransactionType) (int, error) {
	switch trType {
	case model.Remittance:
		return r.CreateRemittance(transaction)
	case model.Deposit:
		return r.CreateDeposit(transaction)
	case model.Reservation:
		return r.CreateReservation(transaction)
	default:
		return 0, fmt.Errorf("unresolved transaction Type: %d", trType)
	}
}

func (r *TransactionPostgres) CreateRemittance(transaction dto.Transaction) (int, error) {
	var balance AccountBalanceModel
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("SELECT id, income, outcome FROM %s WHERE id = $1", usersBalanceView)
	if err = tx.Get(&balance, query, transaction.ProducerId); err != nil {
		fmt.Println("Im probably here!")
		_ = tx.Rollback()
		return 0, err
	}

	if balance.Income-balance.Outcome < transaction.Amount {
		_ = tx.Rollback()
		return 0, fmt.Errorf("not enough money")
	}

	id, err := executeQuery(
		tx,
		fmt.Sprintf("INSERT INTO %s (transaction_type, producer_id, consumer_id, amount, description) "+
			"VALUES ($1, $2, $3, $4, $5) RETURNING id", usersTransactionTable),
		model.Remittance,
		transaction,
	)
	return id, tx.Commit()
}

func (r *TransactionPostgres) CreateDeposit(transaction dto.Transaction) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (transaction_type, consumer_id, amount, description) "+
		"VALUES ($1, $2, $3, $4) RETURNING id", usersTransactionTable)

	return executeQuery(r.db, query, model.Deposit, transaction)
}

func (r *TransactionPostgres) CreateReservation(transaction dto.Transaction) (int, error) {
	var balance AccountBalanceModel
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersBalanceView)
	if err = tx.Get(&balance, query, transaction.ProducerId); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if balance.Income-balance.Outcome < transaction.Amount {
		_ = tx.Rollback()
		return 0, fmt.Errorf("not enough money")
	}

	id, err := executeQuery(
		tx,
		fmt.Sprintf("INSERT INTO %s (transaction_type, producer_id, service_id, amount, description) "+
			"VALUES ($1, $2, $3, $4, $5) RETURNING id", usersTransactionTable),
		model.Reservation,
		transaction,
	)

	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

type queryExecutor interface {
	QueryRow(query string, args ...any) *sql.Row
}

func executeQuery(db queryExecutor, query string, trType model.TransactionType, tr dto.Transaction) (int, error) {
	var id int

	var row *sql.Row
	switch trType {
	case model.Deposit:
		row = db.QueryRow(query, trType, tr.ConsumerId, tr.Amount, tr.Description)
	case model.Remittance:
		row = db.QueryRow(query, trType, tr.ProducerId, tr.ServiceId, tr.ConsumerId, tr.Amount, tr.Description)
	case model.Reservation:
		row = db.QueryRow(query, trType, tr.ProducerId, tr.Amount, tr.Description)
	default:
		return 0, fmt.Errorf("unsupported transaction type")
	}

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
