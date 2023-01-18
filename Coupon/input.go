package coupon

import "time"

type CreateCouponInput struct {
	Name         string    `json:"name" binding:"required"`
	DiscountType string    `json:"discount_type" binding:"required"`
	Discount     int       `json:"discount" binding:"required"`
	MaxValue     int       `json:"max_value"`
	MinValue     int       `json:"min_value"`
	MaxItem      int       `json:"max_item"`
	MinItem      int       `json:"min_item"`
	BranchID     int       `json:"branch_id"`
	Description  string    `json:"description"`
	ValidFrom    time.Time `json:"valid_from"`
	ValidUntil   time.Time `json:"valid_until"`
	Active       int       `json:"active"`
	UserId       int       `json:"user_id"`
}

type UpdateCouponInput struct {
	ID int `json:"id" binding:"required"`
	CreateCouponInput
}

type DeleteCouponInput struct {
	ID int `json:"id" binding:"required"`
}

type SearchInput struct {
	BranchID int `uri:"branch_id"`
}
