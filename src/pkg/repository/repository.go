package repository

import (
	"avito-test/dto"
	"avito-test/model"
	"github.com/jmoiron/sqlx"
)

type UserAccount interface {
	CreateUserAccount(userDto dto.CreateUser) (int, error)
	FindUserAccount(id int) (model.UserAccount, error)
}

type Transaction interface {
	CreateTransaction(transaction dto.Transaction, trType model.TransactionType) (int, error)
}

type Repository struct {
	UserAccount
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		NewUserAccountPostgres(db),
		NewTransactionPostgres(db),
	}
}
