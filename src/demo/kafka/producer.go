package kafka

import "lib/driver/kafka"

var ProducerInstance *kafka.Producer

func init() {
	ProducerInstance = kafka.NewKafkaProducer("kafka.ini")
}
