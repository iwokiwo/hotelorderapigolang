package Models

import (
	"time"

	"gorm.io/gorm"
)

type Branch struct {
	ID          uint   `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	UserId      int    `json:"user_id"`
	BranchId    int    `json:"branch_id"`
	Logo        string `json:"logo"`
	Path        string `json:"path"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (b *Branch) TableName() string {
	return "branchs"
}
