package item

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
