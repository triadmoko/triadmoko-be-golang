package service

import (
	"triadmoko-be-golang/entity"
	"triadmoko-be-golang/mapping"
)

func (s *service) InputFaskes(input mapping.InputFaskes) (entity.Faskes, error) {
	faskes := entity.Faskes{}
	faskes.Name = input.Name

	newFaskes, err := s.repository.SaveFaskes(faskes)
	if err != nil {
		return newFaskes, err
	}
	return newFaskes, nil
}
