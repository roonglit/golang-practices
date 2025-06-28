package main

import (
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

type EmailPayload struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	defer client.Close()

	// Immediate task
	payload, _ := json.Marshal(EmailPayload{
		UserID: 123,
		Email:  "user@example.com",
	})

	task := asynq.NewTask("email:welcome", payload)
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("Enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
