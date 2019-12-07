package service

type IKafkaPublisher interface {
	Publish(topic string, requestStr string) error
}
