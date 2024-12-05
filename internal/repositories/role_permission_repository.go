package repositories

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shohan-joarder/go_pos/internal/models"
	"github.com/shohan-joarder/go_pos/internal/utils"
	"gorm.io/gorm"
)

type RolePermissionRepository struct {
	DB *gorm.DB
}

func NewRolePermissionRepository(db *gorm.DB) *RolePermissionRepository {
	return &RolePermissionRepository{DB: db}
}

func (r *RolePermissionRepository) GetRolePermissionByRoleID(role_id uint) (*models.RolePermission, error) {
	var rolePermission models.RolePermission
	err := r.DB.First(&rolePermission, role_id).Error
	return &rolePermission, err
}

func (r *RolePermissionRepository) CreateRolePermission(rolePermission *models.RolePermission) error {
	return r.DB.Create(rolePermission).Error
}

func (r *RolePermissionRepository) UpdateRolePermission(rolePermission *models.RolePermission) error {
	return r.DB.Save(rolePermission).Error
}

func (r *RolePermissionRepository) DeleteRolePermission(rolePermission *models.RolePermission) error {
	return r.DB.Delete(rolePermission).Error
}

// GetPermissionsByURLMethod filters the RolePermission table based on URL and Method in the permissions JSON
func (r *RolePermissionRepository) FilterPermissionByRoleIdAndURLMethod(roleId uint, url, method string) (bool, error) {
	// Fetch the role permission entry for the given role ID
	var rolePermission models.RolePermission
	err := r.DB.Where("role_id = ?", roleId).First(&rolePermission).Error
	if err != nil {
		return false, fmt.Errorf("failed to fetch role permissions for role ID %d: %w", roleId, err)
	}

	// Parse the permissions JSON stored in the database
	var permissions []models.Permission
	if err := json.Unmarshal([]byte(rolePermission.Permissions), &permissions); err != nil {
		return false, fmt.Errorf("failed to parse permissions JSON: %w", err)
	}

	// Iterate through permissions and check for a match
	for _, permission := range permissions {
		// Check if method matches
		if strings.EqualFold(permission.Methods, method) {
			// Check if the URL matches
			if utils.MatchURL(permission.URL, url) && permission.Status {
				return true, nil // Permission granted
			}
		}
	}

	// No matching permission found
	return false, nil
}
