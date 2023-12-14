package db

import (
	"time"
)

type User struct {
	Username          string    `json:"username" bson:"username"`
	HashedPassword    string    `json:"hashed_password" bson:"hashed_password"`
	FullName          string    `json:"full_name" bson:"full_name"`
	Email             string    `json:"email" bson:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at" bson:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at" bson:"created_at"`
}
