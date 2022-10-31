package repository

import (
	"avito-test/dto"
	"avito-test/model"
	"github.com/jmoiron/sqlx"
)

type Service interface {
	CreateService(service dto.Service) (int, error)
	FindServiceById(id int) (model.Service, error)
}

type UserAccount interface {
	CreateUserAccount(userDto dto.CreateUser) (int, error)
	FindUserAccount(id int) (model.UserAccount, error)
}

type Transaction interface {
	CreateTransaction(transaction dto.Transaction, trType model.TransactionType) (int, error)
	CreateRemittance(transaction dto.Transaction) (int, error)
	CreateDeposit(transaction dto.Transaction) (int, error)
	CreateReservation(transaction dto.Transaction) (int, error)
}

type Repository struct {
	Service
	UserAccount
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		NewServicePostgres(db),
		NewUserAccountPostgres(db),
		NewTransactionPostgres(db),
	}
}
