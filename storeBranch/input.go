package storebranch

import "time"

type CreateStore struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description"`
	Address     string `form:"address"`
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
	BranchId    int    `form:"branch_id"`
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
