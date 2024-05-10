package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type CourseProcedureRepository struct {
	db *sqlx.DB
}

func NewCourseProcedureRepository(db *sqlx.DB) *CourseProcedureRepository {
	return &CourseProcedureRepository{db: db}
}

// Create course procedure in database and get it from database
func (r *CourseProcedureRepository) CreateCourseProcedure(courseProcedure model.CourseProcedure) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var courseProcedureId int
	query := fmt.Sprintf("INSERT INTO %s (patient_course, doctor, begin_date, period, result) VALUES ($1, $2, $3, $4, $5) RETURNING id", courseProcedureTable)
	row := r.db.QueryRow(query,
		courseProcedure.PatientCourse,
		courseProcedure.Doctor,
		courseProcedure.BeginDate,
		courseProcedure.Period,
		courseProcedure.Result,
	)

	err = row.Scan(&courseProcedure)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return courseProcedureId, tx.Commit()
}

// Get course procedure list from database
func (r *CourseProcedureRepository) GetCourseProcedureList() ([]model.CourseProcedure, error) {
	var courseProcedureList []model.CourseProcedure
	query := fmt.Sprintf("SELECT * FROM %s", courseProcedureTable)
	err := r.db.Select(&courseProcedureList, query)
	return courseProcedureList, err
}

// Get course procedure from database by id
func (r *CourseProcedureRepository) GetCourseProcedureById(id string) (model.CourseProcedure, error) {
	var courseProcedure model.CourseProcedure
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", courseProcedureTable)
	err := r.db.Get(&courseProcedure, query, id)
	return courseProcedure, err
}

// Update course procedure fields in database and get it from database
func (r *CourseProcedureRepository) UpdateCourseProcedure(courseProcedure model.CourseProcedure) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET patient_course=$1, doctor=$2, begin_date=$3, period=$4, result=$5 WHERE id=$6", courseProcedureTable)
	_, err = r.db.Exec(query,
		courseProcedure.PatientCourse,
		courseProcedure.Doctor,
		courseProcedure.BeginDate,
		courseProcedure.Period,
		courseProcedure.Result,
		courseProcedure.Id,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Delete course procedure from database by id
func (r *CourseProcedureRepository) DeleteCourseProcedure(id string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", courseProcedureTable)
	_, err = r.db.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
