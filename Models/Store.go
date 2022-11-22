package Models

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID          uint   `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	UserId      int    `json:"user_id"`
	Logo        string `json:"logo"`
	Path        string `json:"path"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (b *Store) TableName() string {
	return "stores"
}
