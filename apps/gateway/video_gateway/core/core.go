package core

import (
	"LiveDanmu/apps/gateway/video_gateway/core/dao"
	"LiveDanmu/apps/public/logger"

	"github.com/bwmarrin/snowflake"
)

var (
	Logger    *logger.NewLogger
	SnowFlake *snowflake.Node // 雪花
	Dao       *dao.Dao
)
