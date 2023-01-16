package unit

import "strings"

type Service interface {
	RegisterUnit(input RegisterUnitInput) (Unit, error)
	UpdateUnit(input UpdateUnitInput) (Unit, error)
	GetUnitById(input GetUnitDetailInput) (Unit, error)
	DeleteUnit(input DeleteUnitInput) (bool, error)
	FindAllUnit() ([]Unit, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUnitById(input GetUnitDetailInput) (Unit, error) {
	unit, err := s.repository.FindById(input.ID)

	if err != nil {
		return unit, err
	}

	return unit, nil
}

func (s *service) FindAllUnit() ([]Unit, error) {
	unit, err := s.repository.FindAll()

	if err != nil {
		return unit, err
	}

	return unit, nil
}

func (s *service) RegisterUnit(input RegisterUnitInput) (Unit, error) {
	unit := Unit{}
	unit.Name = input.Name
	unit.BranchId = input.BranchId
	unit.Slug = strings.ReplaceAll(strings.ToLower(input.Name), " ", "-")

	newUnit, err := s.repository.Save(unit)
	if err != nil {
		return newUnit, err
	}

	return newUnit, nil

}

func (s *service) UpdateUnit(input UpdateUnitInput) (Unit, error) {
	unit, err := s.repository.FindById(input.Id)
	if err != nil {
		return unit, err
	}

	unit.Name = input.Name
	unit.BranchId = input.BranchId
	update, err := s.repository.Update(unit)
	if err != nil {
		return update, err
	}

	return update, nil

}

func (s *service) DeleteUnit(input DeleteUnitInput) (bool, error) {
	unit, err := s.repository.FindById(input.ID)
	if err != nil {
		return false, err
	}

	_, err = s.repository.Delete(input.ID, unit)
	if err != nil {
		return false, err
	}
	return true, nil
}
