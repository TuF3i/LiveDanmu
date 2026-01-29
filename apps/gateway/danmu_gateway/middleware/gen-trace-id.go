package middleware

import (
	"LiveDanmu/apps/gateway/danmu_gateway"
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/hertz/pkg/app"
)

func GenTraceID() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 生成TraceID
		traceID := danmu_gateway.SnowFlake.Generate().String()
		ctx = metainfo.WithValue(ctx, danmu_gateway.TRACE_ID_KEY, traceID)
		c.Next(ctx)
	}
}
