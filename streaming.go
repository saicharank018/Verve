package main

import (
	"github.com/IBM/sarama"
	"log"
)

func sendToStreamingService(count int) error {

	msg := &sarama.ProducerMessage{Topic: "unique-request-count"}
	msg.Key = sarama.StringEncoder(count)

	kafkaProducer.Input() <- msg

	log.Println("Successfully sent message to Kafka")
	return nil
}
