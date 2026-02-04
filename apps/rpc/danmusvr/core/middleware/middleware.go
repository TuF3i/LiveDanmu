package middleware

import (
	"LiveDanmu/apps/rpc/danmusvr/core/handle"
	"LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr"
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func DanmuPoolReleaseMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		// 执行后续业务逻辑
		_ = next(ctx, req, resp)
		// Top型类型断言
		if dmResp, ok := resp.(*danmusvr.GetTopResp); ok {
			// 释放内存
			if dmResp.Data != nil {
				handle.ReleaseDanmuMsg(dmResp.Data)
			}
		}
		// Full型类型断言
		if dmResp, ok := resp.(*danmusvr.GetResp); ok {
			// 释放内存
			if dmResp.Data != nil {
				handle.ReleaseDanmuMsg(dmResp.Data)
			}
		}
		// 其他直接return nil
		return nil
	}
}
