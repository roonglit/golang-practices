package main

import (
	"asynq-demo/tasks"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		&asynq.SchedulerOpts{Location: time.Local},
	)

	// run every minute
	if _, err := scheduler.Register("* * * * *", asynq.NewTask(tasks.TypeEmailDelivery, nil)); err != nil {
		log.Fatal(err)
	}

	if _, err := scheduler.Register("@every 2s", asynq.NewTask(tasks.TypeImageProcessing, nil)); err != nil {
		log.Fatal(err)
	}

	// Run blocks and waits for os signal to terminate the program.
	if err := scheduler.Run(); err != nil {
		log.Fatal(err)
	}
}
