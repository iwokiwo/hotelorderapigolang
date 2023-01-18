package coupon

type Service interface {
	CreateService(input CreateCouponInput, user_id int) (Coupon, error)
	UpdateService(input UpdateCouponInput, user_id int) (Coupon, error)
	DeleteService(input DeleteCouponInput) (bool, error)
	FindAllService(id int) ([]Coupon, error)
	FindAllByBranchService(input SearchInput) ([]Coupon, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateService(input CreateCouponInput, user_id int) (Coupon, error) {
	coupon := Coupon{}
	coupon.Name = input.Name
	coupon.DiscountType = input.DiscountType
	coupon.Discount = input.Discount
	coupon.MaxValue = input.MaxValue
	coupon.MinValue = input.MinValue
	coupon.MaxItem = input.MaxItem
	coupon.MinItem = input.MinItem
	coupon.ValidFrom = input.ValidFrom
	coupon.ValidUntil = input.ValidUntil
	coupon.Active = input.Active
	coupon.BranchID = input.BranchID
	coupon.Description = input.Description
	coupon.UserId = user_id
	coupon.Limit = input.Limit
	coupon.UseLimit = input.UseLimit

	newCoupon, err := s.repository.Create(coupon)

	if err != nil {
		return newCoupon, err
	}

	return newCoupon, nil
}

func (s *service) UpdateService(input UpdateCouponInput, user_id int) (Coupon, error) {
	coupon := Coupon{}
	coupon.ID = uint(input.ID)
	coupon.Name = input.Name
	coupon.DiscountType = input.DiscountType
	coupon.Discount = input.Discount
	coupon.MaxValue = input.MaxValue
	coupon.MinValue = input.MinValue
	coupon.MaxItem = input.MaxItem
	coupon.MinItem = input.MinItem
	coupon.ValidFrom = input.ValidFrom
	coupon.ValidUntil = input.ValidUntil
	coupon.Active = input.Active
	coupon.BranchID = input.BranchID
	coupon.Description = input.Description
	coupon.UserId = user_id
	coupon.Limit = input.Limit
	coupon.UseLimit = input.UseLimit

	newCoupon, err := s.repository.Update(coupon)

	if err != nil {
		return newCoupon, err
	}

	return newCoupon, nil
}

func (s *service) FindAllService(id int) ([]Coupon, error) {

	coupon, err := s.repository.FindAll(id)

	if err != nil {
		return coupon, err
	}

	return coupon, nil
}

func (s *service) FindAllByBranchService(input SearchInput) ([]Coupon, error) {

	coupon, err := s.repository.FindAllByBranch(input)

	if err != nil {
		return coupon, err
	}

	return coupon, nil
}

func (s *service) DeleteService(input DeleteCouponInput) (bool, error) {
	coupon := Coupon{}
	coupon.ID = uint(input.ID)

	newItem, err := s.repository.Delete(coupon)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}
