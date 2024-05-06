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
	query := fmt.Sprintf("INSERT INTO %s (disease, blood_count, coefficient, description) VALUES ($1, $2, $3, $4) RETURNING *", bloodCountValueTable)
	err := r.db.Get(&createdBloodCountValue, query,
		bloodCountValue.Disease,
		bloodCountValue.BloodCount,
		bloodCountValue.Coefficient,
		bloodCountValue.Description,
	)
	return createdBloodCountValue, err
}

func (r *BloodCountValueRepository) GetBloodCountValueList() ([]model.BloodCountValue, error) {
	var bloodCountList []model.BloodCountValue
	query := fmt.Sprintf("SELECT * FROM %s", bloodCountValueTable)
	err := r.db.Select(&bloodCountList, query)
	return bloodCountList, err
}

func (r *BloodCountValueRepository) GetBloodCountValueListByDisease(diseaseId string) ([]model.BloodCountValue, error) {
	var bloodCountList []model.BloodCountValue
	query := fmt.Sprintf("SELECT * FROM %s WHERE disease=$1", bloodCountValueTable)
	err := r.db.Select(&bloodCountList, query, diseaseId)
	return bloodCountList, err
}

func (r *BloodCountValueRepository) GetBloodCountValueListByBloodCount(bloodCountId string) ([]model.BloodCountValue, error) {
	var bloodCountList []model.BloodCountValue
	query := fmt.Sprintf("SELECT * FROM %s WHERE blood_count=$1", bloodCountValueTable)
	err := r.db.Select(&bloodCountList, query, bloodCountId)
	return bloodCountList, err
}

func (r *BloodCountValueRepository) GetBloodCountValueById(diseaseId, bloodCountId string) (model.BloodCountValue, error) {
	var bloodCountValue model.BloodCountValue
	query := fmt.Sprintf("SELECT * FROM %s WHERE disease=$1 AND blood_count=$2", bloodCountValueTable)
	err := r.db.Get(&bloodCountValue, query, diseaseId, bloodCountId)
	return bloodCountValue, err
}

func (r *BloodCountValueRepository) UpdateBloodCountValue(bloodCountValue model.BloodCountValue) (model.BloodCountValue, error) {
	var updatedBloodCountValue model.BloodCountValue
	query := fmt.Sprintf("UPDATE %s SET coefficient=$1, description=$2 WHERE disease=$3 AND blood_count=$4 RETURNING *", bloodCountValueTable)
	err := r.db.Get(&updatedBloodCountValue, query,
		bloodCountValue.Coefficient,
		bloodCountValue.Description,
		bloodCountValue.Disease,
		bloodCountValue.BloodCount,
	)
	return bloodCountValue, err
}

func (r *BloodCountValueRepository) DeleteBloodCountValue(diseaseId, bloodCountId string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE disease=$1 AND blood_count=$2", bloodCountValueTable)
	_, err := r.db.Exec(query, diseaseId, bloodCountId)
	return err
}
