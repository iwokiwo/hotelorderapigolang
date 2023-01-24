package Models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID            uint `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	OrderID       int
	PaymentTypeId int
	PaymentType   string
	TotalPayment  int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

func (b *Payment) TableName() string {
	return "payemnts"
}
