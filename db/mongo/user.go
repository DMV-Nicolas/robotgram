package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateUserParams struct {
	Username       string `json:"username" bson:"username"`
	HashedPassword string `json:"hashed_password" bson:"hashed_password"`
	FullName       string `json:"full_name" bson:"full_name"`
	Email          string `json:"email" bson:"email"`
	Avatar         string `json:"avatar" bson:"avatar"`
	Gender         string `json:"gender" bson:"gender"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (*mongo.InsertOneResult, error) {
	user := User{
		ID:             primitive.NewObjectID(),
		Username:       arg.Username,
		HashedPassword: arg.HashedPassword,
		FullName:       arg.FullName,
		Email:          arg.Email,
		Avatar:         arg.Avatar,
		Description:    "",
		Gender:         arg.Gender,
		CreatedAt:      time.Now(),
	}

	coll := q.db.Collection("users")
	result, err := coll.InsertOne(ctx, user)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	opts := options.FindOne()

	var user User
	coll := q.db.Collection("users")
	err := coll.FindOne(ctx, filter, opts).Decode(&user)

	return user, err
}

func (q *Queries) ListUsers(ctx context.Context, limit int) ([]User, error) {
	filter := bson.D{}

	var users []User
	coll := q.db.Collection("users")
	cursor, err := coll.Find(ctx, filter, options.Find().SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var user User
		err = cursor.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

type UpdateUserParams struct {
	Username       string `json:"username" bson:"username"`
	HashedPassword string `json:"hashed_password" bson:"hashed_password"`
	FullName       string `json:"full_name" bson:"full_name"`
	Description    string `json:"description" bson:"description"`
	Gender         string `json:"gender" bson:"gender"`
	Avatar         string `json:"avatar" bson:"avatar"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (*mongo.UpdateResult, error) {
	filter := bson.M{"username": arg.Username}
	update := bson.M{
		"$set": bson.M{
			"hashed_password": arg.HashedPassword,
			"full_name":       arg.FullName,
			"description":     arg.Description,
			"gender":          arg.Gender,
			"avatar":          arg.Avatar,
			"a":               "a",
		},
	}

	coll := q.db.Collection("users")
	result, err := coll.UpdateOne(ctx, filter, update)

	return result, err
}

func (q *Queries) DeleteUser(ctx context.Context, username string) (*mongo.DeleteResult, error) {
	filter := bson.M{"username": username}

	coll := q.db.Collection("users")
	result, err := coll.DeleteOne(ctx, filter)

	return result, err
}
