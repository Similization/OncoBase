package repository

import (
	"med/pkg/config"

	"github.com/jmoiron/sqlx"
)

const (
	externalUserTable  = "onco_base.external_user"
	internalUserTable  = "onco_base.internal_user"
	patientTable       = "onco_base.patient"
	doctorTable        = "onco_base.doctor"
	doctorPatientTable = "onco_base.doctor_patient"
	diagnosisTable     = "onco_base.diagnosis"
	diseaseTable       = "onco_base.disease"
	patientCourseTable = "onco_base.patient_course"
	unitMeasureTable   = "onco_base.unit_measure"
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
