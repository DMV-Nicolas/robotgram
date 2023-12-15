package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateLikeParams struct {
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	PostID primitive.ObjectID `json:"post_id" bson:"post_id"`
}

func (q *Queries) CreateLike(ctx context.Context, arg CreateLikeParams) (*mongo.InsertOneResult, error) {
	like := Like{
		ID:        primitive.NewObjectID(),
		UserID:    arg.UserID,
		PostID:    arg.PostID,
		CreatedAt: time.Now(),
	}

	coll := q.db.Collection("likes")
	result, err := coll.InsertOne(ctx, like)

	return result, err
}

func (q *Queries) GetLike(ctx context.Context, id primitive.ObjectID) (Like, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	opts := options.FindOne()

	var like Like
	coll := q.db.Collection("likes")
	err := coll.FindOne(ctx, filter, opts).Decode(&like)

	return like, err
}

type ListLikesParams struct {
	PostID primitive.ObjectID `json:"post_id" bson:"post_id"`
	Limit  int                `json:"limit" bson:"limit"`
}

func (q *Queries) ListLikes(ctx context.Context, arg ListLikesParams) ([]Like, error) {
	filter := bson.D{primitive.E{Key: "post_id", Value: arg.PostID}}

	var likes []Like
	coll := q.db.Collection("likes")
	cursor, err := coll.Find(ctx, filter, options.Find().SetLimit(int64(arg.Limit)))
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var like Like
		err = cursor.Decode(&like)
		if err != nil {
			return nil, err
		}

		likes = append(likes, like)
	}

	return likes, nil
}

func (q *Queries) DeleteLike(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}

	coll := q.db.Collection("likes")
	result, err := coll.DeleteOne(ctx, filter)

	return result, err
}
