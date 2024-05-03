package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type DoctorPatientService struct {
	repo repository.DoctorPatient
}

func NewDoctorPatientService(repo repository.DoctorPatient) *DoctorPatientService {
	return &DoctorPatientService{repo: repo}
}

func (s *DoctorPatientService) CreateDoctorPatient(doctorPatient model.DoctorPatient) (model.DoctorPatient, error) {
	return s.repo.CreateDoctorPatient(doctorPatient)
}
func (s *DoctorPatientService) GetDoctorPatientList(doctor_id int) ([]model.DoctorPatient, error) {
	return s.repo.GetDoctorPatientList(doctor_id)
}
func (s *DoctorPatientService) DeleteDoctorPatient(doctorId, patientId int) error {
	return s.repo.DeleteDoctorPatient(model.DoctorPatient{Patient: patientId, Doctor: doctorId})
}
