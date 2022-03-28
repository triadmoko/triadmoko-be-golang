package repository

import (
	"triadmoko-be-golang/entity"
)

func (r *repository) SaveNakes(nakes entity.Nakes) (entity.Nakes, error) {
	err := r.db.Create(&nakes).Error
	if err != nil {
		return nakes, err
	}
	return nakes, nil
}
func (r *repository) UpdateNakes(nakes entity.Nakes) (entity.Nakes, error) {
	err := r.db.Save(&nakes).Error
	if err != nil {
		return nakes, err
	}
	return nakes, nil
}
func (r *repository) FindIDNakes(ID int) (entity.Nakes, error) {
	var nakes entity.Nakes
	err := r.db.Where("ID = ?", ID).Find(&nakes).Error
	if err != nil {
		return nakes, err
	}
	return nakes, nil
}
func (r *repository) FindAllNakes() ([]entity.Nakes, error) {
	var nakes []entity.Nakes
	err := r.db.Find(&nakes).Error
	if err != nil {
		return nakes, err
	}
	return nakes, nil
}
