package repository

import (
	"avito-test/dto"
	"avito-test/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ServicePostgres struct {
	db *sqlx.DB
}

func NewServicePostgres(db *sqlx.DB) *ServicePostgres {
	return &ServicePostgres{db}
}

func (r *ServicePostgres) CreateService(service dto.Service) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", servicesTable)

	row := r.db.QueryRow(query, service.Title, service.Description)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ServicePostgres) FindServiceById(id int) (model.Service, error) {
	return model.Service{}, nil
}
