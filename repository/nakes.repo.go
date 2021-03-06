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
	err := r.db.Where("ID = ?", ID).Preload("Faskes").Find(&nakes).Error
	if err != nil {
		return nakes, err
	}
	return nakes, nil
}
func (r *repository) FindAllNakes() ([]entity.Nakes, error) {
	var nakes []entity.Nakes
	err := r.db.Preload("Faskes").Find(&nakes).Error
	if err != nil {
		return nakes, err
	}
	return nakes, nil
}
func (r *repository) DeleteNakes(ID int) (entity.Nakes, error) {
	nakes := entity.Nakes{}
	err := r.db.Delete(&nakes, ID).Error
	if err != nil {
		return nakes, err
	}
	return nakes, nil
}
