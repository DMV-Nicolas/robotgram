package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Username       string             `json:"username" bson:"username"`
	HashedPassword string             `json:"hashed_password" bson:"hashed_password"`
	FullName       string             `json:"full_name" bson:"full_name"`
	Email          string             `json:"email" bson:"email"`
	Avatar         string             `json:"avatar" bson:"avatar"`
	Description    string             `json:"description" bson:"description"`
	Gender         string             `json:"gender" bson:"gender"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
}
