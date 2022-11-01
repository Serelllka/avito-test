package service

import (
	"avito-test/dto"
	"avito-test/pkg/repository"
	"log"
)

type UserService struct {
	repo repository.UserAccount
}

func NewUserAccountService(repo repository.UserAccount) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUserAccount(user dto.CreateUser) (int, error) {
	id, err := s.repo.CreateUserAccount(user)
	if err != nil {
		log.Printf("error while creating user: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (s *UserService) GetUserAccountBalanceById(id int) (dto.UserAccountBalance, error) {
	userModel, err := s.repo.GetUserAccountBalance(id)
	if err != nil {
		log.Printf("error while finding user from db: %s", err.Error())
		return dto.UserAccountBalance{}, err
	}

	return dto.UserAccountBalance{
		Balance: userModel.Income - userModel.Outcome - userModel.Reserved,
	}, nil
}

func (s *UserService) GetAllUsers() ([]dto.UserAccount, error) {
	return nil, nil
}
