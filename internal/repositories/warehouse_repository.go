package repositories

import (
	"github.com/shohan-joarder/go_pos/internal/models"
	"gorm.io/gorm"
)

type WarehouseRepository struct {
	DB *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) *WarehouseRepository {
	return &WarehouseRepository{DB: db}
}

func (r *WarehouseRepository) GetAllWarehouses() ([]models.Warehouse, error) {
	var warehouses []models.Warehouse
	err := r.DB.Find(&warehouses).Error
	return warehouses, err
}

func (r *WarehouseRepository) CreateWarehouse(warehouse *models.Warehouse) error {
	return r.DB.Create(warehouse).Error
}

func (r *WarehouseRepository) GetWarehouseByID(id uint) (*models.Warehouse, error) {
	var warehouse models.Warehouse
	err := r.DB.First(&warehouse, id).Error
	return &warehouse, err
}

func (r *WarehouseRepository) UpdateWarehouse(warehouse *models.Warehouse) error {
	return r.DB.Save(warehouse).Error
}

func (r *WarehouseRepository) DeleteWarehouse(warehouse *models.Warehouse) error {
	return r.DB.Delete(warehouse).Error
}

func (r *WarehouseRepository) GetWarehouseByUserID(userID string) (*models.Warehouse, error) {
	var warehouse models.Warehouse
	err := r.DB.First(&warehouse, "user_id = ?", userID).Error
	return &warehouse, err
}
