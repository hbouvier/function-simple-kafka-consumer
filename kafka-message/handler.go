package function

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
	"io/ioutil"
)

func getSecretFile(secretFile string) (secretBytes []byte, err error) {
	// read from the openfaas secrets folder
	secretBytes, err = ioutil.ReadFile("/var/openfaas/secrets/" + secretFile)
	if err != nil {
		// read from the original location for backwards compatibility with openfaas <= 0.8.2
		secretBytes, err = ioutil.ReadFile("/run/secrets/" + secretFile)
	}

	return secretBytes, err
}

func getSecretFileAsString(secretFile string) (string, error) {
	secret, err := getSecretFile(secretFile)
	if err != nil {
		return "", err
	}
	return string(secret), nil
}

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
	topic, topicErr := getSecretFileAsString("kafka-response-topic")
	if topicErr != nil {
		fmt.Println(topicErr)
		return "error reading response topic secretr"
	}

	kafkaURL, URLErr := getSecretFileAsString("kafka-url")
	if URLErr != nil {
		fmt.Println(URLErr)
		return "error reading url secret"
	}
	fmt.Sprintf("using kafka url: %s", kafkaURL)
	fmt.Sprintf("using kafka topic: %s", topic)

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
