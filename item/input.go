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
	Name           string `json:"name" binding:"required"`
	Dimensi        string `json:"dimensi"`
	Hpp            int    `json:"hpp"`
	Price          int    `json:"price"`
	AvailableColor int    `json:"available_color"`
	AvailableSize  int    `json:"available_size"`
	Color          string `json:"color"`
	Size           string `json:"size"`
	Stock          int    `json:"stock"`
	Active         int    `json:"active"`
	CategoryId     int    `json:"category_id"`
	UnitId         int    `json:"unit_id"`
	Description    string `json:"description"`
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
