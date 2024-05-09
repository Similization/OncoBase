package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type ProcedureBloodCountService struct {
	repo repository.ProcedureBloodCount
}

func NewProcedureBloodCountService(repo repository.ProcedureBloodCount) *ProcedureBloodCountService {
	return &ProcedureBloodCountService{repo: repo}
}

func (s *ProcedureBloodCountService) CreateProcedureBloodCount(procedureBloodCount model.ProcedureBloodCount) error {
	return s.repo.CreateProcedureBloodCount(procedureBloodCount)
}
func (s *ProcedureBloodCountService) GetProcedureBloodCountById(procedureId int, bloodCountId string) (model.ProcedureBloodCount, error) {
	return s.repo.GetProcedureBloodCountById(procedureId, bloodCountId)
}
func (s *ProcedureBloodCountService) GetProcedureBloodCountListByProcedure(procedureId int) ([]model.ProcedureBloodCount, error) {
	return s.repo.GetProcedureBloodCountListByProcedure(procedureId)
}
func (s *ProcedureBloodCountService) GetProcedureBloodCountListByBloodCount(bloodCountId string) ([]model.ProcedureBloodCount, error) {
	return s.repo.GetProcedureBloodCountListByBloodCount(bloodCountId)
}
func (s *ProcedureBloodCountService) GetProcedureBloodCountList() ([]model.ProcedureBloodCount, error) {
	return s.repo.GetProcedureBloodCountList()
}
func (s *ProcedureBloodCountService) UpdateProcedureBloodCount(procedureBloodCount model.ProcedureBloodCount) error {
	return s.repo.UpdateProcedureBloodCount(procedureBloodCount)
}
func (s *ProcedureBloodCountService) DeleteProcedureBloodCount(procedureId int, bloodCountId string) error {
	return s.repo.DeleteProcedureBloodCount(procedureId, bloodCountId)
}
