package services

import (
	"github.com/shohan-joarder/go_pos/internal/models"
	"github.com/shohan-joarder/go_pos/internal/repositories"
)

type AuthService struct {
	repo *repositories.AuthRepository
}

func NewAuthService(repo *repositories.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(user *models.LoginUser) (string, error) {

	return s.repo.Login(user)

}
