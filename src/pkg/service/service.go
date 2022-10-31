package service

import (
	"avito-test/dto"
	"avito-test/pkg/repository"
)

type UserAccount interface {
	CreateUserAccount(user dto.CreateUser) (int, error)
	GetUserById(id int) (dto.UserAccount, error)
	GetAllUsers() ([]dto.UserAccount, error)
}

type Transaction interface {
	CreateRemittance(transaction dto.Remittance) (int, error)
	CreateDeposit(transaction dto.Deposit) (int, error)
	CreateReservation(transaction dto.Reservation) (int, error)
}

type Maintenance interface {
	CreateService(service dto.Service) (int, error)
}

type Service struct {
	Maintenance
	UserAccount
	Transaction
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		NewMaintenanceService(repo.Service),
		NewUserAccountService(repo.UserAccount),
		NewTransactionService(repo.Transaction),
	}
}
