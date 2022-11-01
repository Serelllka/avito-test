package repository

import (
	"avito-test/dto"
	"avito-test/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ReservedTransactionModel struct {
	ProducerId uint `db:"producer_id"`
	ServiceId  uint `db:"service_id"`
	OrderId    uint `db:"order_id"`
	Amount     uint `db:"amount"`
}

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db}
}

func (r *TransactionPostgres) CreateRemittance(rem dto.Remittance) (int, error) {
	var balance model.UserAccountBalance
	var id int

	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersBalanceView)
	if err = tx.Get(&balance, query, rem.ProducerId); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if balance.Income < rem.Amount+balance.Outcome+balance.Reserved {
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

func (r *TransactionPostgres) CreateReservation(res dto.Reservation) error {
	var balance model.UserAccountBalance
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersBalanceView)
	if err = tx.Get(&balance, query, res.ProducerId); err != nil {
		_ = tx.Rollback()
		return err
	}

	if balance.Income < res.Amount+balance.Outcome+balance.Reserved {
		_ = tx.Rollback()
		return fmt.Errorf("not enough money")
	}

	query = fmt.Sprintf("INSERT INTO %s (producer_id, service_id, order_id, amount) "+
		"VALUES ($1, $2, $3, $4)", reservationsTable)

	if _, err = tx.Exec(query, res.ProducerId, res.ServiceId, res.OrderId, res.Amount); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *TransactionPostgres) CreatePayment(pay dto.Payment) (int, error) {
	var reserved ReservedTransactionModel
	var id int
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE producer_id = $1 AND service_id = $2 AND order_id = $3",
		reservationsTable,
	)

	if err = tx.Get(&reserved, query, pay.ProducerId, pay.ServiceId, pay.OrderId); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf(
		"UPDATE %s SET amount = $1 WHERE producer_id = $2 AND service_id = $3 AND order_id = $4",
		reservationsTable,
	)

	if reserved.Amount < pay.Amount {
		_ = tx.Rollback()
		return 0, fmt.Errorf("not enough money")
	}

	if _, err = tx.Exec(query, reserved.Amount-pay.Amount, pay.ProducerId, pay.ServiceId, pay.OrderId); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (transaction_type, producer_id, service_id, order_id, amount, description) "+
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", usersTransactionTable)

	row := tx.QueryRow(query, model.Payment, pay.ProducerId, pay.ServiceId, pay.OrderId, pay.Amount, pay.Description)

	if err := row.Scan(&id); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
