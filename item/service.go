package item

import (
	"mime/multipart"
)

type Service interface {
	CreateItem(input CreateItem, userId int, filename string, path string) (Product, error)
	UpdateItem(input UpdateItem, userId int, filename string, path string, File []*multipart.FileHeader, deleteFile []string) (Product, error)
	DeleteItem(input DeleteItem) (Product, error)
	SearchAll(input SearchInput, userId int) ([]Product, int64, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateItem(input CreateItem, UserId int, filename string, path string) (Product, error) {
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
	item.Thumbnail = filename
	item.Path = path

	newItem, err := s.repository.CreateItem(item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (s *service) UpdateItem(input UpdateItem, UserId int, filename string, path string, File []*multipart.FileHeader, deleteFile []string) (Product, error) {
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
	item.Thumbnail = filename
	item.Path = path
	item.UserId = UserId

	//fmt.Println("img", FormatInputImgs(File, path, input.ID))

	newItem, err := s.repository.UpdateItem(item, FormatInputImgs(File, input.ID), deleteFile)
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
