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
	CreateRemittance(transaction dto.Remittance) (int, error)
	CreateDeposit(transaction dto.Deposit) (int, error)

	CreateReservation(transaction dto.Reservation) (int, error)
	CreatePayment(transaction dto.Reservation) (int, error)
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
