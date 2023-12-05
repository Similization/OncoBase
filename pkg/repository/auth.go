package repository

import (
	"fmt"
	server "med"

	"github.com/jmoiron/sqlx"
)

type AuthorizationRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db: db}
}

func (r *AuthorizationRepository) CreateUser(user server.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password) VALUES ($1, $2) RETURNING id", patientTable)
	row := r.db.QueryRow(query, user.Email, user.Password)

	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
