package category

import (
	"time"
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
	ID          int    `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	UserId      int    `json:"user_id"`
	Logo        string `json:"logo"`
	Url         string `json:"url"`
	Path        string `json:"path"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}
