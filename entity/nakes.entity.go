package entity

import "gorm.io/gorm"

type Nakes struct {
	gorm.Model
	JumlahNakes int    `json:"jumlah_nakes"`
	Name        string `json:"name"`
	Addres      string `json:"address"`
	FaskesID    int    `json:"faskes_id"`
	Faskes      Faskes `json:"faskes"`
}
