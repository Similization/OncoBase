package services

import (
	"med/pkg/model"
	"med/pkg/repository"
	"med/pkg/utils"
)

type AuthorizationService struct {
	repo repository.Authorization
	salt []byte
}

func NewAuthService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo, salt: utils.Salt}
}

func (s *AuthorizationService) CreateUser(user model.User) (string, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthorizationService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (*AuthorizationService) ParseToken(token string) (string, error) {
	claims, err := utils.ParseToken(token)
	if err != nil {
		return "", err
	}

	return claims.UserEmail, nil
}

func (s *AuthorizationService) generatePasswordHash(password string) string {
	return utils.HashPassword(password, s.salt)
}
