package category

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        int
	Name      string
	Slug      string
	BranchId  int
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
