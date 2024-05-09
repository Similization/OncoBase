package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type PatientCourseService struct {
	repo repository.PatientCourse
}

func NewPatientCourseService(repo repository.PatientCourse) *PatientCourseService {
	return &PatientCourseService{repo: repo}
}

func (s *PatientCourseService) CreatePatientCourse(patientCourse model.PatientCourse) (int, error) {
	return s.repo.CreatePatientCourse(patientCourse)
}
func (s *PatientCourseService) GetPatientCourseById(id int) (model.PatientCourse, error) {
	return s.repo.GetPatientCourseById(id)
}
func (s *PatientCourseService) GetPatientCourseList() ([]model.PatientCourse, error) {
	return s.repo.GetPatientCourseList()
}
func (s *PatientCourseService) UpdatePatientCourse(patientCourse model.PatientCourse) error {
	return s.repo.UpdatePatientCourse(patientCourse)
}
func (s *PatientCourseService) DeletePatientCourse(id int) error {
	return s.repo.DeletePatientCourse(id)
}
