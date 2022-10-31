package service

import (
	"avito-test/dto"
	"avito-test/pkg/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo}
}

func (s *TransactionService) CreateRemittance(transaction dto.Remittance) (int, error) {
	return s.repo.CreateRemittance(transaction)
}

func (s *TransactionService) CreateDeposit(transaction dto.Deposit) (int, error) {
	return s.repo.CreateDeposit(transaction)
}

func (s *TransactionService) CreateReservation(transaction dto.Reservation) (int, error) {
	return s.repo.CreateReservation(transaction)
}

func (s *TransactionService) CreatePayment(transaction dto.Reservation) (int, error) {
	return s.repo.CreatePayment(transaction)
}
