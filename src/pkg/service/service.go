package service

import (
	"avito-test/dto"
	"avito-test/model"
	"avito-test/pkg/repository"
)

type UserAccount interface {
	CreateUserAccount(user dto.CreateUser) (int, error)
	GetUserById(id int) (dto.UserAccount, error)
	GetAllUsers() ([]dto.UserAccount, error)
}

type Transaction interface {
	CreateTransaction(transaction dto.Transaction, trType model.TransactionType) (int, error)
	CreateRemittance(transaction dto.Transaction) (int, error)
	CreateDeposit(transaction dto.Transaction) (int, error)
	CreateReservation(transaction dto.Transaction) (int, error)
}

type Service struct {
	UserAccount
	Transaction
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		NewUserAccountService(repo.UserAccount),
		NewTransactionService(repo.Transaction),
	}
}
