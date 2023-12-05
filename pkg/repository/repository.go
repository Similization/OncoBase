package repository

import (
	server "med"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user server.User) (int, error)
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
