package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateUserParams struct {
	Username       string `json:"username" bson:"username"`
	HashedPassword string `json:"hashed_password" bson:"hashed_password"`
	FullName       string `json:"full_name" bson:"full_name"`
	Email          string `json:"email" bson:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (any, error) {
	coll := q.db.Collection("users")

	result, err := coll.InsertOne(ctx, arg)
	id := result.InsertedID
	return id, err
}

type GetUserParams struct {
	Username string `json:"username" bson:"username"`
}

func (q *Queries) GetUser(ctx context.Context, arg GetUserParams) (User, error) {
	coll := q.db.Collection("users")
	filter := bson.D{primitive.E{Key: "username", Value: arg.Username}}
	opts := options.FindOne()

	var user User
	err := coll.FindOne(ctx, filter, opts).Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
