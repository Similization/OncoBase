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
func (r *CourseProcedureRepository) CreateCourseProcedure(courseProcedure model.CourseProcedure) (model.CourseProcedure, error) {
	var createdCourseProcedure model.CourseProcedure
	query := fmt.Sprintf("INSERT INTO %s (patient_course, doctor, begin_date, period, result) VALUES ($1, $2, $3, $4, $5) RETURNING *", courseProcedureTable)
	err := r.db.Get(&createdCourseProcedure, query,
		courseProcedure.PatientCourse,
		courseProcedure.Doctor,
		courseProcedure.BeginDate,
		courseProcedure.Period,
		courseProcedure.Result,
	)
	return createdCourseProcedure, err
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
func (r *CourseProcedureRepository) UpdateCourseProcedure(courseProcedure model.CourseProcedure) (model.CourseProcedure, error) {
	var updatedCourseProcedure model.CourseProcedure
	query := fmt.Sprintf("UPDATE %s SET patient_course=$1, doctor=$2, begin_date=$3, period=$4, result=$5 WHERE id=$6 RETURNING *", courseProcedureTable)
	err := r.db.Get(&updatedCourseProcedure, query,
		courseProcedure.PatientCourse,
		courseProcedure.Doctor,
		courseProcedure.BeginDate,
		courseProcedure.Period,
		courseProcedure.Result,
		courseProcedure.Id,
	)
	return courseProcedure, err
}

// Delete course procedure from database by id
func (r *CourseProcedureRepository) DeleteCourseProcedure(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", courseProcedureTable)
	_, err := r.db.Exec(query, id)
	return err
}
