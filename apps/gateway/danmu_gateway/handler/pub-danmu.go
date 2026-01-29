package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func PublishDanMuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 调用PubSvr

	}
}
