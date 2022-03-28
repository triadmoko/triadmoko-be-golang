package service

import (
	"triadmoko-be-golang/entity"
	"triadmoko-be-golang/formatter"
	"triadmoko-be-golang/mapping"
	"triadmoko-be-golang/repository"
)

type Service interface {
	InputRegister(input formatter.FormatUser) (entity.User, error)
	GetUserByID(ID int) (entity.User, error)
	Login(input mapping.LoginInput) (entity.User, error)

	InputFaskes(input mapping.InputFaskes) (entity.Faskes, error)
	FindAllFaskes() ([]entity.Faskes, error)

	InputNakes(input mapping.InputNakes) (entity.Nakes, error)
	UpdateNakes(ID int, input mapping.UpdateNakes) (entity.Nakes, error)
	FindIDNakes(ID int) (entity.Nakes, error)
	FindAllNakes() ([]entity.Nakes, error)
	DeleteNakes(ID int) (entity.Nakes, error)
}
type service struct {
	repository repository.Repository
}

func NewServiceUser(repository repository.Repository) *service {
	return &service{repository}
}
