package item

import "time"

type PaginationInput struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Sort      string `json:"sort"`
	Direction string `json:"direction"`
	Active    int    `json:"active"`
	Stock     int    `json:"stock"`
}

type SearchInput struct {
	Page       int    `json:"page"`
	Size       int    `json:"size"`
	Sort       string `json:"sort"`
	Direction  string `json:"direction"`
	Active     int    `json:"active"`
	Stock      int    `json:"stock"`
	Id         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	Search     string `json:"search"`
}

type CreateItem struct {
	Name           string `form:"name" binding:"required"`
	Dimensi        string `form:"dimensi"`
	Hpp            int    `form:"hpp"`
	Price          int    `form:"price"`
	AvailableColor int    `form:"available_color"`
	AvailableSize  int    `form:"available_size"`
	Color          string `form:"color"`
	Size           string `form:"size"`
	Stock          int    `form:"stock"`
	Active         int    `form:"active"`
	CategoryId     int    `form:"category_id"`
	UnitId         int    `form:"unit_id"`
	Description    string `form:"description"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UpdateItem struct {
	ID int `json:"id" binding:"required"`
	CreateItem
}

type DeleteItem struct {
	ID int `json:"id" binding:"required"`
}

type Imgs struct {
	Filename  string
	ProductId int
}
