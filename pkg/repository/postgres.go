package repository

import (
	"med/pkg/config"

	"github.com/jmoiron/sqlx"
)

const (
	externalUserTable = "onco_base.external_user"
	internalUserTable = "onco_base.internal_user"

	bloodCountTable          = "onco_base.blood_count"
	bloodCountValueTable     = "onco_base.blood_count_value"
	courseTable              = "onco_base.course"
	courseProcedureTable     = "onco_base.course_procedure"
	diagnosisTable           = "onco_base.diagnosis"
	diseaseTable             = "onco_base.disease"
	doctorTable              = "onco_base.doctor"
	doctorPatientTable       = "onco_base.doctor_patient"
	drugTable                = "onco_base.drug"
	patientTable             = "onco_base.patient"
	patientCourseTable       = "onco_base.patient_course"
	patientDiseaseTable      = "onco_base.patient_disease"
	procedureBloodCountTable = "onco_base.procedure_blood_count"
	unitMeasureTable         = "onco_base.unit_measure"
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
