package formatter

import "triadmoko-be-golang/entity"

type FormatNakes struct {
	NoNakes  int    `json:"no_nakes" `
	Name     string `json:"name" `
	Addres   string `json:"address" `
	FaskesID int    `json:"faskes_id" `
}

func FormatterNakes(nakes entity.Nakes) FormatNakes {
	formatter := FormatNakes{
		NoNakes:  nakes.NoNakes,
		Name:     nakes.Name,
		Addres:   nakes.Addres,
		FaskesID: nakes.FaskesID,
	}
	return formatter
}
