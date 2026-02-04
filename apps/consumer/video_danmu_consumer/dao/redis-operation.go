package dao

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/utils"
	"context"

	jsoniter "github.com/json-iterator/go"
)

func (r *Dao) checkIfDanmuExistOnRedis(ctx context.Context, data *dao.DanmuData) (bool, error) {
	countHot, err := r.rdb.Exists(ctx, utils.GenHotDanmuKey(data.RVID)).Result()
	if err != nil {
		return false, err
	}
	countFull, err := r.rdb.Exists(ctx, utils.GenFullDanmuKey(data.RVID)).Result()
	if err != nil {
		return false, err
	}

	if countFull > 0 || countHot > 0 {
		return true, nil
	}

	return false, nil
}

func (r *Dao) insertDanmuToRedis(ctx context.Context, data *dao.DanmuData) error {
	keyForHotDanmu := utils.GenHotDanmuKey(data.RVID)
	keyForFullDanmu := utils.GenFullDanmuKey(data.RVID)

	danmuBytes, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(*data)
	if err != nil {
		return err
	}

	pipe := r.rdb.TxPipeline()
	pipe.LPush(ctx, keyForHotDanmu, danmuBytes)
	pipe.LPush(ctx, keyForFullDanmu, danmuBytes)
	_, err = pipe.Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
