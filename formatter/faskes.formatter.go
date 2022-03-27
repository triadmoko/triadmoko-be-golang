package formatter

import "triadmoko-be-golang/entity"

type FormatFaskes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatterFaskes(faskes entity.Faskes) FormatFaskes {
	response := FormatFaskes{
		ID:   faskes.ID,
		Name: faskes.Name,
	}
	return response
}
