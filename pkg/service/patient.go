package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type PatientService struct {
	repo repository.Patient
}

func NewPatientService(repo repository.Patient) *PatientService {
	return &PatientService{repo: repo}
}

func (s *PatientService) CreatePatient(patient model.Patient) (model.Patient, error) {
	return s.repo.CreatePatient(patient)
}
func (s *PatientService) GetPatientById(id int) (model.Patient, error) {
	return s.repo.GetPatientById(id)
}
func (s *PatientService) GetPatientList() ([]model.Patient, error) {
	return s.repo.GetPatientList()
}
func (s *PatientService) UpdatePatient(patient model.Patient) (model.Patient, error) {
	return s.repo.UpdatePatient(patient)
}
func (s *PatientService) DeletePatient(id int) error {
	return s.repo.DeletePatient(id)
}
