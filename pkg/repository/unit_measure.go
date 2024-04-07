package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type UnitMeasureRepository struct {
	db *sqlx.DB
}

func NewUnitMeasureRepository(db *sqlx.DB) *UnitMeasureRepository {
	return &UnitMeasureRepository{db: db}
}

// Create unit measure in database and get him from database
func (r *UnitMeasureRepository) CreateUnitMeasure(unitMeasure model.UnitMeasure) (model.UnitMeasure, error) {
	var createdUnitMeasure model.UnitMeasure
	query := fmt.Sprintf("INSERT INTO %s (id, shorthand, full_text, global) VALUES ($1, $2, $3, $4) RETURNING *", unitMeasureTable)
	err := r.db.Get(&createdUnitMeasure, query,
		unitMeasure.Id,
		unitMeasure.Shorthand,
		unitMeasure.FullText,
		unitMeasure.Global,
	)
	return createdUnitMeasure, err
}

// Get unit measure list from database
func (r *UnitMeasureRepository) GetUnitMeasureList() ([]model.UnitMeasure, error) {
	var unitMeasureList []model.UnitMeasure
	query := fmt.Sprintf("SELECT * FROM %s", unitMeasureTable)
	err := r.db.Select(&unitMeasureList, query)
	return unitMeasureList, err
}

// Get unit measure from database by ID
func (r *UnitMeasureRepository) GetUnitMeasureById(id string) (model.UnitMeasure, error) {
	var unitMeasure model.UnitMeasure
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", unitMeasureTable)
	err := r.db.Get(&unitMeasure, query, id)
	return unitMeasure, err
}

// Update unit measure data in database
func (r *UnitMeasureRepository) UpdateUnitMeasure(unitMeasure model.UnitMeasure) (model.UnitMeasure, error) {
	var updatedUnitMeasure model.UnitMeasure
	query := fmt.Sprintf("UPDATE %s SET shorthand=$2, full_text=$3, global=$4 WHERE id=$1 RETURNING *", unitMeasureTable)
	err := r.db.Get(&updatedUnitMeasure, query,
		unitMeasure.Id,
		unitMeasure.Shorthand,
		unitMeasure.FullText,
		unitMeasure.Global,
	)
	return updatedUnitMeasure, err
}

// Delete unit measure data from database
func (r *UnitMeasureRepository) DeleteUnitMeasure(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", unitMeasureTable)
	_, err := r.db.Exec(query, id)
	return err
}
