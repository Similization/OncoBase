package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type PatientRepository struct {
	db *sqlx.DB
}

func NewPatientRepository(db *sqlx.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

// Create patient in database and get him from database
func (r *PatientRepository) CreatePatient(patient model.Patient) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var patientId int
	query := fmt.Sprintf("INSERT INTO %s (first_name, middle_name, last_name, birth_date, sex, snils, phone) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *", patientTable)
	row := r.db.QueryRow(query,
		patient.FirstName,
		patient.MiddleName,
		patient.LastName,
		patient.BirthDate,
		patient.Sex,
		patient.SNILS,
		patient.Phone,
	)

	err = row.Scan(&patientId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return patientId, tx.Commit()
}

// Get patient list from database
func (r *PatientRepository) GetPatientList() ([]model.Patient, error) {
	var patientList []model.Patient
	query := fmt.Sprintf("SELECT * FROM %s", patientTable)
	err := r.db.Select(&patientList, query)
	return patientList, err
}

// Get patient from database by ID
func (r *PatientRepository) GetPatientById(id int) (model.Patient, error) {
	var patient model.Patient
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", patientTable)
	err := r.db.Get(&patient, query, id)
	return patient, err
}

// Update patient data in database
func (r *PatientRepository) UpdatePatient(patient model.Patient) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// Define the update builder
	updateBuilder := squirrel.Update(patientTable).
		Set("first_name", patient.FirstName).
		Set("middle_name", patient.MiddleName).
		Set("last_name", patient.LastName).
		Set("birth_date", patient.BirthDate).
		Set("sex", patient.Sex).
		Set("snils", patient.SNILS).
		Set("phone", patient.Phone).
		Where(squirrel.Eq{"id": patient.Id}).
		Suffix("RETURNING *")

	// Get the SQL query and arguments from the update builder
	sql, args, err := updateBuilder.ToSql()
	if err != nil {
		tx.Rollback()
		return err
	}

	// Execute the query and scan the result into updatedPatient
	_, err = r.db.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// Delete patient data from database
func (r *PatientRepository) DeletePatient(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", patientTable)
	_, err = r.db.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
