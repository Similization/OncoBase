package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type DiseaseService struct {
	repo repository.Disease
}

func NewDiseaseService(repo repository.Disease) *DiseaseService {
	return &DiseaseService{repo: repo}
}

func (s *DiseaseService) CreateDisease(disease model.Disease) (model.Disease, error) {
	return s.repo.CreateDisease(disease)
}
func (s *DiseaseService) GetDiseaseById(id string) (model.Disease, error) {
	return s.repo.GetDiseaseById(id)
}
func (s *DiseaseService) GetDiseaseList() ([]model.Disease, error) {
	return s.repo.GetDiseaseList()
}
func (s *DiseaseService) UpdateDisease(disease model.Disease) (model.Disease, error) {
	return s.repo.UpdateDisease(disease)
}
func (s *DiseaseService) DeleteDisease(id string) error {
	return s.repo.DeleteDisease(id)
}
