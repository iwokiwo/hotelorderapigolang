package Models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID             uint   `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Sku            string `json:"sku"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	Thumbnail      string `json:"thumbnail"`
	Bahan          string `json:"bahan"`
	Dimensi        string `json:"dimensi"`
	Hpp            int    `json:"hpp"`
	Price          int    `json:"price"`
	AvailableColor int    `json:"available_color"`
	AvailableSize  int    `json:"available_size"`
	Color          string `json:"color"`
	Size           string `json:"size"`
	Stock          int    `json:"stock"`
	Active         int    `json:"active"`
	Inventory      int    `json:"inventory"`
	Views          int    `json:"views"`
	Description    string `json:"description"`
	// CategoryRelation []CategoryRelation `gorm:"foreignKey:ProductID"`
	// SliderRelation   []SliderRelation   `gorm:"foreignKey:ProductID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (b *Product) TableName() string {
	return "products"
}
