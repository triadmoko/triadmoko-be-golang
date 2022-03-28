package entity

import "gorm.io/gorm"

type Nakes struct {
	gorm.Model
	ID          int
	JumlahNakes int
	Name        string
	Addres      string
	FaskesID    int
}
