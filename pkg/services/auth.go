package services

import (
	server "med"
	"med/pkg/repository"
	"med/pkg/utils"
)

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService) CreateUser(user server.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthorizationService) generatePasswordHash(password string) string {
	salt := utils.GenerateRandomSalt(0)
	return utils.HashPassword(password, salt)
}
