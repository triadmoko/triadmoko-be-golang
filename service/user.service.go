package service

import (
	"errors"
	"triadmoko-be-golang/entity"
	"triadmoko-be-golang/mapping"
	"triadmoko-be-golang/repository"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetUser() (entity.User, error)
	InputRegister(user mapping.FormatUser) (entity.User, error)
}
type service struct {
	repository repository.Repository
}

func NewServiceUser(repository repository.Repository) *service {
	return &service{repository}
}
func (s *service) GetUser() (entity.User, error) {
	user, err := s.repository.FindAll()
	if err != nil {
		return user, err
	}
	return user, nil
}
func (s *service) InputRegister(input mapping.FormatUser) (entity.User, error) {
	user := entity.User{}

	user.Email = input.Email
	user.Firstname = input.FirstName
	user.Lastname = input.Lastname
	user.Username = input.Username
	user.Address = input.Address
	user.Phone = input.Phone

	userEmail, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return userEmail, err
	}
	if input.Email == userEmail.Email {
		return user, errors.New("email available")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.Role = 1

	newUser, err := s.repository.Register(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
