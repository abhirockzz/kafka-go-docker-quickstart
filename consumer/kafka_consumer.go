package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var kafkaServer, kafkaTopic string

const groupID = "test-group"

func init() {
	kafkaServer = readFromENV("KAFKA_BROKER", "localhost:9092")
	kafkaTopic = readFromENV("KAFKA_TOPIC", "test")

	fmt.Println("Kafka Broker - ", kafkaServer)
	fmt.Println("Kafka topic - ", kafkaTopic)
}

func main() {
	config := kafka.ConfigMap{"bootstrap.servers": kafkaServer, "group.id": groupID, "go.events.channel.enable": true}
	consumer, consumerCreateErr := kafka.NewConsumer(&config)
	if consumerCreateErr != nil {
		fmt.Println("consumer not created ", consumerCreateErr.Error())
		os.Exit(1)
	}
	subscriptionErr := consumer.Subscribe(kafkaTopic, nil)
	if subscriptionErr != nil {
		fmt.Println("Unable to subscribe to topic " + kafkaTopic + " due to error - " + subscriptionErr.Error())
		os.Exit(1)
	} else {
		fmt.Println("subscribed to topic ", kafkaTopic)
	}

	for {
		fmt.Println("waiting for event...")
		kafkaEvent := <-consumer.Events()
		if kafkaEvent != nil {
			switch event := kafkaEvent.(type) {
			case *kafka.Message:
				fmt.Println("Message " + string(event.Value))
			case kafka.Error:
				fmt.Println("Consumer error ", event.String())
			case kafka.PartitionEOF:
				fmt.Println(kafkaEvent)
			default:
				fmt.Println(kafkaEvent)
			}
		} else {
			fmt.Println("Event was null")
		}
	}

}

func readFromENV(key, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}
