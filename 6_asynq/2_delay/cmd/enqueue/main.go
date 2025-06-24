package main

import (
	"asynq-demo/tasks"
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	defer client.Close()

	// Immediate task
	payload, _ := json.Marshal(tasks.EmailPayload{
		UserID: 123,
		Email:  "user@example.com",
	})

	task := asynq.NewTask(tasks.TypeEmailDelivery, payload)

	// Scheduled task (30 seconds from now)
	info, err := client.Enqueue(task, asynq.ProcessIn(30*time.Second))
	if err != nil {
		log.Fatalf("could not schedule task: %v", err)
	}
	log.Printf("Scheduled task: id=%s queue=%s", info.ID, info.Queue)
}
