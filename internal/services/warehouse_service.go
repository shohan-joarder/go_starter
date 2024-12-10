package services

import (
	"github.com/shohan-joarder/go_pos/internal/models"
	"github.com/shohan-joarder/go_pos/internal/repositories"
)

type WarehouseService struct {
	repo *repositories.WarehouseRepository
}

func NewWarehouseService(repo *repositories.WarehouseRepository) *WarehouseService {
	return &WarehouseService{repo: repo}
}

func (s *WarehouseService) GetAllWarehouses() ([]models.Warehouse, error) {
	return s.repo.GetAllWarehouses()
}

func (s *WarehouseService) GetWarehouseByID(id uint) (*models.Warehouse, error) {
	return s.repo.GetWarehouseByID(id)
}

func (s *WarehouseService) GetWarehouseByUserID(userID string) (*models.Warehouse, error) {
	return s.repo.GetWarehouseByUserID(userID)
}

func (s *WarehouseService) CreateWarehouse(warehouse *models.Warehouse) error {
	return s.repo.CreateWarehouse(warehouse)
}

func (s *WarehouseService) UpdateWarehouse(warehouse *models.Warehouse) error {
	return s.repo.UpdateWarehouse(warehouse)
}

func (s *WarehouseService) DeleteWarehouse(warehouse *models.Warehouse) error {
	return s.repo.DeleteWarehouse(warehouse)
}
