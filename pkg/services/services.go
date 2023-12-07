package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (string, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Account interface {
}

type BloodCount interface {
}

type Service struct {
	Authorization
	Account
	BloodCount
}

func NewService(repos repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
