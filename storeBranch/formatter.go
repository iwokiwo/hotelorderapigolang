package storebranch

import "os"

type StoreFormatter struct {
	ID          int    `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	UserId      int    `json:"user_id"`
	Logo        string `json:"logo"`
	Url         string `json:"url"`
	Path        string `json:"path"`
}

func FormatStore(store Store) StoreFormatter {
	formatter := StoreFormatter{
		ID:          store.ID,
		Name:        store.Name,
		Description: store.Description,
		Address:     store.Address,
		Logo:        store.Logo,
		Url:         os.Getenv("APP_URL"),
		Path:        store.Path,
	}
	return formatter
}

func FormatStores(items []Store) []StoreFormatter {

	storeFormatter := []StoreFormatter{}

	for _, item := range items {
		formatter := FormatStore(item)
		storeFormatter = append(storeFormatter, formatter)
	}

	return storeFormatter
}

type BranchFormatter struct {
	ID          int    `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	UserId      int    `json:"user_id"`
	Logo        string `json:"logo"`
	Url         string `json:"url"`
	Path        string `json:"path"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}

func FormatBranch(store Branch) BranchFormatter {
	formatter := BranchFormatter{
		ID:          store.ID,
		Name:        store.Name,
		Description: store.Description,
		Address:     store.Address,
		Logo:        store.Logo,
		Url:         os.Getenv("APP_URL"),
		Path:        store.Path,
		Phone:       store.Phone,
		Email:       store.Email,
	}
	return formatter
}

func FormatBranchs(items []Branch) []BranchFormatter {

	storeFormatter := []BranchFormatter{}

	for _, item := range items {
		formatter := FormatBranch(item)
		storeFormatter = append(storeFormatter, formatter)
	}

	return storeFormatter
}
