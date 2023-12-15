package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListUsers(ctx context.Context, limit int) ([]User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	DeleteUser(ctx context.Context, username string) error
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	GetPost(ctx context.Context, id primitive.ObjectID) (Post, error)
	ListPosts(ctx context.Context, limit int) ([]Post, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) error
	DeletePost(ctx context.Context, id primitive.ObjectID) error
}

var _ Querier = (*Queries)(nil)
