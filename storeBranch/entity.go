package storebranch

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID          int    `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	UserId      int    `json:"user_id"`
	Logo        string `json:"logo"`
	Path        string `json:"path"`
	//Branch      []Branch `gorm:"foreignkey:StoreId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Branch struct {
	ID          int    `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	UserId      int    `json:"user_id"`
	StoreId     int    `json:"store_id"`
	Logo        string `json:"logo"`
	Path        string `json:"path"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
