package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

// Message represents the data structure for Kafka messages
type Message struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

// ConsumerGroupHandler implements sarama.ConsumerGroupHandler
type ConsumerGroupHandler struct {
	messageCount int
}

func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		h.messageCount++

		var msg Message
		if err := json.Unmarshal(message.Value, &msg); err != nil {
			log.Printf("❌ Error unmarshaling message: %v", err)
			continue
		}

		fmt.Printf("📨 Message #%d\n", h.messageCount)
		fmt.Printf("   Topic: %s\n", message.Topic)
		fmt.Printf("   Partition: %d\n", message.Partition)
		fmt.Printf("   Offset: %d\n", message.Offset)
		fmt.Printf("   Key: %s\n", string(message.Key))
		fmt.Printf("   Value: %+v\n\n", msg)

		session.MarkMessage(message, "")
	}

	return nil
}

func main() {
	// Get parameters
	topic := "test-topic"
	if len(os.Args) > 1 {
		topic = os.Args[1]
	}

	// Kafka configuration
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Create consumer group
	consumerGroup, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "go-consumer-group", config)
	if err != nil {
		log.Fatalf("❌ Failed to create consumer group: %v", err)
	}
	defer consumerGroup.Close()

	fmt.Printf("📥 Starting to consume messages from topic '%s'...\n", topic)
	fmt.Println("Press Ctrl+C to stop\n")

	// Handle signals
	ctx, cancel := context.WithCancel(context.Background())
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	handler := &ConsumerGroupHandler{}

	// Consumer loop
	go func() {
		for {
			if err := consumerGroup.Consume(ctx, []string{topic}, handler); err != nil {
				log.Printf("❌ Error from consumer: %v", err)
			}

			if ctx.Err() != nil {
				return
			}
		}
	}()

	<-sigterm
	fmt.Println("\n⚠️  Consumer interrupted by user")
	cancel()
	fmt.Println("🔒 Closing consumer...")
}
