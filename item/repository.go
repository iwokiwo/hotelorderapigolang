package item

import (
	"gorm.io/gorm"
	"strings"
)

type Repository interface {
	CreateItem(item Product) (Product, error)
	UpdateItem(item Product) (Product, error)
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

func (r *repository) UpdateItem(item Product) (Product, error) {
	err := r.db.Save(&item).Error

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
			Where("user_id = ?", userId).
			Preload("Category").
			Preload("Unit").
			Preload("Img").
			Find(&items)
	}

	if input.CategoryID > 0 {

		find = find.Where("active = ?", input.Active).
			Where("category_id = ?", input.CategoryID).
			Where("stock > ?", 0).
			Where("name LIKE ?", "%"+input.Search+"%").
			Where("user_id = ?", userId).
			Preload("Category").
			Preload("Unit").
			Preload("Img").
			Find(&items)
	}
	if input.CategoryID == 0 {
		//fmt.Println("search", input.Search)
		find = find.Where("active = ?", input.Active).
			Where("user_id = ?", userId).
			Where("stock > ?", 0).
			Where("LOWER(name) LIKE ?", "%"+strings.ToLower(input.Search)+"%").
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

	return data, total, nil

}
