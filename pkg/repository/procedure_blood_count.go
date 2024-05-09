package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type ProcedureBloodCountRepository struct {
	db *sqlx.DB
}

func NewProcedureBloodCountRepository(db *sqlx.DB) *ProcedureBloodCountRepository {
	return &ProcedureBloodCountRepository{db: db}
}

// Create patient in database and get him from database
func (r *ProcedureBloodCountRepository) CreateProcedureBloodCount(procedureBloodCount model.ProcedureBloodCount) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (value, measure_code, procedure, blood_count) VALUES ($1, $2, $3, $4)", procedureBloodCountTable)
	_, err = r.db.Exec(query,
		procedureBloodCount.Value,
		procedureBloodCount.MeasureCode,
		procedureBloodCount.Procedure,
		procedureBloodCount.BloodCount,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Get patient list from database
func (r *ProcedureBloodCountRepository) GetProcedureBloodCountList() ([]model.ProcedureBloodCount, error) {
	var procedureBloodCountList []model.ProcedureBloodCount
	query := fmt.Sprintf("SELECT * FROM %s", procedureBloodCountTable)
	err := r.db.Select(&procedureBloodCountList, query)
	return procedureBloodCountList, err
}

// Get patient list from database
func (r *ProcedureBloodCountRepository) GetProcedureBloodCountListByProcedure(procedureId int) ([]model.ProcedureBloodCount, error) {
	var procedureBloodCountList []model.ProcedureBloodCount
	query := fmt.Sprintf("SELECT * FROM %s WHERE procedure=$1", procedureBloodCountTable)
	err := r.db.Select(&procedureBloodCountList, query, procedureId)
	return procedureBloodCountList, err
}

// Get patient list from database
func (r *ProcedureBloodCountRepository) GetProcedureBloodCountListByBloodCount(bloodCountId string) ([]model.ProcedureBloodCount, error) {
	var procedureBloodCountList []model.ProcedureBloodCount
	query := fmt.Sprintf("SELECT * FROM %s WHERE bloodCount=$1", procedureBloodCountTable)
	err := r.db.Select(&procedureBloodCountList, query, bloodCountId)
	return procedureBloodCountList, err
}

// Get patient from database by ID
func (r *ProcedureBloodCountRepository) GetProcedureBloodCountById(procedureId int, bloodCountId string) (model.ProcedureBloodCount, error) {
	var procedureBloodCount model.ProcedureBloodCount
	query := fmt.Sprintf("SELECT * FROM %s WHERE procedure=$1 AND blood_count=$2", procedureBloodCountTable)
	err := r.db.Get(&procedureBloodCount, query, procedureId, bloodCountId)
	return procedureBloodCount, err
}

// Update patient data in database
func (r *ProcedureBloodCountRepository) UpdateProcedureBloodCount(procedureBloodCount model.ProcedureBloodCount) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("UPDATE %s SET value=$1, measure_code=$2 WHERE procedure=$3 AND blood_count=$4", procedureBloodCountTable)
	_, err = r.db.Exec(query,
		procedureBloodCount.Value,
		procedureBloodCount.MeasureCode,
		procedureBloodCount.Procedure,
		procedureBloodCount.BloodCount,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Delete patient data from database
func (r *ProcedureBloodCountRepository) DeleteProcedureBloodCount(procedureId int, bloodCountId string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE procedure=$1 AND blood_count=$2", procedureBloodCountTable)
	_, err = r.db.Exec(query, procedureId, bloodCountId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
