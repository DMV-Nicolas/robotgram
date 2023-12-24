package token

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Maker is an interface for managing tokens.
type Maker interface {
	// CreateToken creates a new token for the specific username and duration
	CreateToken(userID primitive.ObjectID, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
