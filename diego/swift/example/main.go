package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	fmt.Println("all is well")
}

func ProductKafka() {
	// Set up configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Create a new synchronous producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln("Failed to close Sarama producer:", err)
		}
	}()

	// Produce a message
	msg := &sarama.ProducerMessage{
		Topic: "test_topic",
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalln("Failed to send message:", err)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}

func ComsumerKafka() {
	// Set up configuration
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a new consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama consumer:", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln("Failed to close Sarama consumer:", err)
		}
	}()

	// Consume messages from the specified topic and partition
	partitionConsumer, err := consumer.ConsumePartition("test_topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln("Failed to start partition consumer:", err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln("Failed to close partition consumer:", err)
		}
	}()

	// Consume messages
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Message received: key=%s value=%s\n", string(msg.Key), string(msg.Value))
	}
}
