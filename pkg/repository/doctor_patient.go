package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type DoctorPatientRepository struct {
	db *sqlx.DB
}

func NewDoctorPatientRepository(db *sqlx.DB) *DoctorPatientRepository {
	return &DoctorPatientRepository{db: db}
}

// Create doctor patient in database and get him from database
func (r *DoctorPatientRepository) CreateDoctorPatient(doctorPatient model.DoctorPatient) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (patient, doctor) VALUES ($1, $2)", doctorPatientTable)
	_, err = r.db.Exec(query,
		doctorPatient.Patient,
		doctorPatient.Doctor,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Get doctor patient list from database
func (r *DoctorPatientRepository) GetDoctorPatientListByDoctor(doctor_id int) ([]model.DoctorPatient, error) {
	var doctorPatientList []model.DoctorPatient
	query := fmt.Sprintf("SELECT * FROM %s WHERE doctor=$1", doctorPatientTable)
	err := r.db.Select(&doctorPatientList, query, doctor_id)
	return doctorPatientList, err
}

// Get doctor patient list from database
func (r *DoctorPatientRepository) GetDoctorPatientListByPatient(patient_id int) ([]model.DoctorPatient, error) {
	var doctorPatientList []model.DoctorPatient
	query := fmt.Sprintf("SELECT * FROM %s WHERE patient=$1", doctorPatientTable)
	err := r.db.Select(&doctorPatientList, query, patient_id)
	return doctorPatientList, err
}

// Delete doctor patient data from database
func (r *DoctorPatientRepository) DeleteDoctorPatient(doctorPatient model.DoctorPatient) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE patient=$1 AND doctor=$2", doctorPatientTable)
	_, err = r.db.Exec(query,
		doctorPatient.Patient,
		doctorPatient.Doctor,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
