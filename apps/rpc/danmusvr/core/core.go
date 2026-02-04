package core

import (
	"LiveDanmu/apps/rpc/danmusvr/core/dao"
	"LiveDanmu/apps/rpc/danmusvr/core/kafka"
)

var (
	KClient *kafka.KClient
	Dao     *dao.Dao
)
