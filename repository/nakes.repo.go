package repository

import "triadmoko-be-golang/entity"

func (r *repository) SaveNakes(nakes entity.Nakes) (entity.Nakes, error) {
	err := r.db.Create(&nakes).Error
	if err != nil {
		return nakes, err
	}
	return nakes, nil
}
