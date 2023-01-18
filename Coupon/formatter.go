package coupon

import "time"

type CouponFormatter struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	DiscountType string    `json:"discount_type"`
	Discount     int       `json:"discount"`
	MaxValue     int       `json:"max_value"`
	MinValue     int       `json:"min_value"`
	MaxItem      int       `json:"max_item"`
	MinItem      int       `json:"min_item"`
	Description  string    `json:"description"`
	ValidFrom    time.Time `json:"valid_from"`
	ValidUntil   time.Time `json:"valid_until"`
	Active       int       `json:"active"`
	Branch       Branch    `json:"branch"`
	UserId       int       `json:"user_id"`
	Limit        int       `json:"limit"`
	UseLimit     int       `json:"use_limit"`
}

func FormatCoupon(coupon Coupon) CouponFormatter {
	formatter := CouponFormatter{
		ID:           coupon.ID,
		Name:         coupon.Name,
		DiscountType: coupon.DiscountType,
		Discount:     coupon.Discount,
		MaxValue:     coupon.MaxValue,
		MinValue:     coupon.MinValue,
		MaxItem:      coupon.MaxItem,
		MinItem:      coupon.MinItem,
		Description:  coupon.Description,
		ValidFrom:    coupon.ValidFrom,
		ValidUntil:   coupon.ValidUntil,
		Active:       coupon.Active,
		Branch:       coupon.Branch,
		Limit:        coupon.Limit,
		UseLimit:     coupon.UseLimit,
	}

	return formatter
}

func FormatCoupons(unit []Coupon) []CouponFormatter {
	Formatter := []CouponFormatter{}
	for _, unit := range unit {
		formatter := FormatCoupon(unit)
		Formatter = append(Formatter, formatter)
	}
	return Formatter
}
