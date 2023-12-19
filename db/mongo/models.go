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

type Post struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Images      []string           `json:"images" bson:"images"`
	Videos      []string           `json:"videos" bson:"videos"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

type Like struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	TargetID  primitive.ObjectID `json:"target_id" bson:"target_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type Session struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	RefreshToken string             `json:"refresh_token" bson:"refresh_token"`
	UserAgent    string             `json:"user_agent" bson:"user_agent"`
	ClientIP     string             `json:"client_ip" bson:"client_ip"`
	IsBlocked    bool               `json:"is_blocked" bson:"is_blocked"`
	ExpiresAt    time.Time          `json:"expires_at" bson:"expires_at"`
}
