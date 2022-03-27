package service

import (
	"triadmoko-be-golang/entity"
	"triadmoko-be-golang/mapping"
)

func (s *service) InputNakes(input mapping.InputNakes) (entity.Nakes, error) {
	nakes := entity.Nakes{}

	nakes.FaskesID = input.FaskesID
	nakes.NoNakes = input.NoNakes
	nakes.Addres = input.Addres
	nakes.Name = input.Name

	newNakes, err := s.repository.SaveNakes(nakes)

	if err != nil {
		return nakes, err
	}
	return newNakes, nil
}
