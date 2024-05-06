package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type DoctorRepository struct {
	db *sqlx.DB
}

func NewDoctorRepository(db *sqlx.DB) *DoctorRepository {
	return &DoctorRepository{db: db}
}

// Create doctor in database and get him from database
func (r *DoctorRepository) CreateDoctor(doctor model.Doctor) (model.Doctor, error) {
	var createdDoctor model.Doctor
	query := fmt.Sprintf("INSERT INTO %s (first_name, middle_name, last_name, qualification, phone, user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *", doctorTable)
	err := r.db.Get(&createdDoctor, query,
		doctor.FirstName,
		doctor.MiddleName,
		doctor.LastName,
		doctor.Qualification,
		doctor.Phone,
		doctor.UserId,
	)
	return createdDoctor, err
}

// Get doctor list from database
func (r *DoctorRepository) GetDoctorList() ([]model.Doctor, error) {
	var doctorList []model.Doctor
	query := fmt.Sprintf("SELECT * FROM %s", doctorTable)
	err := r.db.Select(&doctorList, query)
	return doctorList, err
}

// Get doctor from database by ID
func (r *DoctorRepository) GetDoctorById(id int) (model.Doctor, error) {
	var doctor model.Doctor
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", doctorTable)
	err := r.db.Get(&doctor, query, id)
	return doctor, err
}

// Update doctor data in database
func (r *DoctorRepository) UpdateDoctor(doctor model.Doctor) (model.Doctor, error) {
	var updatedDoctor model.Doctor
	query := fmt.Sprintf("UPDATE %s SET first_name=$1, middle_name=$2, last_name=$3, qualification=$4, phone=$5, user_id=$6 WHERE id=$7 RETURNING *", doctorTable)
	err := r.db.Get(&updatedDoctor, query,
		doctor.FirstName,
		doctor.MiddleName,
		doctor.LastName,
		doctor.Qualification,
		doctor.Phone,
		doctor.UserId,
		doctor.Id,
	)
	return updatedDoctor, err
}

// Delete doctor data from database
func (r *DoctorRepository) DeleteDoctor(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", doctorTable)
	_, err := r.db.Exec(query, id)
	return err
}
