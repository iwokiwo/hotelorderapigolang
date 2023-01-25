package payment

import "strings"

type Service interface {
	CreatePaymentType(input CreatePaymentTypeInput) (PaymentType, error)
	UpdatePaymentType(input UpdatePaymentTypeInput) (PaymentType, error)
	GetPaymentTypeById(input GetPaymentType) (PaymentType, error)
	DeletePaymentType(input DeletePaymentTypeInput) (bool, error)
	FindAllPaymentType() ([]PaymentType, error)
	FindAllPaymentTypeByBranch(branch_id GetPaymentTypeByBranch) ([]PaymentType, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetPaymentTypeById(input GetPaymentType) (PaymentType, error) {
	item, err := s.repository.FindById(input.ID)

	if err != nil {
		return item, err
	}

	return item, nil
}

func (s *service) FindAllPaymentTypeByBranch(baranch_id GetPaymentTypeByBranch) ([]PaymentType, error) {
	item, err := s.repository.FindAllByBranch(baranch_id)

	if err != nil {
		return item, err
	}

	return item, nil
}

func (s *service) FindAllPaymentType() ([]PaymentType, error) {
	unit, err := s.repository.FindAll()

	if err != nil {
		return unit, err
	}

	return unit, nil
}

func (s *service) CreatePaymentType(input CreatePaymentTypeInput) (PaymentType, error) {
	item := PaymentType{}
	item.Name = input.Name
	item.BranchId = input.BranchId
	item.Slug = strings.ReplaceAll(strings.ToLower(input.Name), " ", "-")

	newItem, err := s.repository.Save(item)
	if err != nil {
		return newItem, err
	}

	return item, nil

}

func (s *service) UpdatePaymentType(input UpdatePaymentTypeInput) (PaymentType, error) {
	item, err := s.repository.FindById(input.Id)
	if err != nil {
		return item, err
	}

	item.Name = input.Name
	item.BranchId = input.BranchId
	update, err := s.repository.Update(item)
	if err != nil {
		return update, err
	}

	return item, nil

}

func (s *service) DeletePaymentType(input DeletePaymentTypeInput) (bool, error) {
	item, err := s.repository.FindById(input.ID)
	if err != nil {
		return false, err
	}

	_, err = s.repository.Delete(input.ID, item)
	if err != nil {
		return false, err
	}
	return true, nil
}
