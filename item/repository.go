package item

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	// Save(item Item) (Item, error)
	// Update(item Item) (Item, error)
	SearchAll(input SearchInput) ([]Product, int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SearchAll(input SearchInput) ([]Product, int64, error) {

	var items []Product
	totalRows := 0
	total := int64(totalRows)

	page := input.Page

	if page == 0 {
		page = 1
	}

	//query
	offset := (page - 1) * input.Size
	find := r.db.Limit(input.Size).Offset(offset).Order(input.Sort + " " + input.Direction)

	if input.Id > 0 {
		find = find.Where("active = ?", input.Active).
			Where("stock > ?", 0).
			Where("id = ?", input.Id).
			Preload("Category").
			Preload("Unit").
			Preload("Img").
			Find(&items)
	}

	if input.CategoryID > 0 {
		fmt.Println("lewat", input.CategoryID)
		find = find.Where("active = ?", input.Active).
			Where("category_id = ?", input.CategoryID).
			Where("stock > ?", 0).
			Where("name LIKE ?", "%"+input.Search+"%").
			Preload("Category").
			Preload("Unit").
			Preload("Img").
			Find(&items)
	}
	if input.CategoryID == 0 {

		find = find.Where("active = ?", input.Active).
			Where("stock > ?", 0).
			Where("name LIKE ?", "%"+input.Search+"%").
			Or("slug LIKE ?", "%"+input.Search+"%").
			Or("description LIKE ?", "%"+input.Search+"%").
			Preload("Category").
			Preload("Unit").
			Preload("Img").
			Find(&items)

	}

	err := find.Error
	if err != nil {
		return items, total, err
	}

	for i := range items {
		total = total + 1
		i++
	}

	data := items
	fmt.Println("data", data)
	return data, total, nil

}
