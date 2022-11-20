package Models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint   `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name        string `json:"name"`
	Img         string `json:"img"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	UserId      int    `json:"user_id"`
	BranchId    int    `json:"branch_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (b *Category) TableName() string {
	return "categories"
}
