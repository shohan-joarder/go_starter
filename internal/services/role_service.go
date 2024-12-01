package services

import (
	"github.com/shohan-joarder/go_pos/internal/models"
	"github.com/shohan-joarder/go_pos/internal/repositories"
)

type RoleService struct {
	repo *repositories.RoleRepository
}

func NewRoleService(repo *repositories.RoleRepository) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	return s.repo.GetAllRoles()
}

func (s *RoleService) CreateRole(role *models.Role) error {
	return s.repo.CreateRole(role)
}

func (s *RoleService) GetRoleByID(id uint) (*models.Role, error) {
	return s.repo.GetRoleByID(id)
}

func (s *RoleService) UpdateRole(role *models.Role) error {
	return s.repo.UpdateRole(role)
}

func (s *RoleService) DeleteRole(role *models.Role) error {
	return s.repo.DeleteRole(role)
}
