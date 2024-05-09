package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type DiagnosisService struct {
	repo repository.Diagnosis
}

func NewDiagnosisService(repo repository.Diagnosis) *DiagnosisService {
	return &DiagnosisService{repo: repo}
}

func (s *DiagnosisService) CreateDiagnosis(diagnosis model.Diagnosis) (model.Diagnosis, error) {
	return s.repo.CreateDiagnosis(diagnosis)
}
func (s *DiagnosisService) GetDiagnosisById(id string) (model.Diagnosis, error) {
	return s.repo.GetDiagnosisById(id)
}
func (s *DiagnosisService) GetDiagnosisList() ([]model.Diagnosis, error) {
	return s.repo.GetDiagnosisList()
}
func (s *DiagnosisService) UpdateDiagnosis(diagnosis model.Diagnosis) (model.Diagnosis, error) {
	return s.repo.UpdateDiagnosis(diagnosis)
}
func (s *DiagnosisService) DeleteDiagnosis(id string) error {
	return s.repo.DeleteDiagnosis(id)
}
