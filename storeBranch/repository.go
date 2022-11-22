package storebranch

import "gorm.io/gorm"

type Repository interface {
	CreateStore(store Store) (Store, error)
	UpdateStore(store Store) (Store, error)
	DeleteStore(store Store) (Store, error)
	SearchAllStore(userId int) ([]Store, error)

	CreateBranch(branch Branch) (Branch, error)
	UpdateBranch(branch Branch) (Branch, error)
	DeleteBranch(branch Branch) (Branch, error)
	SearchAllBranch(userId int) ([]Branch, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateBranch(branch Branch) (Branch, error) {
	err := r.db.Create(&branch).Error

	if err != nil {
		return branch, err
	}

	return branch, nil
}

func (r *repository) UpdateBranch(branch Branch) (Branch, error) {
	err := r.db.Save(&branch).Error

	if err != nil {
		return branch, err
	}

	return branch, nil
}

func (r *repository) DeleteBranch(branch Branch) (Branch, error) {
	err := r.db.Delete(&branch, branch.ID).Error

	if err != nil {
		return branch, err
	}

	return branch, nil
}

func (r *repository) SearchAllBranch(userId int) ([]Branch, error) {
	var branchs []Branch

	err := r.db.Where("user_id", userId).Find(&branchs).Error

	if err != nil {
		return branchs, err
	}

	return branchs, nil
}

func (r *repository) CreateStore(store Store) (Store, error) {
	err := r.db.Create(&store).Error

	if err != nil {
		return store, err
	}

	return store, nil
}

func (r *repository) UpdateStore(store Store) (Store, error) {
	err := r.db.Save(&store).Error

	if err != nil {
		return store, err
	}

	return store, nil
}

func (r *repository) DeleteStore(store Store) (Store, error) {
	err := r.db.Delete(&store, store.ID).Error

	if err != nil {
		return store, err
	}

	return store, nil
}

func (r *repository) SearchAllStore(userId int) ([]Store, error) {
	var stores []Store

	err := r.db.Where("user_id", userId).Find(&stores, userId).Error

	if err != nil {
		return stores, err
	}

	return stores, nil
}
