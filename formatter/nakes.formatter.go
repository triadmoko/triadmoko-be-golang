package formatter

import (
	"triadmoko-be-golang/entity"
)

type FormatNakes struct {
	JumlahNakes int    `json:"jumlah_nakes" `
	Name        string `json:"name" `
	Addres      string `json:"address" `
	FaskesID    int    `json:"faskes_id" `
}

func FormatterNakes(nakes entity.Nakes) FormatNakes {
	formatter := FormatNakes{
		JumlahNakes: nakes.JumlahNakes,
		Name:        nakes.Name,
		Addres:      nakes.Addres,
		FaskesID:    nakes.FaskesID,
	}
	return formatter
}
func FormattNakesAll(faskes []entity.Nakes) []FormatNakes {
	var faskesAll []FormatNakes
	for _, v := range faskes {
		formatter := FormatterNakes(v)
		faskesAll = append(faskesAll, formatter)
	}
	return faskesAll
}
