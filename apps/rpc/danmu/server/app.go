package main

import (
	"LiveDanmu/apps/rpc/danmu/core/handle"
	"LiveDanmu/apps/rpc/danmu/core/middleware"
	danmusvr "LiveDanmu/apps/rpc/danmu/kitex_gen/danmusvr/danmusvr"
	"log"

	"github.com/cloudwego/kitex/server"
)

func main() {
	svr := danmusvr.NewServer(new(handle.DanmuSvrImpl), server.WithMiddleware(middleware.DanmuPoolReleaseMiddleware))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
