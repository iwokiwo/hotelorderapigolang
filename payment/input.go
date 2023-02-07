package payment

type CreatePaymentTypeInput struct {
	Name     string `json:"name" binding:"required"`
	BranchId int    `json:"branch_id"`
}

type UpdatePaymentTypeInput struct {
	Id       int    `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	BranchId int    `json:"branch_id"`
}

type DeletePaymentTypeInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetPaymentType struct {
	ID int `uri:"id" binding:"required"`
}

type GetPaymentTypeByBranch struct {
	BranchID int `uri:"branch_id"`
}
