package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type BloodCountValueRepository struct {
	db *sqlx.DB
}

func NewBloodCountValueRepository(db *sqlx.DB) *BloodCountValueRepository {
	return &BloodCountValueRepository{db: db}
}

func (r *BloodCountValueRepository) CreateBloodCountValue(bloodCountValue model.BloodCountValue) (model.BloodCountValue, error) {
	var createdBloodCountValue model.BloodCountValue
	query := fmt.Sprintf("INSERT INTO %s (coefficient, description, disease, blood_count) VALUES ($1, $2, $3, $4)", bloodCountTable)
	err := r.db.Get(&createdBloodCountValue, query,
		bloodCountValue.Coefficient,
		bloodCountValue.Description,
		bloodCountValue.Disease,
		bloodCountValue.BloodCount,
	)
	return createdBloodCountValue, err
}

func (r *BloodCountValueRepository) GetBloodCountValueList() ([]model.BloodCountValue, error) {
	var bloodCountList []model.BloodCountValue
	query := fmt.Sprintf("SELECT FROM %s", bloodCountTable)
	err := r.db.Select(&bloodCountList, query)
	return bloodCountList, err
}

func (r *BloodCountValueRepository) GetBloodCountValueListByDisease(diseaseId string) ([]model.BloodCountValue, error) {
	var bloodCountList []model.BloodCountValue
	query := fmt.Sprintf("SELECT FROM %s", bloodCountTable)
	err := r.db.Select(&bloodCountList, query)
	return bloodCountList, err
}

func (r *BloodCountValueRepository) GetBloodCountValueListByBloodCount(bloodCountId string) ([]model.BloodCountValue, error) {
	var bloodCountList []model.BloodCountValue
	query := fmt.Sprintf("SELECT FROM %s", bloodCountTable)
	err := r.db.Select(&bloodCountList, query)
	return bloodCountList, err
}

func (r *BloodCountValueRepository) GetBloodCountValueById(diseaseId, bloodCountId string) (model.BloodCountValue, error) {
	var bloodCountValue model.BloodCountValue
	query := fmt.Sprintf("SELECT FROM %s WHERE disease=$1 AND blood_count=$2", bloodCountTable)
	err := r.db.Get(&bloodCountValue, query, diseaseId, bloodCountId)
	return bloodCountValue, err
}

func (r *BloodCountValueRepository) UpdateBloodCountValue(bloodCountValue model.BloodCountValue) (model.BloodCountValue, error) {
	var updatedBloodCountValue model.BloodCountValue
	query := fmt.Sprintf("UPDATE %s SET coefficient=$1, description=$2 WHERE disease=$3 AND blood_count=$4", bloodCountTable)
	err := r.db.Get(&updatedBloodCountValue, query,
		bloodCountValue.Coefficient,
		bloodCountValue.Description,
		bloodCountValue.Disease,
		bloodCountValue.BloodCount,
	)
	return bloodCountValue, err
}

func (r *BloodCountValueRepository) DeleteBloodCountValue(diseaseId, bloodCountId string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE disease=$1 AND blood_count=$2", bloodCountTable)
	_, err := r.db.Exec(query, diseaseId, bloodCountId)
	return err
}
