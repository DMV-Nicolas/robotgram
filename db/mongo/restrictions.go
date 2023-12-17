package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrUsernameTaken  = errors.New("The username must be unique")
	ErrEmailTaken     = errors.New("The email must be unique")
	ErrDuplicatedLike = errors.New("The like has already been given")
)

// UsernameTaken verifies in the database if the provided username is taken or not
func (q *Queries) UsernameTaken(ctx context.Context, username string) error {
	_, err := q.GetUser(ctx, "username", username)
	if err == mongo.ErrNoDocuments {
		return nil
	}

	if err != nil {
		return err
	}

	return ErrUsernameTaken
}

// EmailTaken verifies in the database if the provided email is taken or not
func (q *Queries) EmailTaken(ctx context.Context, email string) error {
	_, err := q.GetUser(ctx, "email", email)
	if err == mongo.ErrNoDocuments {
		return nil
	}

	if err != nil {
		return err
	}

	return ErrEmailTaken
}

// DuplicatedLike verfies in the database if the entered like is already registered
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
