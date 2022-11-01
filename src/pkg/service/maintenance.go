package service

import (
	"avito-test/dto"
	"avito-test/pkg/repository"
)

type MaintenanceService struct {
	repo repository.Service
}

func NewMaintenanceService(repo repository.Service) *MaintenanceService {
	return &MaintenanceService{
		repo: repo,
	}
}

func (r *MaintenanceService) CreateService(service dto.Service) (int, error) {
	return r.repo.CreateService(service)
}
