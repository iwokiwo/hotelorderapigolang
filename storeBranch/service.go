package storebranch

type Service interface {
	CreateStore(input CreateStore, userId int) (Store, error)
	UpdateStore(input UpdateStore) (Store, error)
	DeleteStore(input DeleteStore) (Store, error)
	SearchAllStore(userId int) ([]Store, error)

	CreateBranch(input CreateBranch, userId int) (Branch, error)
	UpdateBranch(input UpdateBranch) (Branch, error)
	DeleteBranch(input DeleteBranch) (Branch, error)
	SearchAllBranch(userId int) ([]Branch, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateStore(input CreateStore, UserId int) (Store, error) {
	item := Store{}
	item.Name = input.Name
	item.Address = input.Address
	item.Description = input.Description
	item.UserId = UserId

	newItem, err := s.repository.CreateStore(item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (s *service) UpdateStore(input UpdateStore) (Store, error) {
	item := Store{}
	item.ID = input.ID
	item.Name = input.Name
	item.Address = input.Address
	item.Description = input.Description

	newItem, err := s.repository.UpdateStore(item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (s *service) DeleteItem(input DeleteStore) (Store, error) {
	item := Store{}
	item.ID = input.ID

	newItem, err := s.repository.DeleteStore(item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (s *service) SearchAllStore(userId int) ([]Store, error) {
	items, err := s.repository.SearchAllStore(userId)
	if err != nil {
		return items, err
	}

	return items, nil
}

func (s *service) CreateBranch(input CreateBranch, UserId int) (Branch, error) {
	item := Branch{}
	item.Name = input.Name
	item.Address = input.Address
	item.Description = input.Description
	item.UserId = UserId

	newItem, err := s.repository.CreateBranch(item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (s *service) UpdateBranch(input UpdateBranch) (Branch, error) {
	item := Branch{}
	item.ID = input.ID
	item.Name = input.Name
	item.Address = input.Address
	item.Description = input.Description

	newItem, err := s.repository.UpdateBranch(item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (s *service) DeleteBranch(input DeleteBranch) (Branch, error) {
	item := Branch{}
	item.ID = input.ID

	newItem, err := s.repository.DeleteBranch(item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (s *service) SearchAllBranch(userId int) ([]Branch, error) {
	items, err := s.repository.SearchAllBranch(userId)
	if err != nil {
		return items, err
	}

	return items, nil
}
