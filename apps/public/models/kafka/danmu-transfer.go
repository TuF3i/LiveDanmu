package kafka

import "LiveDanmu/apps/public/models/dao"

const (
	LIVE_DANMU_PUB_TOPIC       = "danmusvr.live"
	LIVE_DANMU_BOARDCAST_TOPIC = "danmusvr.live.boardcast"
	VIDEO_DANMU_PUB_TOPIC      = "danmusvr.video"
)

type DanmuKMsg struct {
	RVID    int64
	TraceID string
	Data    dao.DanmuData
}
