package repository

import (
	"med/pkg/config"

	"github.com/jmoiron/sqlx"
)

const (
	adminTable           = "admin"
	patientTable         = "onco_base.patient"
	doctorTable          = "doctor"
	doctorPatientTable   = "doctor_patient"
	diagnosisTable       = "diagnosis"
	drugTable            = "drug"
	unitMeasureTable     = "unit_measure"
	courseTable          = "course"
	bloodCountTable      = "blood_count"
	diseaseTable         = "disease"
	patientDiseaseTable  = "patient_disease"
	patientCourseTable   = "patient_course"
	bloodCountValueTable = "blood_count_value"
	courseProcedureTable = "course_procedure"
)

func NewPostgresDB(cfg *config.ConfigDatabase) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.GetDataSourceName())

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
