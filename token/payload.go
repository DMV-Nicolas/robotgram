package token

import (
	"errors"
	"time"
)

var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewPayload(username string, duration time.Duration) *Payload {
	return &Payload{
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
}

func (p Payload) Verify() error {
	if time.Now().After(p.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}
