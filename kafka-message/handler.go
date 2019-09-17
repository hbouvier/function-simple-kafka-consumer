package function

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

// Handle a serverless request
func Handle(req []byte) string {
	// to produce messages
	topic := "response"
	kafkaURL := "kafka.openfaas.9092"

	writer := newKafkaWriter(kafkaURL, topic)
	defer writer.Close()

	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("Key: %s", uuid.New())),
		Value: []byte(fmt.Sprintf("Body: %s", string(req))),
	}
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("Hello, Go. You said: %s", string(req))
}
