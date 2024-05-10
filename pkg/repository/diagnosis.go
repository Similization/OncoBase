package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type DiagnosisRepository struct {
	db *sqlx.DB
}

func NewDiagnosisRepository(db *sqlx.DB) *DiagnosisRepository {
	return &DiagnosisRepository{db: db}
}

// Create diagnosis in database and get it from database
func (r *DiagnosisRepository) CreateDiagnosis(diagnosis model.Diagnosis) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var id interface{} = diagnosis.Id
	if diagnosis.Id == "" {
		id = nil
	}

	query := fmt.Sprintf("INSERT INTO %s (id, description) VALUES ($1, $2)", diagnosisTable)
	_, err = r.db.Exec(query,
		id,
		diagnosis.Description,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Get diagnosis list from database
func (r *DiagnosisRepository) GetDiagnosisList() ([]model.Diagnosis, error) {
	var diagnosisList []model.Diagnosis
	query := fmt.Sprintf("SELECT * FROM %s", diagnosisTable)
	err := r.db.Select(&diagnosisList, query)
	return diagnosisList, err
}

// Get diagnosis from database by ID
func (r *DiagnosisRepository) GetDiagnosisById(id string) (model.Diagnosis, error) {
	var diagnosis model.Diagnosis
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", diagnosisTable)
	err := r.db.Get(&diagnosis, query, id)
	return diagnosis, err
}

// Update diagnosis data in database
func (r *DiagnosisRepository) UpdateDiagnosis(diagnosis model.Diagnosis) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET description=$2 WHERE id=$1", diagnosisTable)
	_, err = r.db.Exec(query,
		diagnosis.Id,
		diagnosis.Description,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Delete diagnosis data from database
func (r *DiagnosisRepository) DeleteDiagnosis(id string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", diagnosisTable)
	_, err = r.db.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
