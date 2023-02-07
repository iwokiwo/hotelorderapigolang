package payment

import "gorm.io/gorm"

type Repository interface {
	Save(item PaymentType) (PaymentType, error)
	Update(item PaymentType) (PaymentType, error)
	FindById(id int) (PaymentType, error)
	Delete(item PaymentType) (bool, error)
	FindAll(id int) ([]PaymentType, error)
	FindAllByBranch(branch_id GetPaymentTypeByBranch) ([]PaymentType, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(item PaymentType) (PaymentType, error) {
	err := r.db.Create(&item).Error

	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *repository) Update(item PaymentType) (PaymentType, error) {
	err := r.db.Save(&item).Error

	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *repository) FindAllByBranch(branch_id GetPaymentTypeByBranch) ([]PaymentType, error) {
	var items []PaymentType

	err := r.db.Where("branch_id = ?", branch_id.BranchID).Preload("Branch").Find(&items).Error

	if err != nil {
		return items, err
	}

	return items, nil
}

func (r *repository) FindAll(id int) ([]PaymentType, error) {
	var items []PaymentType

	err := r.db.Where("user_id = ?", id).Preload("Branch").Find(&items).Error

	if err != nil {
		return items, err
	}

	return items, nil
}

func (r *repository) FindById(id int) (PaymentType, error) {
	var item PaymentType

	err := r.db.Where("id = ?", id).Find(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil

}
func (r *repository) FindBySlug(slug string) (PaymentType, error) {
	var item PaymentType

	err := r.db.Where("slug = ?", slug).Find(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil

}

func (r *repository) Delete(item PaymentType) (bool, error) {
	err := r.db.Delete(&item, item.ID).Error

	if err != nil {
		return false, err
	}

	return true, err
}
