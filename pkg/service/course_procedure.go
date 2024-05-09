package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type CourseProcedureService struct {
	repo repository.CourseProcedure
}

func NewCourseProcedureService(repo repository.CourseProcedure) *CourseProcedureService {
	return &CourseProcedureService{repo: repo}
}

func (s *CourseProcedureService) CreateCourseProcedure(courseProcedure model.CourseProcedure) (int, error) {
	return s.repo.CreateCourseProcedure(courseProcedure)
}
func (s *CourseProcedureService) GetCourseProcedureById(id string) (model.CourseProcedure, error) {
	return s.repo.GetCourseProcedureById(id)
}
func (s *CourseProcedureService) GetCourseProcedureList() ([]model.CourseProcedure, error) {
	return s.repo.GetCourseProcedureList()
}
func (s *CourseProcedureService) UpdateCourseProcedure(courseProcedure model.CourseProcedure) error {
	return s.repo.UpdateCourseProcedure(courseProcedure)
}
func (s *CourseProcedureService) DeleteCourseProcedure(id string) error {
	return s.repo.DeleteCourseProcedure(id)
}
