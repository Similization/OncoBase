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
func (r *UnitMeasureRepository) CreateUnitMeasure(unitMeasure model.UnitMeasure) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (id, shorthand, full_text, global) VALUES ($1, $2, $3, $4)", unitMeasureTable)
	_, err = r.db.Exec(query,
		unitMeasure.Id,
		unitMeasure.Shorthand,
		unitMeasure.FullText,
		unitMeasure.Global,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
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
func (r *UnitMeasureRepository) UpdateUnitMeasure(unitMeasure model.UnitMeasure) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET shorthand=$2, full_text=$3, global=$4 WHERE id=$1 RETURNING *", unitMeasureTable)
	_, err = r.db.Exec(query,
		unitMeasure.Id,
		unitMeasure.Shorthand,
		unitMeasure.FullText,
		unitMeasure.Global,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Delete unit measure data from database
func (r *UnitMeasureRepository) DeleteUnitMeasure(id string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", unitMeasureTable)
	_, err = r.db.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
