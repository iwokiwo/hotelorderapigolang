package item

type Service interface {
	CreateItem(input CreateItem, userId int) (Product, error)
	UpdateItem(input UpdateItem) (Product, error)
	DeleteItem(input DeleteItem) (Product, error)
	SearchAll(input SearchInput, userId int) ([]Product, int64, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateItem(input CreateItem, UserId int) (Product, error) {
	item := Product{}
	item.Name = input.Name
	item.Hpp = input.Hpp
	item.Price = input.Price
	item.Stock = input.Stock
	item.Active = input.Active
	item.CategoryId = input.CategoryId
	item.UnitId = input.UnitId
	item.Description = input.Description
	item.UserId = UserId

	newItem, err := s.repository.CreateItem(item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (s *service) UpdateItem(input UpdateItem) (Product, error) {
	item := Product{}
	item.ID = input.ID
	item.Name = input.Name
	item.Hpp = input.Hpp
	item.Price = input.Price
	item.Stock = input.Stock
	item.Active = input.Active
	item.CategoryId = input.CategoryId
	item.UnitId = input.UnitId
	item.Description = input.Description

	newItem, err := s.repository.UpdateItem(item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (s *service) DeleteItem(input DeleteItem) (Product, error) {
	item := Product{}
	item.ID = input.ID

	newItem, err := s.repository.DeleteItem(item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (s *service) SearchAll(input SearchInput, userId int) ([]Product, int64, error) {
	items, total, err := s.repository.SearchAll(input, userId)
	if err != nil {
		return items, total, err
	}
	if total != 0 {
		return items, total, nil
	}
	return items, total, nil
}
