package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUsernameTaken  = errors.New("the username must be unique")
	ErrEmailTaken     = errors.New("the email must be unique")
	ErrDuplicatedLike = errors.New("the like has already been given")
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
