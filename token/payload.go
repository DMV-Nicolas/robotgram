package token

import (
	"errors"
	"time"
)

var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Payload contains the payload data of the token
type Payload struct {
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) *Payload {
	return &Payload{
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
}

// Valid checks if the payload is valid or not.
func (p Payload) Valid() error {
	if time.Now().After(p.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}
