package service

import (
	"context"
	"math/rand"
	"strconv"

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
	uniqueID := strconv.Itoa(rand.Intn(1000000))
	w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(uniqueID),
			Value: []byte(requestStr),
		},
	)

	w.Close()
	return nil
}
