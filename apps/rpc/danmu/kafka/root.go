package kafka

import "github.com/segmentio/kafka-go"

type KClient struct {
	HotDanmuWriter    *kafka.Writer
	NormalDanmuWriter *kafka.Writer
	VideoDanmuWriter  *kafka.Writer
}

func InitKClient() *KClient {}
