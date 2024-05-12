package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type PatientDiseaseRepository struct {
	db *sqlx.DB
}

func NewPatientDiseaseRepository(db *sqlx.DB) *PatientDiseaseRepository {
	return &PatientDiseaseRepository{db: db}
}

// Create patient in database and get him from database
func (r *PatientDiseaseRepository) CreatePatientDisease(patientDisease model.PatientDisease) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (stage, diagnosis, patient, disease) VALUES ($1, $2, $3, $4)", patientDiseaseTable)
	_, err = r.db.Exec(query,
		patientDisease.Stage,
		patientDisease.Diagnosis,
		patientDisease.Patient,
		patientDisease.Disease,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Get patient list from database
func (r *PatientDiseaseRepository) GetPatientDiseaseList() ([]model.PatientDisease, error) {
	var patientDiseaseList []model.PatientDisease
	query := fmt.Sprintf("SELECT * FROM %s", patientDiseaseTable)
	err := r.db.Select(&patientDiseaseList, query)
	return patientDiseaseList, err
}

// Get patient list from database
func (r *PatientDiseaseRepository) GetPatientDiseaseListByPatient(patientId int) ([]model.PatientDisease, error) {
	var patientDiseaseList []model.PatientDisease
	query := fmt.Sprintf("SELECT * FROM %s WHERE patient=$1", patientDiseaseTable)
	err := r.db.Select(&patientDiseaseList, query, patientId)
	return patientDiseaseList, err
}

// Get patient list from database
func (r *PatientDiseaseRepository) GetPatientDiseaseListByDisease(diseaseId string) ([]model.PatientDisease, error) {
	var patientDiseaseList []model.PatientDisease
	query := fmt.Sprintf("SELECT * FROM %s WHERE disease=$1", patientDiseaseTable)
	err := r.db.Select(&patientDiseaseList, query, diseaseId)
	return patientDiseaseList, err
}

// Get patient from database by ID
func (r *PatientDiseaseRepository) GetPatientDiseaseById(patientId int, diseaseId string) (model.PatientDisease, error) {
	var patientDisease model.PatientDisease
	query := fmt.Sprintf("SELECT * FROM %s WHERE patient=$1 AND disease=$2", patientDiseaseTable)
	err := r.db.Get(&patientDisease, query, patientId, diseaseId)
	return patientDisease, err
}

// Update patient data in database
func (r *PatientDiseaseRepository) UpdatePatientDisease(patientDisease model.PatientDisease) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET stage=$1, diagnosis=$2 WHERE patient=$3 AND disease=$4", patientDiseaseTable)
	_, err = r.db.Exec(query,
		patientDisease.Stage,
		patientDisease.Diagnosis,
		patientDisease.Patient,
		patientDisease.Disease,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Delete patient data from database
func (r *PatientDiseaseRepository) DeletePatientDisease(patientId int, diseaseId string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE patient=$1 AND disease=$2", patientDiseaseTable)
	_, err = r.db.Exec(query, patientId, diseaseId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
