package repository

import (
	"fmt"
	"med/pkg/model"
)

// Get patient list from database
func (r *AuthorizationRepository) GetPatientList() ([]model.Patient, error) {
	var patientList []model.Patient
	query := fmt.Sprintf("SELECT * FROM %s", patientTable)
	err := r.db.Select(&patientList, query)
	return patientList, err
}

// Get patient from database by ID
func (r *AuthorizationRepository) GetPatientById(id int) (model.Patient, error) {
	var patient model.Patient
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", patientTable)
	err := r.db.Get(&patient, query, id)
	return patient, err
}

// Create patient in database and get him from database
func (r *AuthorizationRepository) CreatePatient(patient model.Patient) (model.Patient, error) {
	var createdPatient model.Patient
	query := fmt.Sprintf("INSERT INTO %s (first_name, middle_name, last_name, birth_date, sex, snils, phone) VALUES ($1, $2, $3, $4, $5, $6, $7)", patientTable)
	err := r.db.Get(&createdPatient, query,
		patient.FirstName,
		patient.MiddleName,
		patient.LastName,
		patient.BirthDate,
		patient.Sex,
		patient.SNILS,
		patient.Phone,
	)
	return createdPatient, err
}

// Update patient data in database
func (r *AuthorizationRepository) UpdatePatient(patient model.Patient) (model.Patient, error) {
	var updatedPatient model.Patient
	query := fmt.Sprintf("UPDATE %s SET first_name=$1, middle_name=$2, last_name=$3, birth_date=$4, sex=$5, snils=$6, phone=$7 WHERE id=$8", patientTable)
	err := r.db.Get(&updatedPatient, query,
		patient.FirstName,
		patient.MiddleName,
		patient.LastName,
		patient.BirthDate,
		patient.Sex,
		patient.SNILS,
		patient.Phone,
		patient.Id,
	)
	return patient, err
}

// Delete patient data from database
func (r *AuthorizationRepository) DeletePatient(patient model.Patient) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", patientTable)
	_, err := r.db.Exec(query, patient.Id)
	return err
}
