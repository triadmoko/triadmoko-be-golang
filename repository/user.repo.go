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
}
type repository struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (entity.User, error) {
	var user entity.User
	err := r.db.Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *repository) Register(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *repository) FindByID(ID int) (entity.User, error) {
	var user entity.User
	err := r.db.Where("ID = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
