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
	if err := q.UsernameTaken(ctx, arg.Username); err != nil {
		return nil, err
	}

	if err := q.EmailTaken(ctx, arg.Email); err != nil {
		return nil, err
	}

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

	return result, err
}

func (q *Queries) GetUser(ctx context.Context, key string, value any) (User, error) {
	filter := bson.D{primitive.E{Key: key, Value: value}}
	opts := options.FindOne()

	var user User
	coll := q.db.Collection("users")
	err := coll.FindOne(ctx, filter, opts).Decode(&user)

	return user, err
}

type ListUsersParams struct {
	Offset int64 `json:"offset" bson:"offset"`
	Limit  int64 `json:"limit" bson:"limit"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	filter := bson.D{}
	projection := bson.D{
		primitive.E{Key: "hashed_password", Value: 0},
		primitive.E{Key: "email", Value: 0},
		primitive.E{Key: "description", Value: 0},
	}

	var users []User
	coll := q.db.Collection("users")
	cursor, err := coll.Find(ctx, filter, options.Find().
		SetSkip(arg.Offset).
		SetLimit(arg.Limit).
		SetProjection(projection))

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
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	HashedPassword string             `json:"hashed_password" bson:"hashed_password"`
	FullName       string             `json:"full_name" bson:"full_name"`
	Description    string             `json:"description" bson:"description"`
	Gender         string             `json:"gender" bson:"gender"`
	Avatar         string             `json:"avatar" bson:"avatar"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": arg.ID}
	update := bson.M{
		"$set": bson.M{
			"hashed_password": arg.HashedPassword,
			"full_name":       arg.FullName,
			"description":     arg.Description,
			"gender":          arg.Gender,
			"avatar":          arg.Avatar,
		},
	}

	coll := q.db.Collection("users")
	result, err := coll.UpdateOne(ctx, filter, update)

	return result, err
}

func (q *Queries) DeleteUser(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}

	coll := q.db.Collection("users")
	result, err := coll.DeleteOne(ctx, filter)

	return result, err
}
