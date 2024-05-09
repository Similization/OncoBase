package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type UnitMeasureService struct {
	repo repository.UnitMeasure
}

func NewUnitMeasureService(repo repository.UnitMeasure) *UnitMeasureService {
	return &UnitMeasureService{repo: repo}
}

func (s *UnitMeasureService) CreateUnitMeasure(unitMeasure model.UnitMeasure) error {
	return s.repo.CreateUnitMeasure(unitMeasure)
}
func (s *UnitMeasureService) GetUnitMeasureById(id string) (model.UnitMeasure, error) {
	return s.repo.GetUnitMeasureById(id)
}
func (s *UnitMeasureService) GetUnitMeasureList() ([]model.UnitMeasure, error) {
	return s.repo.GetUnitMeasureList()
}
func (s *UnitMeasureService) UpdateUnitMeasure(unitMeasure model.UnitMeasure) error {
	return s.repo.UpdateUnitMeasure(unitMeasure)
}
func (s *UnitMeasureService) DeleteUnitMeasure(id string) error {
	return s.repo.DeleteUnitMeasure(id)
}
