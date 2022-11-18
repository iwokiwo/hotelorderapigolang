package item

type Service interface {
	CreateItem(input CreateItem) (Product, error)
	SearchAll(input SearchInput) ([]Product, int64, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateItem(input CreateItem) (Product, error) {
	item := Product{}
	item.Name = input.Name

	newItem, err := s.repository.CreateItem(item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (s *service) SearchAll(input SearchInput) ([]Product, int64, error) {
	items, total, err := s.repository.SearchAll(input)
	if err != nil {
		return items, total, err
	}
	if total != 0 {
		return items, total, nil
	}
	return items, total, nil
}
