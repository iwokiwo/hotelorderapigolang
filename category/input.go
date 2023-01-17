package category

type RegisterCategoryInput struct {
	Name     string `json:"name" binding:"required"`
	BranchId int    `json:"branch_id"`
}

type UpdateCategoryInput struct {
	Id       int    `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	BranchId int    `json:"branch_id"`
}

type DeleteCategoryInput struct {
	ID int `json:"id" binding:"required"`
}

type GetCategoryDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetCategorySlugInput struct {
	Slug string `uri:"slug" binding:"required"`
}

type SearchInput struct {
	BranchID int `uri:"branch_id"`
}
