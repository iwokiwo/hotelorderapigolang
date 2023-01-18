package coupon

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	ID           uint      `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name         string    `json:"name"`
	DiscountType string    `json:"discount_type"`
	Discount     int       `json:"discount"`
	MaxValue     int       `json:"max_value"`
	MinValue     int       `json:"min_value"`
	MaxItem      int       `json:"max_item"`
	MinItem      int       `json:"min_item"`
	BranchID     int       `json:"branch_id"`
	Description  string    `json:"description"`
	ValidFrom    time.Time `json:"valid_from"`
	ValidUntil   time.Time `json:"valid_until"`
	Filename     string    `json:"filename"`
	Path         string    `json:"path"`
	Active       int       `json:"active"`
	UserId       int       `json:"user_id"`
	Limit        int       `json:"limit"`
	UseLimit     int       `json:"use_limit"`
	Branch       Branch    `gorm:"foreignkey:BranchID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type Branch struct {
	ID          int    `json:"id"`
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
