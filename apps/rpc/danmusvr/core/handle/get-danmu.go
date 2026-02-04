package handle

import (
	"LiveDanmu/apps/public/dto"
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/rpc/danmusvr/core"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
	"context"
	"errors"
	"sync"
)

// 可复用指针内存池，避免重复申请内存
var danmuMsgPool = sync.Pool{New: func() interface{} { return &danmusvr.DanmuMsg{} }}

// 生成DanmuMsg
func genGetMsg(data []dao.DanmuData) []*danmusvr.DanmuMsg {
	results := make([]*danmusvr.DanmuMsg, 0, len(data))
	for _, dm := range data {
		// 从内存池里取出一个对象
		result := danmuMsgPool.Get().(*danmusvr.DanmuMsg)
		// 覆写内存字段
		*result = danmusvr.DanmuMsg{
			RoomId:  dm.RVID,
			UserId:  dm.UserId,
			Content: dm.Content,
			Color:   dm.Color,
			Ts:      dm.Ts,
		}
		results = append(results, result)
	}

	return results
}

// ReleaseDanmuMsg 使用完释放
func ReleaseDanmuMsg(ori []*danmusvr.DanmuMsg) {
	for _, msg := range ori {
		danmuMsgPool.Put(msg)
	}
	ori = nil
}

func GetHotDanmu(ctx context.Context, req *danmusvr.GetTopReq) ([]*danmusvr.DanmuMsg, dto.Response) {
	// 从数据库读弹幕
	dm, resp := core.Dao.ReadHotDanmu(ctx, req.BV)
	if !errors.Is(resp, dto.OperationSuccess) {
		return nil, resp
	}
	// 转换结构体类型
	data := genGetMsg(dm)

	return data, dto.OperationSuccess
}

func GetFullDanmu(ctx context.Context, req *danmusvr.GetReq) ([]*danmusvr.DanmuMsg, dto.Response) {
	// 从数据库读弹幕
	dm, resp := core.Dao.ReadFullDanmu(ctx, req.BV)
	if !errors.Is(resp, dto.OperationSuccess) {
		return nil, resp
	}
	// 转换结构体类型
	data := genGetMsg(dm)

	return data, dto.OperationSuccess
}
