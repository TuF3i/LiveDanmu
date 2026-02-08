package main

import (
	"LiveDanmu/apps/rpc/danmusvr/core/handle"
	"LiveDanmu/apps/rpc/danmusvr/core/middleware"
	danmusvr "LiveDanmu/apps/rpc/danmusvr/kitex_gen/danmusvr/danmusvr"
	"log"

	"github.com/cloudwego/kitex/server"
)

func main() {
	svr := danmusvr.NewServer(new(handle.DanmuSvrImpl), server.WithMiddleware(middleware.PreInit), server.WithMiddleware(middleware.DanmuPoolReleaseMiddleware))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
