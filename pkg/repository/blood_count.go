package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type BloodCountRepository struct {
	db *sqlx.DB
}

func NewBloodCountRepository(db *sqlx.DB) *BloodCountRepository {
	return &BloodCountRepository{db: db}
}

func (r *BloodCountRepository) CreateBloodCount(bloodCount model.BloodCount) (model.BloodCount, error) {
	var createdBloodCount model.BloodCount
	query := fmt.Sprintf("INSERT INTO %s (id, description, min_normal_value, max_normal_value, min_possible_value, max_possible_value, measure_code) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *", bloodCountTable)
	err := r.db.Get(&createdBloodCount, query,
		bloodCount.Id,
		bloodCount.Description,
		bloodCount.MinNormalValue,
		bloodCount.MaxNormalValue,
		bloodCount.MinPossibleValue,
		bloodCount.MaxPossibleValue,
		bloodCount.MeasureCode,
	)
	return createdBloodCount, err
}

func (r *BloodCountRepository) GetBloodCountList() ([]model.BloodCount, error) {
	var bloodCountList []model.BloodCount
	query := fmt.Sprintf("SELECT * FROM %s", bloodCountTable)
	err := r.db.Select(&bloodCountList, query)
	return bloodCountList, err
}

func (r *BloodCountRepository) GetBloodCountById(id string) (model.BloodCount, error) {
	var bloodCount model.BloodCount
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", bloodCountTable)
	err := r.db.Get(&bloodCount, query, id)
	return bloodCount, err
}

func (r *BloodCountRepository) UpdateBloodCount(bloodCount model.BloodCount) (model.BloodCount, error) {
	var updatedBloodCount model.BloodCount
	query := fmt.Sprintf("UPDATE %s SET description=$1, min_normal_value=$2, max_normal_value=$3, min_possible_value=$4, max_possible_value=$5, measure_code=$6 WHERE id=$7 RETURNING *", bloodCountTable)
	err := r.db.Get(&updatedBloodCount, query,
		bloodCount.Description,
		bloodCount.MinNormalValue,
		bloodCount.MaxNormalValue,
		bloodCount.MinPossibleValue,
		bloodCount.MaxPossibleValue,
		bloodCount.MeasureCode,
		bloodCount.Id,
	)
	return bloodCount, err
}

func (r *BloodCountRepository) DeleteBloodCount(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", bloodCountTable)
	_, err := r.db.Exec(query, id)
	return err
}
