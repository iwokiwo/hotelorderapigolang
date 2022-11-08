package Models

import (
	"time"

	"gorm.io/gorm"
)

type Unit struct {
	ID          uint   `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (b *Unit) TableName() string {
	return "units"
}
