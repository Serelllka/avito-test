package repository

import (
	"avito-test/dto"
	"avito-test/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserAccountPostgres struct {
	db *sqlx.DB
}

func NewUserAccountPostgres(db *sqlx.DB) *UserAccountPostgres {
	return &UserAccountPostgres{db}
}

func (r *UserAccountPostgres) CreateUserAccount(userDto dto.CreateUser) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", usersTable)

	row := r.db.QueryRow(query, userDto.Name)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserAccountPostgres) GetUserAccountBalance(id int) (model.UserAccountBalance, error) {
	var balance model.UserAccountBalance

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersBalanceView)
	if err := r.db.Get(&balance, query, id); err != nil {
		return model.UserAccountBalance{}, err
	}

	return balance, nil
}
