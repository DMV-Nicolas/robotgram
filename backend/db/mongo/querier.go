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
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (*mongo.UpdateResult, error)
	DeleteUser(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)

	CreatePost(ctx context.Context, arg CreatePostParams) (*mongo.InsertOneResult, error)
	GetPost(ctx context.Context, key string, value any) (Post, error)
	ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (*mongo.UpdateResult, error)
	DeletePost(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)

	GetLike(ctx context.Context, id primitive.ObjectID) (Like, error)
	ListLikes(ctx context.Context, arg ListLikesParams) ([]Like, error)
	CountLikes(ctx context.Context, targetID primitive.ObjectID) (int64, error)
	ToggleLike(ctx context.Context, arg ToggleLikeParams) (*mongo.InsertOneResult, *mongo.DeleteResult, error)

	CreateComment(ctx context.Context, arg CreateCommentParams) (*mongo.InsertOneResult, error)
	GetComment(ctx context.Context, id primitive.ObjectID) (Comment, error)
	ListComments(ctx context.Context, arg ListCommentsParams) ([]Comment, error)

	CreateSession(ctx context.Context, arg CreateSessionParams) (*mongo.InsertOneResult, error)
	GetSession(ctx context.Context, id primitive.ObjectID) (Session, error)
	DeleteSession(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)
	BlockSession(ctx context.Context, id primitive.ObjectID) (*mongo.UpdateResult, error)
}

var _ Querier = (*Queries)(nil)
