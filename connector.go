package main

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/redis/go-redis/v9"
	"log"
)

var redisClient *redis.Client
var kafkaProducer sarama.AsyncProducer

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Default DB
	})

	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")
}

func initKafka() {

	var err error
	kafkaProducer, err = sarama.NewAsyncProducer([]string{"localhost:9092"}, sarama.NewConfig())
	if err != nil {
		log.Fatalf("Failed to connect to Kafka: %v", err)
		return
	}
}
