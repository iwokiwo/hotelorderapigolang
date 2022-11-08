package unit

type UnitFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func FormatCategory(unit Unit) UnitFormatter {
	formatter := UnitFormatter{
		ID:   unit.ID,
		Name: unit.Name,
		Slug: unit.Slug,
	}

	return formatter
}

func FormatCategories(unit []Unit) []UnitFormatter {
	UnitFormatter := []UnitFormatter{}
	for _, unit := range unit {
		formatter := FormatCategory(unit)
		UnitFormatter = append(UnitFormatter, formatter)
	}
	return UnitFormatter
}
