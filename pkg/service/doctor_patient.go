package services

import (
	"med/pkg/model"
	"med/pkg/repository"

	"github.com/guregu/null/v5"
)

type DoctorPatientService struct {
	repo repository.DoctorPatient
}

func NewDoctorPatientService(repo repository.DoctorPatient) *DoctorPatientService {
	return &DoctorPatientService{repo: repo}
}

func (s *DoctorPatientService) CreateDoctorPatient(doctorPatient model.DoctorPatient) error {
	return s.repo.CreateDoctorPatient(doctorPatient)
}
func (s *DoctorPatientService) GetDoctorPatientListByDoctor(doctor_id int) ([]model.DoctorPatient, error) {
	return s.repo.GetDoctorPatientListByDoctor(doctor_id)
}
func (s *DoctorPatientService) GetDoctorPatientListByPatient(patientId int) ([]model.DoctorPatient, error) {
	return s.repo.GetDoctorPatientListByPatient(patientId)
}
func (s *DoctorPatientService) DeleteDoctorPatient(doctorId, patientId int) error {
	return s.repo.DeleteDoctorPatient(model.DoctorPatient{
		Patient: null.IntFrom(int64(patientId)),
		Doctor:  null.IntFrom(int64(doctorId)),
	})
}
