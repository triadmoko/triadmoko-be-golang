package entity

import (
	"gorm.io/gorm"
)

type Faskes struct {
	gorm.Model
	ID   int
	Name string
}
