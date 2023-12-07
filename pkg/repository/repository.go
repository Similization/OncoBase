package repository

import (
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (string, error)
	GetUser(email, password string) (model.User, error)
}

type Account interface {
}

type BloodCount interface {
}

type Repository struct {
	Authorization
	Account
	BloodCount
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
