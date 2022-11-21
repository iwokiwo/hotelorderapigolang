package storebranch

import "time"

type CreateStore struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Address     string `json:"address"`
	UserId      int    `json:"user_id"`
	Logo        string `json:"logo"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UpdateStore struct {
	ID int `json:"id" binding:"required"`
	CreateStore
}

type DeleteStore struct {
	ID int `json:"id" binding:"required"`
}

type CreateBranch struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Address     string `json:"address"`
	UserId      int    `json:"user_id"`
	BranchId    int    `json:"branch_id"`
	Logo        string `json:"logo"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UpdateBranch struct {
	ID int `json:"id" binding:"required"`
	CreateBranch
}

type DeleteBranch struct {
	ID int `json:"id" binding:"required"`
}
