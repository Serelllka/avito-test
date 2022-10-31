package repository

import (
	"avito-test/dto"
	"avito-test/model"
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

func (r *TransactionPostgres) CreateRemittance(rem dto.Remittance) (int, error) {
	var balance AccountBalanceModel
	var id int

	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("SELECT id, income, outcome FROM %s WHERE id = $1", usersBalanceView)
	if err = tx.Get(&balance, query, rem.ProducerId); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if balance.Income-balance.Outcome < rem.Amount {
		_ = tx.Rollback()
		return 0, fmt.Errorf("not enough money")
	}

	query = fmt.Sprintf("INSERT INTO %s (transaction_type, producer_id, consumer_id, amount, description) "+
		"VALUES ($1, $2, $3, $4, $5) RETURNING id", usersTransactionTable)

	row := tx.QueryRow(query, model.Remittance, rem.ProducerId, rem.ConsumerId, rem.Amount, rem.Description)

	if err := row.Scan(&id); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TransactionPostgres) CreateDeposit(dep dto.Deposit) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (transaction_type, consumer_id, amount, description) "+
		"VALUES ($1, $2, $3, $4) RETURNING id", usersTransactionTable)

	row := r.db.QueryRow(query, model.Deposit, dep.ConsumerId, dep.Amount, dep.Description)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TransactionPostgres) CreateReservation(res dto.Reservation) (int, error) {
	var balance AccountBalanceModel
	var id int
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersBalanceView)
	if err = tx.Get(&balance, query, res.ProducerId); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if balance.Income-balance.Outcome < res.Amount {
		_ = tx.Rollback()
		return 0, fmt.Errorf("not enough money")
	}

	query = fmt.Sprintf("INSERT INTO %s (transaction_type, producer_id, service_id, amount, description, order_id) "+
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", usersTransactionTable)

	row := tx.QueryRow(query, model.Reservation, res.ProducerId, res.ServiceId, res.Amount, res.Description, res.OrderId)

	if err := row.Scan(&id); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
