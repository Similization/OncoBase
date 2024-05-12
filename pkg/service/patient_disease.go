package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type PatientDiseaseService struct {
	repo repository.PatientDisease
}

func NewPatientDiseaseService(repo repository.PatientDisease) *PatientDiseaseService {
	return &PatientDiseaseService{repo: repo}
}

func (s *PatientDiseaseService) CreatePatientDisease(patientDisease model.PatientDisease) error {
	return s.repo.CreatePatientDisease(patientDisease)
}
func (s *PatientDiseaseService) GetPatientDiseaseById(patientId int, diseaseId string) (model.PatientDisease, error) {
	return s.repo.GetPatientDiseaseById(patientId, diseaseId)
}
func (s *PatientDiseaseService) GetPatientDiseaseListByDisease(diseaseId string) ([]model.PatientDisease, error) {
	return s.repo.GetPatientDiseaseListByDisease(diseaseId)
}
func (s *PatientDiseaseService) GetPatientDiseaseListByPatient(patientId int) ([]model.PatientDisease, error) {
	return s.repo.GetPatientDiseaseListByPatient(patientId)
}
func (s *PatientDiseaseService) GetPatientDiseaseList() ([]model.PatientDisease, error) {
	return s.repo.GetPatientDiseaseList()
}
func (s *PatientDiseaseService) UpdatePatientDisease(patientDisease model.PatientDisease) error {
	return s.repo.UpdatePatientDisease(patientDisease)
}
func (s *PatientDiseaseService) DeletePatientDisease(diseaseId int, patientId string) error {
	return s.repo.DeletePatientDisease(diseaseId, patientId)
}
