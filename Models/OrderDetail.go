package Models

import (
	"time"

	"gorm.io/gorm"
)

type OrderDetail struct {
	ID          uint `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	OrderId     int
	ItemId      int
	Qty         int
	Price       int
	Discount    int
	TotalAmount int
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (b *OrderDetail) TableName() string {
	return "order_details"
}
