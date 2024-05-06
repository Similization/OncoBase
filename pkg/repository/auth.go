package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type AuthorizationRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db: db}
}

func (r *AuthorizationRepository) CreateUser(user model.User) (string, error) {
	var email string
	query := fmt.Sprintf("INSERT INTO %s (email, password, role) VALUES ($1, $2, $3) RETURNING email", externalUserTable)
	row := r.db.QueryRow(query, user.Email, user.Password, user.Role)

	if err := row.Scan(&email); err != nil {
		return "", err
	}
	return email, nil
}

func (r *AuthorizationRepository) GetUser(email, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 AND password=$2", internalUserTable)
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 AND password=$2", externalUserTable)
		err = r.db.Get(&user, query, email, password)
	}
	return user, err
}

// func (r *AuthorizationRepository) GetUser(email, password string) (model.User, error) {
// 	var user model.User

// 	// Create a squirrel.SelectBuilder for the internal user table
// 	internalSelect := squirrel.Select("*").From(internalUserTable).Where(squirrel.Eq{"email": email, "password": password})

// 	// Create a squirrel.SelectBuilder for the external user table
// 	externalSelect := squirrel.Select("*").From(externalUserTable).Where(squirrel.Eq{"email": email, "password": password})

// 	// Combine the two queries with OR
// 	query, args, err := squirrel.Or{internalSelect, externalSelect}.ToSql()
// 	if err != nil {
// 		return user, err
// 	}

// 	// Execute the query
// 	err = r.db.Get(&user, query, args...)
// 	return user, err
// }
