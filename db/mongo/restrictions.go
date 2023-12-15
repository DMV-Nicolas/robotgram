package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrUsernameTaken  = errors.New("The username must be unique")
	ErrEmailTaken     = errors.New("The email must be unique")
	ErrDuplicatedLike = errors.New("The like has already been given")
)

func (q *Queries) UsernameTaken(ctx context.Context, username string) error {
	_, err := q.GetUser(ctx, username)
	if err != nil {
		return nil
	}
	return ErrUsernameTaken
}

func (q *Queries) EmailTaken(ctx context.Context, email string) error {
	filter := bson.D{primitive.E{Key: "email", Value: email}}
	opts := options.FindOne()

	var user User
	coll := q.db.Collection("users")
	err := coll.FindOne(ctx, filter, opts).Decode(&user)

	if err != nil {
		return nil
	}

	return ErrEmailTaken
}

func (q *Queries) DuplicatedLike(ctx context.Context, arg CreateLikeParams) error {
	filter := bson.D{
		primitive.E{Key: "user_id", Value: arg.UserID},
		primitive.E{Key: "post_id", Value: arg.PostID},
	}
	opts := options.FindOne()

	var like Like
	coll := q.db.Collection("likes")
	err := coll.FindOne(ctx, filter, opts).Decode(&like)

	if err != nil {
		return nil
	}

	return ErrDuplicatedLike
}
