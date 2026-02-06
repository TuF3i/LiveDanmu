package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func JWTMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求头提取AccessToken
		authHeader := string(c.GetHeader("Authorization"))
		if authHeader == "" {

		}
	}
}
