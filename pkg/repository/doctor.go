package repository

import (
	"fmt"
	"med/pkg/model"
)

func (r *AuthorizationRepository) GetDoctorList() ([]model.Doctor, error) {
	var doctorList []model.Doctor
	query := fmt.Sprintf("SELECT FROM %s", doctorTable)
	err := r.db.Select(&doctorList, query)
	return doctorList, err
}

func (r *AuthorizationRepository) GetDoctorById(id int) (model.Doctor, error) {
	var doctor model.Doctor
	query := fmt.Sprintf("SELECT FROM %s WHERE id=$1", doctorTable)
	err := r.db.Get(&doctor, query, id)
	return doctor, err
}

func (r *AuthorizationRepository) CreateDoctor(patient model.Patient) (model.Doctor, error) {
	var createdDoctor model.Doctor
	query := fmt.Sprintf("INSERT INTO %s (first_name, middle_name, last_name, birth_date, sex, phone) VALUES ($1, $2, $3, $4, $5, $6)", patientTable)
	err := r.db.Get(&createdDoctor, query,
		patient.FirstName,
		patient.MiddleName,
		patient.LastName,
		patient.BirthDate,
		patient.Sex,
		patient.Phone,
	)
	return createdDoctor, err
}

func (r *AuthorizationRepository) UpdateDoctor(doctor model.Doctor) (model.Doctor, error) {
	var updatedDoctor model.Doctor
	query := fmt.Sprintf("UPDATE %s SET first_name=$1, middle_name=$2, last_name=$3, birth_date=$4, sex=$5, phone=$6 WHERE id=$7", doctorTable)
	err := r.db.Get(&updatedDoctor, query,
		doctor.FirstName,
		doctor.MiddleName,
		doctor.LastName,
		doctor.BirthDate,
		doctor.Sex,
		doctor.Phone,
		doctor.Id,
	)
	return doctor, err
}

func (r *AuthorizationRepository) DeleteDoctor(doctor model.Doctor) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", doctorTable)
	_, err := r.db.Exec(query, doctor.Id)
	return err
}
