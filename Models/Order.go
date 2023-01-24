package Models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           uint `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	CouponId     int
	Discount     int
	DiscountType string
	SubTotal     int
	Tax          int
	Service      int
	GrandTotal   int
	OrderStatus  int
	UserId       int
	BranchId     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

func (b *Order) TableName() string {
	return "orders"
}
