package kafka

import (
	kafkaCfg "LiveDanmu/apps/public/models/kafka"
	"time"

	"github.com/segmentio/kafka-go"
)

func (r *KClient) InitKafkaClient() {
	// 连接拨号器
	dialer := &kafka.Dialer{
		ClientID:  "", // TODO
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	// 热弹幕生产者
	r.HotDanmuWriter = kafka.NewWriter(
		kafka.WriterConfig{
			Brokers:          []string{}, // TODO
			Topic:            kafkaCfg.HOT_DANMU_PUB_TOPIC,
			Balancer:         &kafka.LeastBytes{},
			RequiredAcks:     1,
			Async:            false,           // 同步写入
			BatchSize:        100,             // 批量大小
			BatchTimeout:     1 * time.Second, // 批量超时时间
			Dialer:           dialer,
			CompressionCodec: kafka.Snappy.Codec(),
		},
	)

	// 正常弹幕生产者
	r.NormalDanmuWriter = kafka.NewWriter(
		kafka.WriterConfig{
			Brokers:          []string{}, // TODO
			Topic:            kafkaCfg.NORMAL_DANMU_PUB_TOPIC,
			Balancer:         &kafka.LeastBytes{},
			RequiredAcks:     1,
			Async:            false,           // 同步写入
			BatchSize:        100,             // 批量大小
			BatchTimeout:     1 * time.Second, // 批量超时时间
			Dialer:           dialer,
			CompressionCodec: kafka.Snappy.Codec(),
		},
	)

	r.VideoDanmuWriter = kafka.NewWriter(
		kafka.WriterConfig{
			Brokers:          []string{}, // TODO
			Topic:            kafkaCfg.VIDEO_DANMU_PUB_TOPIC,
			Balancer:         &kafka.LeastBytes{},
			RequiredAcks:     1,
			Async:            false,           // 同步写入
			BatchSize:        100,             // 批量大小
			BatchTimeout:     1 * time.Second, // 批量超时时间
			Dialer:           dialer,
			CompressionCodec: kafka.Snappy.Codec(),
		},
	)
}
