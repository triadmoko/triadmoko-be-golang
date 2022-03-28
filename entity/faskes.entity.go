package entity

import "gorm.io/gorm"

type Faskes struct {
	gorm.Model
	Name string `json:"name"`
}
