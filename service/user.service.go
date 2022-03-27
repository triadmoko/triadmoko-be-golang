package service

import (
	"errors"
	"time"
	"triadmoko-be-golang/entity"
	"triadmoko-be-golang/formatter"
	"triadmoko-be-golang/mapping"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) InputRegister(input formatter.FormatUser) (entity.User, error) {
	user := entity.User{}

	user.Email = input.Email
	user.Firstname = input.FirstName
	user.Lastname = input.Lastname
	user.Username = input.Username
	user.Address = input.Address
	user.Phone = input.Phone
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

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

func (s *service) GetUserByID(ID int) (entity.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (s *service) Login(input mapping.LoginInput) (entity.User, error) {

	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("password is wrong")
	}
	return user, nil
}
