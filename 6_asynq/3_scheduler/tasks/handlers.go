package tasks

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
)

const (
	TypeEmailDelivery   = "email:deliver"
	TypeImageProcessing = "image:process"
)

type EmailPayload struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}

func HandleEmailDelivery(ctx context.Context, t *asynq.Task) error {
	log.Printf("Do email delivery")
	return nil
}

func HandleImageProcessing(ctx context.Context, t *asynq.Task) error {
	log.Printf("Do image processing")
	return nil
}
