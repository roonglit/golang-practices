package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

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
	var p EmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v", err)
	}

	// Simulate email sending
	log.Printf("Sending email to %s (User: %d)", p.Email, p.UserID)
	time.Sleep(2 * time.Second)
	log.Printf("Email sent successfully to %s", p.Email)
	return nil
}
