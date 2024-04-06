package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (string, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (*UserData, error)
}

type Account interface {
}

type UnitMeasure interface {
	CreateUnitMeasure(unitMeasure model.UnitMeasure) (model.UnitMeasure, error)
	GetUnitMeasureById(id string) (model.UnitMeasure, error)
	GetUnitMeasureList() ([]model.UnitMeasure, error)
	UpdateUnitMeasure(unitMeasure model.UnitMeasure) (model.UnitMeasure, error)
	DeleteUnitMeasure(id string) error
}

type Diagnosis interface {
	CreateDiagnosis(unitMeasure model.Diagnosis) (model.Diagnosis, error)
	GetDiagnosisById(id string) (model.Diagnosis, error)
	GetDiagnosisList() ([]model.Diagnosis, error)
	UpdateDiagnosis(unitMeasure model.Diagnosis) (model.Diagnosis, error)
	DeleteDiagnosis(id string) error
}

type Disease interface {
	CreateDisease(unitMeasure model.Disease) (model.Disease, error)
	GetDiseaseById(id string) (model.Disease, error)
	GetDiseaseList() ([]model.Disease, error)
	UpdateDisease(unitMeasure model.Disease) (model.Disease, error)
	DeleteDisease(id string) error
}

type BloodCount interface {
}

type Service struct {
	Authorization
	Account
	BloodCount
	UnitMeasure
	Disease
	Diagnosis
}

func NewService(repos repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		UnitMeasure:   NewUnitMeasureService(repos.UnitMeasure),
		Disease:       NewDiseaseService(repos.Disease),
		Diagnosis:     NewDiagnosisService(repos.Diagnosis),
	}
}
