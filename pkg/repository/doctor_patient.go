package repository

import (
	"fmt"
	"med/pkg/model"
)

func (r *AuthorizationRepository) GetDoctorPatientList(doctor_id int) ([]model.Patient, error) {
	var doctorPatients []model.Patient
	query := fmt.Sprintf("SELECT (id, first_name, middle_name, last_name, birth_date, sex, snils) FROM %s JOIN onco_base.patient p on p.id = doctor_patient.patient WHERE doctor = $1", doctorPatientTable)
	err := r.db.Select(&doctorPatients, query, doctor_id)
	return doctorPatients, err
}

func (r *AuthorizationRepository) GetDoctorPatientBy() {
}

func (r *AuthorizationRepository) CreateDoctorPatient(doctor_id, patient_id int) error {
	query := fmt.Sprintf("INSERT INTO %s SET patient=$1, doctor=$2", doctorPatientTable)
	_, err := r.db.Exec(query, patient_id, doctor_id)
	return err
}

func (r *AuthorizationRepository) DeleteDoctorPatient(doctor_id, patient_id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE patient=$1 AND doctor=$2", doctorPatientTable)
	_, err := r.db.Exec(query, patient_id, doctor_id)
	return err
}
