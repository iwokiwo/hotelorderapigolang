package unit

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(unit Unit) (Unit, error)
	Update(unit Unit) (Unit, error)
	FindById(id int) (Unit, error)
	Delete(id int, unit Unit) (bool, error)
	FindAll() ([]Unit, error)
	FindAllByBranch(branch_id SearchInput) ([]Unit, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(unit Unit) (Unit, error) {
	err := r.db.Create(&unit).Error

	if err != nil {
		return unit, err
	}

	return unit, nil
}

func (r *repository) Update(unit Unit) (Unit, error) {
	err := r.db.Save(&unit).Error

	if err != nil {
		return unit, err
	}

	return unit, nil
}

func (r *repository) FindAllByBranch(branch_id SearchInput) ([]Unit, error) {
	var units []Unit

	err := r.db.Where("branch_id = ?", branch_id.BranchID).Preload("Branch").Find(&units).Error

	if err != nil {
		return units, err
	}

	return units, nil
}

func (r *repository) FindAll() ([]Unit, error) {
	var units []Unit

	err := r.db.Preload("Branch").Find(&units).Error

	if err != nil {
		return units, err
	}

	return units, nil
}

func (r *repository) FindById(id int) (Unit, error) {
	var unit Unit

	err := r.db.Where("id = ?", id).Find(&unit).Error
	if err != nil {
		return unit, err
	}

	return unit, nil

}
func (r *repository) FindBySlug(slug string) (Unit, error) {
	var unit Unit

	err := r.db.Where("slug = ?", slug).Find(&unit).Error
	if err != nil {
		return unit, err
	}

	return unit, nil

}

func (r *repository) Delete(id int, unit Unit) (bool, error) {
	err := r.db.Where("id = ?", id).Delete(&unit).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
