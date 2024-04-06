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

// Create diagnosis in database and get him from database
func (r *DiagnosisRepository) CreateDiagnosis(diagnosis model.Diagnosis) (model.Diagnosis, error) {
	var createdDiagnosis model.Diagnosis
	query := fmt.Sprintf("INSERT INTO %s (id, description) VALUES ($1, $2) RETURNING *", diagnosisTable)
	err := r.db.Get(&createdDiagnosis, query,
		diagnosis.Id,
		diagnosis.Description,
	)
	return createdDiagnosis, err
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
func (r *DiagnosisRepository) UpdateDiagnosis(diagnosis model.Diagnosis) (model.Diagnosis, error) {
	var cre model.Diagnosis
	query := fmt.Sprintf("UPDATE %s SET description=$2 WHERE id=$1 RETURNING *", diagnosisTable)
	err := r.db.Get(&cre, query,
		diagnosis.Id,
		diagnosis.Description,
	)
	return cre, err
}

// Delete diagnosis data from database
func (r *DiagnosisRepository) DeleteDiagnosis(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", diagnosisTable)
	_, err := r.db.Exec(query, id)
	return err
}
