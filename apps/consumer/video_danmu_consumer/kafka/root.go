package kafka

import (
	"LiveDanmu/apps/public/config/config_template"

	"github.com/segmentio/kafka-go"
)

type ConsumerGroup struct {
	conf    *config_template.LiveDanmuConsumerConfig
	kClient *kafka.Reader
}

func GetConsumerGroup(conf *config_template.LiveDanmuConsumerConfig) *ConsumerGroup {
	c := ConsumerGroup{conf: conf}
	c.initKClient()
	return &c
}
