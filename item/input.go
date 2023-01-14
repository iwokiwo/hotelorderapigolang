package item

import (
	"mime/multipart"
	"os"
	"time"
)

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
	Hpp            int    `form:"price"`
	Price          int    `form:"sale_price"`
	AvailableColor int    `form:"available_color"`
	AvailableSize  int    `form:"available_size"`
	Color          string `form:"color"`
	Size           string `form:"size"`
	Stock          int    `form:"quantity"`
	Active         int    `form:"active"`
	CategoryId     int    `form:"category_id"`
	UnitId         int    `form:"unit_id"`
	Description    string `form:"description"`
	BranchId       int    `form:"branch_id"`
	//Gallery        []Img  `form:"gallery"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateItem struct {
	ID int `form:"id" binding:"required"`
	CreateItem
}

type DeleteItem struct {
	ID int `uri:"id" binding:"required"`
}

// type Imgs struct {
// 	Filename string
// 	Path     string
// }

func FormatInputImg(filename string, producId int) Img {
	formatter := Img{
		Filename:  filename,
		Path:      os.Getenv("IMG_GALLERY"),
		ProductId: producId,
	}

	return formatter
}

func FormatInputImgs(imgs []*multipart.FileHeader, productId int) []Img {
	imgFormatter := []Img{}

	for _, img := range imgs {
		formatter := FormatInputImg(img.Filename, productId)
		imgFormatter = append(imgFormatter, formatter)
	}

	return imgFormatter
}
