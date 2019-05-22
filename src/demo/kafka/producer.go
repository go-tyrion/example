package kafka

import "lib/driver/kafka"

var ProducerInstance *kafka.Producer

func init() {
	ProducerInstance = kafka.NewKakfaProducer("kafka.ini")
}
