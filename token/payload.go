package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("token invalid")
var ErrInvalidKeySize = fmt.Errorf("invalid key size: must be at least %d characters", minSecertKeySize)

type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	IssueAt  time.Time `json:"issue_at"`
	ExpireAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:       tokenID,
		Username: username,
		IssueAt:  time.Now(),
		ExpireAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpireAt) {
		return ErrExpiredToken
	}
	return nil
}
