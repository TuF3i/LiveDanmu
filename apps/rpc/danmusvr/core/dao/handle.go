package dao

import (
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/danmusvr/core/dto"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
	"context"
	"errors"
)

// Compare cy的小工具,防止切片越界
func Compare(max int, data []dao.DanmuData) []dao.DanmuData {
	length := len(data)
	if length >= max {
		return data[:max-1]
	}
	return data
}

func (r *Dao) ReadHotDanmu(ctx context.Context, vid int64) ([]dao.DanmuData, dto.Response) {
	// 从redis读数据
	data, resp := r.getHotDanmuR(ctx, vid)
	if !errors.Is(resp, dto.OperationSuccess) {
		return []dao.DanmuData{}, resp
	}
	// redis里没有就穿透到pgsql
	if len(data) == 0 {
		// 从pgsql里拉数据
		data, resp := r.getFullDanmuP(ctx, vid)
		if !errors.Is(resp, dto.OperationSuccess) {
			return []dao.DanmuData{}, resp
		}
		// 向redis里写入hotDanmu
		resp = r.setHotDanmuR(ctx, vid, Compare(1000, data))
		if !errors.Is(resp, dto.OperationSuccess) {
			return []dao.DanmuData{}, resp
		}
		// 向redis里写入fullDanmu
		resp = r.setFullDanmuR(ctx, vid, data)
		if !errors.Is(resp, dto.OperationSuccess) {
			return []dao.DanmuData{}, resp
		}

		return data, dto.OperationSuccess
	}
	// 计数器递增
	resp = r.incrementHotR(ctx, vid)
	if !errors.Is(resp, dto.OperationSuccess) {
		return []dao.DanmuData{}, resp
	}

	return data, dto.OperationSuccess
}

func (r *Dao) ReadFullDanmu(ctx context.Context, vid int64) ([]dao.DanmuData, dto.Response) {
	// 从redis拉数据
	data, resp := r.getFullDanmuR(ctx, vid)
	if !errors.Is(resp, dto.OperationSuccess) {
		return []dao.DanmuData{}, resp
	}
	// 如果mysql里没就走pgsql
	if len(data) == 0 {
		// 从pgsql里拉数据
		data, resp := r.getFullDanmuP(ctx, vid)
		if !errors.Is(resp, dto.OperationSuccess) {
			return []dao.DanmuData{}, resp
		}
		// 向redis里写入fullDanmu
		resp = r.setFullDanmuR(ctx, vid, data)
		if !errors.Is(resp, dto.OperationSuccess) {
			return []dao.DanmuData{}, resp
		}

		return data, dto.OperationSuccess
	}
	// 计数器递增
	resp = r.incrementFullR(ctx, vid)
	if !errors.Is(resp, dto.OperationSuccess) {
		return []dao.DanmuData{}, resp
	}

	return data, dto.OperationSuccess
}

func (r *Dao) DelVideoDanmu(ctx context.Context, msg *danmusvr.DanmuMsg) dto.Response {
	// 从redis中删除整个key，下次访问时自动补位
	resp := r.delDanmuInRedis(ctx, msg.RoomId)
	if !errors.Is(resp, dto.OperationSuccess) {
		return resp
	}
	// 从pgsql中删除弹幕
	tx := r.pgdb.Begin()
	ok, resp := r.checkIfDanmuExistOnPgSQL(tx, msg)
	if !errors.Is(resp, dto.OperationSuccess) {
		tx.Rollback()
		return resp
	}
	// 字段不存在直接返回
	if !ok {
		tx.Commit()
		return dto.OperationSuccess
	}
	// 删除弹幕
	resp = r.delVideoDanmu(tx, msg)
	if !errors.Is(resp, dto.OperationSuccess) {
		tx.Rollback()
		return resp
	}

	tx.Commit()
	return dto.OperationSuccess
}
