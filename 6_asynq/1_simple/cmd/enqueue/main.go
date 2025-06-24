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
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("Enqueued task: id=%s queue=%s", info.ID, info.Queue)

	// Scheduled task (5 minutes from now)
	info, err = client.Enqueue(task, asynq.ProcessIn(5*time.Minute))
	if err != nil {
		log.Fatalf("could not schedule task: %v", err)
	}
	log.Printf("Scheduled task: id=%s queue=%s", info.ID, info.Queue)
}
