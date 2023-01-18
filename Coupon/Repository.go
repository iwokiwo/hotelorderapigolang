package coupon

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(coupon Coupon) (Coupon, error)
	Update(coupon Coupon) (Coupon, error)
	Delete(id int, input Coupon) (bool, error)
	FindAll(id int) ([]Coupon, error)
	FindAllByBranch(id SearchInput) ([]Coupon, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(coupon Coupon) (Coupon, error) {
	err := r.db.Create(&coupon).Error

	if err != nil {
		return coupon, err
	}
	return coupon, nil
}

func (r *repository) Update(coupon Coupon) (Coupon, error) {
	err := r.db.Save(&coupon).Error

	if err != nil {
		return coupon, err
	}

	return coupon, nil
}

func (r *repository) FindAll(id int) ([]Coupon, error) {
	var coupons []Coupon

	err := r.db.Where("user_id = ?", id).Preload("Branch").Find(&coupons).Error

	if err != nil {
		return coupons, err
	}

	return coupons, nil
}

func (r *repository) FindAllByBranch(input SearchInput) ([]Coupon, error) {
	var coupons []Coupon

	err := r.db.Where("branch_id = ?", input.BranchID).Preload("Branch").Find(&coupons).Error

	if err != nil {
		return coupons, err
	}

	return coupons, nil
}

func (r *repository) Delete(id int, coupon Coupon) (bool, error) {
	err := r.db.Where("id = ?", id).Delete(&coupon).Error

	if err != nil {
		return false, err
	}

	return true, err
}
