package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Querier is an interface that contains all the methods of the database
type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (*mongo.InsertOneResult, error)
	GetUser(ctx context.Context, key string, value any) (User, error)
	ListUsers(ctx context.Context, limit int) ([]User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (*mongo.UpdateResult, error)
	DeleteUser(ctx context.Context, username string) (*mongo.DeleteResult, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (*mongo.InsertOneResult, error)
	GetPost(ctx context.Context, key string, value any) (Post, error)
	ListPosts(ctx context.Context, limit int) ([]Post, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (*mongo.UpdateResult, error)
	DeletePost(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)
	CreateLike(ctx context.Context, arg CreateLikeParams) (*mongo.InsertOneResult, error)
	GetLike(ctx context.Context, id primitive.ObjectID) (Like, error)
	ListLikes(ctx context.Context, arg ListLikesParams) ([]Like, error)
	DeleteLike(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)

	UsernameTaken(ctx context.Context, username string) error
	EmailTaken(ctx context.Context, email string) error
}

var _ Querier = (*Queries)(nil)
