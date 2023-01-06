package item

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	CreateItem(item Product) (Product, error)
	UpdateItem(item Product, img []Img) (Product, error)
	DeleteItem(item Product) (Product, error)
	SearchAll(input SearchInput, userId int) ([]Product, int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateItem(item Product) (Product, error) {
	err := r.db.Create(&item).Error

	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *repository) UpdateItem(item Product, img []Img) (Product, error) {
	err := r.db.Save(&item).Error
	errs := r.db.Create(&img).Error
	fmt.Println("image save", errs)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *repository) DeleteItem(item Product) (Product, error) {
	//err := r.db.Where("id = ?", item.ID).Delete(&item).Error
	err := r.db.Delete(&item, item.ID).Error
	if err != nil {
		return item, err
	}
	return item, nil
}

func (r *repository) SearchAll(input SearchInput, userId int) ([]Product, int64, error) {

	//fmt.Println("userId", userId)
	var items []Product
	var totalItems []Product

	page := input.Page

	if page == 0 {
		page = 1
	} else {
		page = page + 1
	}

	//query
	offset := (page - 1) * input.Size

	find := r.db.Limit(input.Size).Offset(offset).Order(input.Sort + " " + input.Direction)

	if input.Id > 0 {
		find = find.Where("active = ?", input.Active).
			Where("id = ?", input.Id).
			Where("user_id = ?", userId).
			Preload("Category").
			Preload("Unit").
			Preload("Img").
			Find(&items)

		r.db.Where("active = ?", input.Active).
			Where("id = ?", input.Id).
			Where("user_id = ?", userId).
			Preload("Category").
			Preload("Unit").
			Preload("Img").
			Find(&totalItems)

	}

	if input.CategoryID > 0 {

		find = find.Where("active = ?", input.Active).
			Where("category_id = ?", input.CategoryID).
			Where("name LIKE ?", "%"+input.Search+"%").
			Where("user_id = ?", userId).
			Preload("Category").
			Preload("Unit").
			Preload("Img").
			Find(&items)

		r.db.Where("active = ?", input.Active).
			Where("category_id = ?", input.CategoryID).
			Where("name LIKE ?", "%"+input.Search+"%").
			Where("user_id = ?", userId).
			Preload("Category").
			Preload("Unit").
			Preload("Img").
			Find(&totalItems)
	}
	if input.CategoryID == 0 {
		//fmt.Println("search", input.Search)
		find = find.Where("active = ?", input.Active).
			Where("user_id = ?", userId).
			Where("LOWER(name) LIKE ?", "%"+strings.ToLower(input.Search)+"%").
			Preload("Category").
			Preload("Unit").
			Preload("Img").
			Find(&items)

		r.db.Where("active = ?", input.Active).
			Where("user_id = ?", userId).
			Where("LOWER(name) LIKE ?", "%"+strings.ToLower(input.Search)+"%").
			Preload("Category").
			Preload("Unit").
			Preload("Img").
			Find(&totalItems)

	}
	fmt.Println("page", len(totalItems))

	total := int64(len(totalItems))
	err := find.Error
	if err != nil {
		return items, total, err
	}

	data := items

	return data, total, nil

}
