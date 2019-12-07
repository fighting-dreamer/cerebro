package config

func GetKafkaWriterConfig() KafkaWriterConfig {
	return KafkaWriterConfig{
		Brokers: []string{"localhost:9097", "localhost:9098", "localhost:9099"},
	}
}
