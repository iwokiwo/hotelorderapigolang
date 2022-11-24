package storebranch

import "time"

type CreateStore struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description"`
	Address     string `form:"address" json:"address"`
	UserId      int    `form:"user_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UpdateStore struct {
	ID int `form:"id" binding:"required"`
	CreateStore
}

type DeleteStore struct {
	ID int `json:"id" binding:"required"`
}

type CreateBranch struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description"`
	Address     string `form:"address"`
	UserId      int    `form:"user_id"`
	StoreId     int    `form:"store_id"`
	Phone       string `form:"phone"`
	Email       string `form:"email"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UpdateBranch struct {
	ID int `form:"id" binding:"required"`
	CreateBranch
}

type DeleteBranch struct {
	ID int `json:"id" binding:"required"`
}
