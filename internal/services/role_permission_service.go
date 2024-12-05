package services

import (
	"fmt"

	"github.com/shohan-joarder/go_pos/internal/models"
	"github.com/shohan-joarder/go_pos/internal/repositories"
)

type RolePermissionService struct {
	repo *repositories.RolePermissionRepository
}

func NewRolePermissionService(repo *repositories.RolePermissionRepository) *RolePermissionService {
	return &RolePermissionService{repo: repo}
}

func (s *RolePermissionService) GetRolePermissionByRoleID(role_id uint) (*models.RolePermission, error) {
	return s.repo.GetRolePermissionByRoleID(role_id)
}

func (s *RolePermissionService) CreateRolePermission(rolePermission *models.RolePermission) error {
	return s.repo.CreateRolePermission(rolePermission)
}

func (s *RolePermissionService) UpdateRolePermission(rolePermission *models.RolePermission) error {
	return s.repo.UpdateRolePermission(rolePermission)
}

func (s *RolePermissionService) FilterPermissionByRoleIdAndURLMethod(roleId uint, url, method string) (bool, error) {
	// Call repository to filter permissions
	isAllowed, err := s.repo.FilterPermissionByRoleIdAndURLMethod(roleId, url, method)
	if err != nil {
		return false, fmt.Errorf("error fetching permissions: %w", err)
	}

	fmt.Printf("Permission check result: roleID=%d, url=%s, method=%s, allowed=%v\n", roleId, url, method, isAllowed)

	return isAllowed, nil
}
