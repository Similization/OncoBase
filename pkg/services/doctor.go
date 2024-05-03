package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type DoctorService struct {
	repo repository.Doctor
}

func NewDoctorService(repo repository.Doctor) *DoctorService {
	return &DoctorService{repo: repo}
}

func (s *DoctorService) CreateDoctor(doctor model.Doctor) (model.Doctor, error) {
	return s.repo.CreateDoctor(doctor)
}
func (s *DoctorService) GetDoctorById(id int) (model.Doctor, error) {
	return s.repo.GetDoctorById(id)
}
func (s *DoctorService) GetDoctorList() ([]model.Doctor, error) {
	return s.repo.GetDoctorList()
}
func (s *DoctorService) UpdateDoctor(doctor model.Doctor) (model.Doctor, error) {
	return s.repo.UpdateDoctor(doctor)
}
func (s *DoctorService) DeleteDoctor(id int) error {
	return s.repo.DeleteDoctor(id)
}
