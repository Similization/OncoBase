package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type DiseaseRepository struct {
	db *sqlx.DB
}

func NewDiseaseRepository(db *sqlx.DB) *DiseaseRepository {
	return &DiseaseRepository{db: db}
}

// Create disease in database and get it from database
func (r *DiseaseRepository) CreateDisease(disease model.Disease) (model.Disease, error) {
	var createdDisease model.Disease
	query := fmt.Sprintf("INSERT INTO %s (id, description) VALUES ($1, $2) RETURNING *", diseaseTable)
	err := r.db.Get(&createdDisease, query,
		disease.Id,
		disease.Description,
	)
	return createdDisease, err
}

// Get disease list from database
func (r *DiseaseRepository) GetDiseaseList() ([]model.Disease, error) {
	var diseaseList []model.Disease
	query := fmt.Sprintf("SELECT * FROM %s", diseaseTable)
	err := r.db.Select(&diseaseList, query)
	return diseaseList, err
}

// Get disease from database by ID
func (r *DiseaseRepository) GetDiseaseById(id string) (model.Disease, error) {
	var disease model.Disease
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", diseaseTable)
	err := r.db.Get(&disease, query, id)
	return disease, err
}

// Update disease data in database
func (r *DiseaseRepository) UpdateDisease(disease model.Disease) (model.Disease, error) {
	var updatedDisease model.Disease
	query := fmt.Sprintf("UPDATE %s SET description=$2 WHERE id=$1 RETURNING *", diseaseTable)
	err := r.db.Get(&updatedDisease, query,
		disease.Id,
		disease.Description,
	)
	return updatedDisease, err
}

// Delete disease data from database
func (r *DiseaseRepository) DeleteDisease(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", diseaseTable)
	_, err := r.db.Exec(query, id)
	return err
}
