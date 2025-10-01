package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/IBM/sarama"
)

// Message represents the data structure for Kafka messages
type Message struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	// Get parameters
	topic := "test-topic"
	numMessages := 10

	if len(os.Args) > 1 {
		topic = os.Args[1]
	}
	if len(os.Args) > 2 {
		if n, err := strconv.Atoi(os.Args[2]); err == nil {
			numMessages = n
		}
	}

	// Kafka configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	// Create producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("❌ Failed to create producer: %v", err)
	}
	defer producer.Close()

	fmt.Printf("📤 Starting to produce %d messages to topic '%s'...\n\n", numMessages, topic)

	// Produce messages
	for i := 0; i < numMessages; i++ {
		msg := Message{
			ID:        i,
			Message:   fmt.Sprintf("Hello Kafka message #%d", i),
			Timestamp: time.Now().Unix(),
		}

		// Serialize to JSON
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			log.Printf("❌ Failed to marshal message: %v", err)
			continue
		}

		// Create Kafka message
		kafkaMsg := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(strconv.Itoa(i)),
			Value: sarama.ByteEncoder(msgBytes),
		}

		// Send message
		partition, offset, err := producer.SendMessage(kafkaMsg)
		if err != nil {
			log.Printf("❌ Failed to send message: %v", err)
		} else {
			fmt.Printf("✅ Message delivered to %s [%d] at offset %d\n", topic, partition, offset)
		}

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("\n✨ All messages sent successfully!")
}
