package Models

import (
	"time"

	"gorm.io/gorm"
)

type Img struct {
	ID        uint   `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name      string `json:"name"`
	Filename  string `json:"filename"`
	ProductId string `json:"product_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (b *Img) TableName() string {
	return "imgs"
}
