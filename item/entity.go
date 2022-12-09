package item

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID             int      `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name           string   `json:"name"`
	Slug           string   `json:"slug"`
	Bahan          string   `json:"bahan"`
	Thumbnail      string   `json:"thumbnail"`
	Dimensi        string   `json:"dimensi"`
	Hpp            int      `json:"hpp"`
	Price          int      `json:"price"`
	AvailableColor int      `json:"available_color"`
	AvailableSize  int      `json:"available_size"`
	Color          string   `json:"color"`
	Size           string   `json:"size"`
	Stock          int      `json:"stock"`
	Active         int      `json:"active"`
	CategoryId     int      `json:"category_id"`
	UnitId         int      `json:"unit_id"`
	Views          int      `json:"views"`
	Description    string   `json:"description"`
	UserId         int      `json:"user_id"`
	BranchId       int      `json:"branch_id"`
	Path           string   `json:"path"`
	Category       Category `gorm:"foreignkey:CategoryId"`
	Unit           Unit     `gorm:"foreignkey:UnitId"`
	Img            []Img    `gorm:"foreignkey:ProductId"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

type Unit struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Img struct {
	ID        int
	Filename  string
	ProductId int
}
