package main

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var kafkaServer, kafkaTopic string

func init() {
	kafkaServer = readFromENV("KAFKA_BROKER", "localhost:9092")
	kafkaTopic = readFromENV("KAFKA_TOPIC", "test")

	fmt.Println("Kafka Broker - ", kafkaServer)
	fmt.Println("Kafka topic - ", kafkaTopic)
}
func main() {
	producer, producerCreateErr := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})

	if producerCreateErr != nil {
		fmt.Println("Failed to create producer due to ", producerCreateErr)
		os.Exit(1)
	}
	for {
		for i := 0; i < 5; i++ {
			value := time.Now().String()
			producerErr := producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
				Value:          []byte(value),
			}, nil)

			if producerErr != nil {
				fmt.Println("unable to enqueue message ", value)
			}
			event := <-producer.Events()

			message := event.(*kafka.Message)

			if message.TopicPartition.Error != nil {
				fmt.Println("Delivery failed due to error ", message.TopicPartition.Error)
			} else {
				fmt.Println("Delivered message to offset " + message.TopicPartition.Offset.String() + " in partition " + message.TopicPartition.String())
			}
		}

		time.Sleep(time.Second * 5) //wait for 5 seconds before sending another batch of messages
	}

}
func readFromENV(key, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}
