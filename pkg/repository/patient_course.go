package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type PatientCourseRepository struct {
	db *sqlx.DB
}

func NewPatientCourseRepository(db *sqlx.DB) *PatientCourseRepository {
	return &PatientCourseRepository{db: db}
}

// Create patient course in database and get him from database
func (r *PatientCourseRepository) CreatePatientCourse(patientCourse model.PatientCourse) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var patientCourseId int
	query := fmt.Sprintf("INSERT INTO %s (patient, disease, course, doctor, begin_date, end_date, diagnosis) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", patientCourseTable)
	row := r.db.QueryRow(query,
		patientCourse.Patient,
		patientCourse.Disease,
		patientCourse.Course,
		patientCourse.Doctor,
		patientCourse.BeginDate,
		patientCourse.EndDate,
		patientCourse.Diagnosis,
	)

	err = row.Scan(&patientCourseId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return patientCourseId, tx.Commit()
}

// Get patient course list from database
func (r *PatientCourseRepository) GetPatientCourseList() ([]model.PatientCourse, error) {
	var patientCourseList []model.PatientCourse
	query := fmt.Sprintf("SELECT * FROM %s", patientCourseTable)
	err := r.db.Select(&patientCourseList, query)
	return patientCourseList, err
}

// Get patient course from database by ID
func (r *PatientCourseRepository) GetPatientCourseById(patientCourseId int) (model.PatientCourse, error) {
	var patientCourse model.PatientCourse
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", patientCourseTable)
	err := r.db.Get(&patientCourse, query, patientCourseId)
	return patientCourse, err
}

// Update patient course data in database
func (r *PatientCourseRepository) UpdatePatientCourse(patientCourse model.PatientCourse) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET patient=$1, disease=$2, course=$3, doctor=$4, begin_date=$5, end_date=$6, diagnosis=$7 WHERE id=$8", patientCourseTable)
	_, err = r.db.Exec(query,
		patientCourse.Patient,
		patientCourse.Disease,
		patientCourse.Course,
		patientCourse.Doctor,
		patientCourse.BeginDate,
		patientCourse.EndDate,
		patientCourse.Diagnosis,
		patientCourse.Id,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Delete patient course data from database
func (r *PatientCourseRepository) DeletePatientCourse(patientCourseId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", patientCourseTable)
	_, err = r.db.Exec(query, patientCourseId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
