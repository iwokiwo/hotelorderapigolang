package unit

import (
	"time"

	"gorm.io/gorm"
)

type Unit struct {
	ID        int
	Name      string
	Slug      string
	BranchId  int    `json:"branch_id"`
	Branch    Branch `gorm:"foreignkey:BranchId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Branch struct {
	ID          int
	Name        string
	Description string
	Address     string
	UserId      int
	StoreId     int
	Logo        string
	Path        string
	Phone       string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
