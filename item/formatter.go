package item

type ItemFormatter struct {
	ID             int      `json:"id"`
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
	Views          int      `json:"views"`
	Description    string   `json:"description"`
	CategoryId     int      `json:"category_id"`
	Category       Category `json:"category"`
	Unit           Unit     `json:"unit"`
	Img            []Img    `json:"img"`
}

func FormatItem(product Product) ItemFormatter {
	formatter := ItemFormatter{
		ID:             product.ID,
		Name:           product.Name,
		Slug:           product.Slug,
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
		Img:            product.Img,
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
