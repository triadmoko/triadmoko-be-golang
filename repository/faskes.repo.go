package repository

import "triadmoko-be-golang/entity"

func (r *repository) SaveFaskes(faskes entity.Faskes) (entity.Faskes, error) {
	err := r.db.Create(&faskes).Error
	if err != nil {
		return faskes, err
	}
	return faskes, nil
}
func (r *repository) FindAllFaskes() ([]entity.Faskes, error) {
	var faskes []entity.Faskes
	err := r.db.Find(&faskes).Error
	if err != nil {
		return faskes, err
	}
	return faskes, nil
}
