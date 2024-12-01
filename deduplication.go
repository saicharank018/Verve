package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func addUniqueID(id int) bool {
	ctx := context.Background()
	key := fmt.Sprintf("unique_id:%d", id)

	success, err := redisClient.SetNX(ctx, key, true, time.Minute).Result()
	if err != nil {
		log.Printf("Redis error: %v", err)
		return false
	}
	return success
}

func getUniqueCount() int {
	ctx := context.Background()
	keys, err := redisClient.Keys(ctx, "unique_id:*").Result()
	if err != nil {
		log.Printf("Failed to fetch unique keys: %v", err)
		return 0
	}
	return len(keys)
}

func clearUniqueIDs() {
	ctx := context.Background()
	keys, err := redisClient.Keys(ctx, "unique_id:*").Result()
	if err != nil {
		log.Printf("Failed to fetch unique keys for clearing: %v", err)
		return
	}
	for _, key := range keys {
		if err := redisClient.Del(ctx, key).Err(); err != nil {
			log.Printf("Failed to delete key %s: %v", key, err)
		}
	}
}
