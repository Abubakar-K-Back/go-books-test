package kafka

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// ProduceMessage sends a message to the Kafka topic
func ProduceMessage(topic string, message string) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
	})
	if err != nil {
		log.Fatal("Failed to create Kafka producer:", err)
	}
	defer producer.Close()

	// Produce the message
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		log.Println("Error producing message:", err)
	} else {
		log.Println("Produced Kafka message to topic:", topic)
	}

	// Flush messages
	producer.Flush(15000)
}
