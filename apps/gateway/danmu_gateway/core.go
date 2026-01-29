package danmu_gateway

import (
	"LiveDanmu/apps/gateway/danmu_gateway/models"
	"LiveDanmu/apps/rpc/danmu/kitex_gen/pub/pub"

	"github.com/bwmarrin/snowflake"
	hertzzap "github.com/hertz-contrib/logger/zap"
)

const (
	TRACE_ID_KEY = "trace_id"
)

var (
	Logger    *hertzzap.Logger // 日志组件
	Config    *models.Config   // 配置文件
	SnowFlake *snowflake.Node  // 雪花
)

var (
	PubSvr pub.Client // 发弹幕微服务
)
