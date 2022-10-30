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

func (r *UserAccountPostgres) FindUserAccount(id int) (model.UserAccount, error) {
	return model.UserAccount{}, nil
}
