package entity

import "gorm.io/gorm"

type Faskes struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
}
