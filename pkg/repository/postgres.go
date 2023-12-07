package repository

import (
	"med/pkg/config"

	"github.com/jmoiron/sqlx"
)

const (
	userTable = "onco_base.external_user"
)

func NewPostgresDB(cfg *config.ConfigDatabase) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.GetDataSourceName())

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
