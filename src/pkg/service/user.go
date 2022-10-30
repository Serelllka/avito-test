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

func (s *UserService) GetUserById(id int) (dto.UserAccount, error) {
	userModel, err := s.repo.FindUserAccount(id)
	if err != nil {
		log.Printf("error while finding user from db: %s", err.Error())
		return dto.UserAccount{}, err
	}

	return dto.UserAccount{
		Id:   userModel.Id,
		Name: userModel.Name,
	}, nil
}

func (s *UserService) GetAllUsers() ([]dto.UserAccount, error) {
	return nil, nil
}
