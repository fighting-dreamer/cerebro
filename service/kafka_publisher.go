package service

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
	"kilvish.io/cerebro/config"
)

type KafkaPublisher struct {
	// KWriter kafka.Writer
}

func (kp *KafkaPublisher) Publish(topic string, requestStr string) error {
	brokers := config.GetKafkaWriterConfig().Brokers
	// balancer := BalancerMapping[balancerType]

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.Hash{},
	})

	w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
	)

	w.Close()
	return nil
}
