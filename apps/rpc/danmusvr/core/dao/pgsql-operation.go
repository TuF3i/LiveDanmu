package dao

import (
	publicDao "LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/danmusvr/core/dto"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *Dao) getFullDanmuP(ctx context.Context, vid int64) ([]publicDao.DanmuData, dto.Response) {
	var results []publicDao.DanmuData
	err := r.pgdb.Where("rv_id = ?", vid).Order("ts DESC").Find(&results).Error
	if err != nil {
		return []publicDao.DanmuData{}, dto.ServerInternalError(err)
	}

	return results, dto.OperationSuccess
}

func (r *Dao) checkIfDanmuExistOnPgSQL(Tx *gorm.DB, data *danmusvr.DanmuMsg) (bool, dto.Response) {
	var dest publicDao.DanmuData
	if err := Tx.
		Where("room_id = ?", data.RoomId).
		Where("user_id = ?", data.UserId).
		Where("ts = ?", data.Ts).
		First(&dest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, dto.OperationSuccess
		}
		return false, dto.ServerInternalError(err)
	}
	return true, dto.OperationSuccess
}

func (r *Dao) delVideoDanmu(Tx *gorm.DB, data *danmusvr.DanmuMsg) dto.Response {
	var dest publicDao.DanmuData
	if err := Tx.
		Where("room_id = ?", data.RoomId).
		Where("user_id = ?", data.UserId).
		Where("ts = ?", data.Ts).
		Delete(&dest).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.OperationSuccess
		}
		return dto.ServerInternalError(err)
	}
	return dto.OperationSuccess
}
