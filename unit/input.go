package unit

type RegisterUnitInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateUnitInput struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type DeleteUnitInput struct {
	ID int `json:"id" binding:"required"`
}

type GetUnitDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
