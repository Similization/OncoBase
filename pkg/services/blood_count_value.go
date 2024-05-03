package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type BloodCountValueService struct {
	repo repository.BloodCountValue
}

func NewBloodCountValueService(repo repository.BloodCountValue) *BloodCountValueService {
	return &BloodCountValueService{repo: repo}
}

func (s *BloodCountValueService) CreateBloodCountValue(bloodCountValue model.BloodCountValue) (model.BloodCountValue, error) {
	return s.repo.CreateBloodCountValue(bloodCountValue)
}
func (s *BloodCountValueService) GetBloodCountValueById(diseaseId, bloodCountId string) (model.BloodCountValue, error) {
	return s.repo.GetBloodCountValueById(diseaseId, bloodCountId)
}
func (s *BloodCountValueService) GetBloodCountValueListByDisease(diseaseId string) ([]model.BloodCountValue, error) {
	return s.repo.GetBloodCountValueListByDisease(diseaseId)
}
func (s *BloodCountValueService) GetBloodCountValueListByBloodCount(bloodCountId string) ([]model.BloodCountValue, error) {
	return s.repo.GetBloodCountValueListByBloodCount(bloodCountId)
}
func (s *BloodCountValueService) GetBloodCountValueList() ([]model.BloodCountValue, error) {
	return s.repo.GetBloodCountValueList()
}
func (s *BloodCountValueService) UpdateBloodCountValue(bloodCountValue model.BloodCountValue) (model.BloodCountValue, error) {
	return s.repo.UpdateBloodCountValue(bloodCountValue)
}
func (s *BloodCountValueService) DeleteBloodCountValue(diseaseId, bloodCountId string) error {
	return s.repo.DeleteBloodCountValue(diseaseId, bloodCountId)
}
