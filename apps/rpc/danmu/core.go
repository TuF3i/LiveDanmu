package danmu

import (
	"LiveDanmu/apps/rpc/danmu/kafka"
	"LiveDanmu/apps/public/config/config_template"
)

var (
	Config *
	KClient *kafka.KClient
)
