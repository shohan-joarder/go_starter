package repositories

import (
	"github.com/shohan-joarder/go_pos/internal/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (r *RoleRepository) GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	err := r.DB.Find(&roles).Error
	return roles, err
}

func (r *RoleRepository) CreateRole(role *models.Role) error {
	return r.DB.Create(role).Error
}

func (r *RoleRepository) GetRoleByID(id uint) (*models.Role, error) {
	var role models.Role
	err := r.DB.First(&role, id).Error
	return &role, err
}

func (r *RoleRepository) UpdateRole(role *models.Role) error {
	return r.DB.Save(role).Error
}

func (r *RoleRepository) DeleteRole(role *models.Role) error {
	return r.DB.Delete(role).Error
}
