package dao

import (
	"LiveDanmu/apps/public/models/dao"
	"errors"

	"gorm.io/gorm"
)

func (r *Dao) checkIfDanmuExistOnPgSQL(Tx *gorm.DB, data *dao.DanmuData) (bool, error) {
	var dest dao.DanmuData
	if err := Tx.
		Where("room_id = ?", data.RVID).
		Where("user_id = ?", data.UserId).
		Where("ts = ?", data.Ts).
		First(&dest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *Dao) insertDanmuToPgSQL(Tx *gorm.DB, data *dao.DanmuData) error {
	if err := Tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}
