package services

import (
	server "med"
	"med/pkg/repository"
)

type Authorization interface {
	CreateUser(user server.User) (int, error)
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
