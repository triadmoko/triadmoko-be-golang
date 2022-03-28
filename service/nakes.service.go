package service

import (
	"errors"
	"time"
	"triadmoko-be-golang/entity"
	"triadmoko-be-golang/mapping"
)

func (s *service) InputNakes(input mapping.InputNakes) (entity.Nakes, error) {
	nakes := entity.Nakes{}

	nakes.FaskesID = input.FaskesID
	nakes.NoNakes = input.NoNakes
	nakes.Addres = input.Addres
	nakes.Name = input.Name
	nakes.CreateAt = time.Now()

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
	nakes.NoNakes = input.NoNakes
	nakes.Addres = input.Addres
	nakes.Name = input.Name
	nakes.UpdateAt = time.Now()

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
