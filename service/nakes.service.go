package service

import (
	"errors"
	"triadmoko-be-golang/entity"
	"triadmoko-be-golang/mapping"
)

func (s *service) InputNakes(input mapping.InputNakes) (entity.Nakes, error) {
	nakes := entity.Nakes{}

	nakes.FaskesID = input.FaskesID
	nakes.JumlahNakes = input.JumlahNakes
	nakes.Addres = input.Addres
	nakes.Name = input.Name

	newNakes, err := s.repository.SaveNakes(nakes)

	if err != nil {
		return nakes, err
	}
	return newNakes, nil
}
func (s *service) UpdateNakes(ID int, input mapping.UpdateNakes) (entity.Nakes, error) {
	nakes, err := s.repository.FindIDNakes(ID)
	if err != nil {
		return nakes, err
	}
	nakes.FaskesID = input.FaskesID
	nakes.JumlahNakes = input.JumlahNakes
	nakes.Addres = input.Addres
	nakes.Name = input.Name

	newNakes, err := s.repository.UpdateNakes(nakes)

	if err != nil {
		return nakes, err
	}
	return newNakes, nil
}
func (s *service) FindIDNakes(ID int) (entity.Nakes, error) {
	nakes, err := s.repository.FindIDNakes(ID)
	if err != nil {
		return nakes, err
	}
	if nakes.ID == 0 {
		return nakes, errors.New("nakes not found")
	}
	return nakes, nil
}
func (s *service) FindAllNakes() ([]entity.Nakes, error) {
	nakes, err := s.repository.FindAllNakes()
	if err != nil {
		return nakes, err
	}
	return nakes, err
}
func (s *service) DeleteNakes(ID int) (entity.Nakes, error) {
	nakes, err := s.repository.DeleteNakes(ID)
	if err != nil {
		return nakes, err
	}
	return nakes, err
}
