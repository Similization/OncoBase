package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type DrugService struct {
	repo repository.Drug
}

func NewDrugService(repo repository.Drug) *DrugService {
	return &DrugService{repo: repo}
}

func (s *DrugService) CreateDrug(drug model.Drug) (model.Drug, error) {
	return s.repo.CreateDrug(drug)
}
func (s *DrugService) GetDrugById(id string) (model.Drug, error) {
	return s.repo.GetDrugById(id)
}
func (s *DrugService) GetDrugList() ([]model.Drug, error) {
	return s.repo.GetDrugList()
}
func (s *DrugService) UpdateDrug(drug model.Drug) (model.Drug, error) {
	return s.repo.UpdateDrug(drug)
}
func (s *DrugService) DeleteDrug(id string) error {
	return s.repo.DeleteDrug(id)
}
