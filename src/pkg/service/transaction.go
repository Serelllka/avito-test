package service

import (
	"avito-test/dto"
	"avito-test/model"
	"avito-test/pkg/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo}
}

func (s *TransactionService) CreateTransaction(transaction dto.Transaction, trType model.TransactionType) (int, error) {
	return s.repo.CreateTransaction(transaction, trType)
}

func (s *TransactionService) CreateRemittance(transaction dto.Transaction) (int, error) {
	return s.repo.CreateRemittance(transaction)
}

func (s *TransactionService) CreateDeposit(transaction dto.Transaction) (int, error) {
	return s.repo.CreateDeposit(transaction)
}

func (s *TransactionService) CreateReservation(transaction dto.Transaction) (int, error) {
	return s.repo.CreateReservation(transaction)
}
