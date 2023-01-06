package item

import (
	"os"

	"github.com/gosimple/slug"
)

type ItemFormatter struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Slug           string   `json:"slug"`
	Bahan          string   `json:"bahan"`
	Thumbnail      string   `json:"thumbnail"`
	Dimensi        string   `json:"dimensi"`
	Hpp            int      `json:"price"`
	Price          int      `json:"sale_price"`
	AvailableColor int      `json:"available_color"`
	AvailableSize  int      `json:"available_size"`
	Color          string   `json:"color"`
	Size           string   `json:"size"`
	Stock          int      `json:"quantity"`
	Active         int      `json:"active"`
	Views          int      `json:"views"`
	Description    string   `json:"description"`
	CategoryId     int      `json:"category_id"`
	Url            string   `json:"url"`
	Path           string   `json:"path"`
	Category       Category `json:"category"`
	Unit           Unit     `json:"unit"`
	Img            []Img    `json:"gallery"`
}

type ItemFormatterFrontEnd struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	Price       int      `json:"price"`
	SalePrice   int      `json:"sale_price"`
	Quantity    int      `json:"quantity"`
	Description string   `json:"description"`
	Category    Category `json:"category"`
	Unit        Unit     `json:"unit"`
	Gallery     []Img    `json:"gallery"`
	Image       string   `json:"image"`
}

// type ImgFormatter struct {
// 	ID        int
// 	Filename  string
// 	ProductId int
// 	Path      string
// 	Url       string
// }

func FormatGallery(img Img) Img {
	formatter := Img{
		ID:        img.ID,
		Filename:  img.Filename,
		Path:      img.Path,
		ProductId: img.ProductId,
		Url:       os.Getenv("APP_URL"),
	}

	return formatter
}

func FormatGallerys(items []Img) []Img {

	itemFormatter := []Img{}

	for _, item := range items {
		formatter := FormatGallery(item)
		itemFormatter = append(itemFormatter, formatter)
	}

	return itemFormatter
}

func FormatItem(product Product) ItemFormatter {
	formatter := ItemFormatter{
		ID:             product.ID,
		Name:           product.Name,
		Slug:           slug.Make(product.Name),
		Bahan:          product.Bahan,
		Thumbnail:      product.Thumbnail,
		Dimensi:        product.Dimensi,
		Hpp:            product.Hpp,
		Price:          product.Price,
		AvailableColor: product.AvailableColor,
		AvailableSize:  product.AvailableSize,
		Color:          product.Color,
		Size:           product.Size,
		Stock:          product.Stock,
		Active:         product.Active,
		Views:          product.Views,
		Description:    product.Description,
		Category:       product.Category,
		CategoryId:     product.CategoryId,
		Unit:           product.Unit,
		Img:            FormatGallerys(product.Img),
		Path:           product.Path,
		Url:            os.Getenv("APP_URL"),
	}
	return formatter
}

func FormatItemFrontEnd(product Product) ItemFormatterFrontEnd {
	formatter := ItemFormatterFrontEnd{
		ID:          product.ID,
		Name:        product.Name,
		Slug:        slug.Make(product.Name),
		Gallery:     FormatGallerys(product.Img),
		SalePrice:   product.Price,
		Price:       product.Hpp,
		Category:    product.Category,
		Unit:        product.Unit,
		Description: product.Description,
		Image:       os.Getenv("APP_URL") + "" + product.Path + "" + product.Thumbnail,
	}

	return formatter
}

func FormatItems(items []Product) []ItemFormatter {

	itemFormatter := []ItemFormatter{}

	for _, item := range items {
		formatter := FormatItem(item)
		itemFormatter = append(itemFormatter, formatter)
	}

	return itemFormatter
}

func FormatItemsFrontEnds(items []Product) []ItemFormatterFrontEnd {

	itemFormatter := []ItemFormatterFrontEnd{}

	for _, item := range items {
		formatter := FormatItemFrontEnd(item)
		itemFormatter = append(itemFormatter, formatter)
	}

	return itemFormatter
}
