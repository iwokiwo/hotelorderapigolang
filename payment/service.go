package payment

import "strings"

type Service interface {
	CreatePaymentType(input CreatePaymentTypeInput, user_id int) (PaymentType, error)
	UpdatePaymentType(input UpdatePaymentTypeInput) (PaymentType, error)
	GetPaymentTypeById(input GetPaymentType) (PaymentType, error)
	DeletePaymentType(input DeletePaymentTypeInput) (bool, error)
	FindAllServicePaymentType(id int) ([]PaymentType, error)
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

func (s *service) FindAllServicePaymentType(id int) ([]PaymentType, error) {
	item, err := s.repository.FindAll(id)

	if err != nil {
		return item, err
	}

	return item, nil
}

func (s *service) CreatePaymentType(input CreatePaymentTypeInput, user_id int) (PaymentType, error) {
	item := PaymentType{}
	item.Name = input.Name
	item.BranchId = input.BranchId
	item.Slug = strings.ReplaceAll(strings.ToLower(input.Name), " ", "-")
	item.UserId = user_id

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
	item := PaymentType{}
	item.ID = int(input.ID)

	newItem, err := s.repository.Delete(item)
	if err != nil {
		return newItem, err
	}
	return true, nil
}
