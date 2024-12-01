package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	logFile, err := os.OpenFile("requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	// Initialize Redis (for multi-instance deduplication)
	initRedis()

	// Initialize Kafka (for distributed streaming)
	initKafka()

	// Initialize Gin Router
	router := gin.Default()
	setupRoutes(router)

	// Start unique request logging
	go startLogging()

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
