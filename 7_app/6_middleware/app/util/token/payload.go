package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpireToken  = errors.New("token has expire")
)

type Payload struct {
	ID         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	IssueDate  time.Time `json:"issue_date"`
	ExpireDate time.Time `json:"expire_date"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	token, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:         token,
		Username:   username,
		IssueDate:  time.Now(),
		ExpireDate: time.Now().Add(duration),
	}

	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpireDate) {
		return ErrExpireToken
	}

	return nil
}
