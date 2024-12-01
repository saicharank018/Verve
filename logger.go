package main

import (
	"log"
	"time"
)

func startLogging() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		count := getUniqueCount()
		log.Printf("Unique requests in the last minute: %d\n", count)

		if err := sendToStreamingService(count); err != nil {
			log.Printf("Failed to send to streaming service: %v\n", err)
		}

		clearUniqueIDs()
	}
}
