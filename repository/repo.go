package repository

import (
	"triadmoko-be-golang/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (entity.User, error)
	FindByID(ID int) (entity.User, error)
	Register(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)

	SaveFaskes(faskes entity.Faskes) (entity.Faskes, error)
	FindAllFaskes() ([]entity.Faskes, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}
