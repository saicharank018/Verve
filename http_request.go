package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func sendHTTPRequest(endpoint string, count int) {

	payload := map[string]int{"unique_request_count": count}
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to serialize payload: %v\n", err)
		return
	}

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Printf("HTTP request failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("HTTP request to %s returned status: %s\n", endpoint, resp.Status)
}
