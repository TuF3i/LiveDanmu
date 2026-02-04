package kafka

import (
	kafka2 "LiveDanmu/apps/public/models/kafka"

	"github.com/segmentio/kafka-go"
)

func (r *ConsumerGroup) initKClient() {
	r.kClient = kafka.NewReader(kafka.ReaderConfig{
		Brokers: r.conf.KafKa.Urls,
		GroupID: r.conf.PodUID,
		Topic:   kafka2.VIDEO_DANMU_PUB_TOPIC,
	})
}
