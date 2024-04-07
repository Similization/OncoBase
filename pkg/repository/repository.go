package repository

import (
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (string, error)
	GetUser(email, password string) (model.User, error)
}

type Account interface {
}

type BloodCount interface {
}

type UnitMeasure interface {
	CreateUnitMeasure(unitMeasure model.UnitMeasure) (model.UnitMeasure, error)
	GetUnitMeasureById(id string) (model.UnitMeasure, error)
	GetUnitMeasureList() ([]model.UnitMeasure, error)
	UpdateUnitMeasure(unitMeasure model.UnitMeasure) (model.UnitMeasure, error)
	DeleteUnitMeasure(id string) error
}

type Disease interface {
	CreateDisease(disease model.Disease) (model.Disease, error)
	GetDiseaseById(id string) (model.Disease, error)
	GetDiseaseList() ([]model.Disease, error)
	UpdateDisease(disease model.Disease) (model.Disease, error)
	DeleteDisease(id string) error
}

type Diagnosis interface {
	CreateDiagnosis(diagnosis model.Diagnosis) (model.Diagnosis, error)
	GetDiagnosisById(id string) (model.Diagnosis, error)
	GetDiagnosisList() ([]model.Diagnosis, error)
	UpdateDiagnosis(diagnosis model.Diagnosis) (model.Diagnosis, error)
	DeleteDiagnosis(id string) error
}

type Drug interface {
	CreateDrug(drug model.Drug) (model.Drug, error)
	GetDrugById(id string) (model.Drug, error)
	GetDrugList() ([]model.Drug, error)
	UpdateDrug(drug model.Drug) (model.Drug, error)
	DeleteDrug(id string) error
}

type Patient interface {
	CreatePatient(patient model.Patient) (model.Patient, error)
	GetPatientById(id int) (model.Patient, error)
	GetPatientList() ([]model.Patient, error)
	UpdatePatient(patient model.Patient) (model.Patient, error)
	DeletePatient(id int) error
}

type Repository struct {
	Authorization
	Account
	BloodCount
	UnitMeasure
	Diagnosis
	Disease
	Drug
	Patient
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		UnitMeasure:   NewUnitMeasureRepository(db),
		Diagnosis:     NewDiagnosisRepository(db),
		Disease:       NewDiseaseRepository(db),
		Drug:          NewDrugRepository(db),
		Patient:       NewPatientRepository(db),
	}
}
