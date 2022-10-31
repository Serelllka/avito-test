package repository

import (
	"avito-test/dto"
	"avito-test/model"
	"github.com/jmoiron/sqlx"
)

type ServicePostgres struct {
	db *sqlx.DB
}

func NewServicePostgres(db *sqlx.DB) *ServicePostgres {
	return &ServicePostgres{db}
}

func (r *ServicePostgres) CreateService(service dto.Service) (int, error) {
	return 0, nil
}

func (r *ServicePostgres) FindServiceById(id int) (model.Service, error) {
	return model.Service{}, nil
}
