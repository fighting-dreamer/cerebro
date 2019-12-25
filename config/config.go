package config

func GetKafkaWriterConfig() KafkaWriterConfig {
	return KafkaWriterConfig{
		Brokers: []string{"localhost:9087", "localhost:9088", "localhost:9089"},
	}
}
