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
	CreateDisease(unitMeasure model.Disease) (model.Disease, error)
	GetDiseaseById(id string) (model.Disease, error)
	GetDiseaseList() ([]model.Disease, error)
	UpdateDisease(unitMeasure model.Disease) (model.Disease, error)
	DeleteDisease(id string) error
}

type Diagnosis interface {
	CreateDiagnosis(unitMeasure model.Diagnosis) (model.Diagnosis, error)
	GetDiagnosisById(id string) (model.Diagnosis, error)
	GetDiagnosisList() ([]model.Diagnosis, error)
	UpdateDiagnosis(unitMeasure model.Diagnosis) (model.Diagnosis, error)
	DeleteDiagnosis(id string) error
}

type Repository struct {
	Authorization
	Account
	BloodCount
	UnitMeasure
	Diagnosis
	Disease
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		UnitMeasure:   NewUnitMeasureRepository(db),
		Diagnosis:     NewDiagnosisRepository(db),
		Disease:       NewDiseaseRepository(db),
	}
}
