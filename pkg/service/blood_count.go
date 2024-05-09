package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type BloodCountService struct {
	repo repository.BloodCount
}

func NewBloodCountService(repo repository.BloodCount) *BloodCountService {
	return &BloodCountService{repo: repo}
}

func (s *BloodCountService) CreateBloodCount(bloodCount model.BloodCount) (model.BloodCount, error) {
	return s.repo.CreateBloodCount(bloodCount)
}
func (s *BloodCountService) GetBloodCountById(id string) (model.BloodCount, error) {
	return s.repo.GetBloodCountById(id)
}
func (s *BloodCountService) GetBloodCountList() ([]model.BloodCount, error) {
	return s.repo.GetBloodCountList()
}
func (s *BloodCountService) UpdateBloodCount(bloodCount model.BloodCount) (model.BloodCount, error) {
	return s.repo.UpdateBloodCount(bloodCount)
}
func (s *BloodCountService) DeleteBloodCount(id string) error {
	return s.repo.DeleteBloodCount(id)
}
