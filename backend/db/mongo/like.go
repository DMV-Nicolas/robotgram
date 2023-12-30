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
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	TargetID primitive.ObjectID `json:"target_id" bson:"target_id"`
}

func (q *Queries) CreateLike(ctx context.Context, arg CreateLikeParams) (*mongo.InsertOneResult, error) {
	if _, liked := q.IsLiked(ctx, arg); liked {
		return nil, ErrDuplicatedLike
	}

	like := Like{
		ID:        primitive.NewObjectID(),
		UserID:    arg.UserID,
		TargetID:  arg.TargetID,
		CreatedAt: time.Now(),
	}

	coll := q.db.Collection("likes")
	result, err := coll.InsertOne(ctx, like)

	return result, err
}

func (q *Queries) DeleteLike(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}

	coll := q.db.Collection("likes")
	result, err := coll.DeleteOne(ctx, filter)

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
	TargetID primitive.ObjectID `json:"target_id" bson:"target_id"`
	Offset   int64              `json:"offset" bson:"limit"`
	Limit    int64              `json:"limit" bson:"limit"`
}

func (q *Queries) ListLikes(ctx context.Context, arg ListLikesParams) ([]Like, error) {
	filter := bson.D{primitive.E{Key: "target_id", Value: arg.TargetID}}

	var likes []Like
	coll := q.db.Collection("likes")
	cursor, err := coll.Find(ctx, filter, options.Find().SetSkip(arg.Offset).SetLimit(arg.Limit))
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

func (q *Queries) CountLikes(ctx context.Context, targetID primitive.ObjectID) (int64, error) {
	filter := bson.D{primitive.E{Key: "target_id", Value: targetID}}

	coll := q.db.Collection("likes")
	nLikes, err := coll.CountDocuments(ctx, filter, nil)

	return nLikes, err
}

func (q *Queries) ToggleLike(ctx context.Context, arg CreateLikeParams) (bool, error) {
	if like, liked := q.IsLiked(ctx, arg); liked {
		_, err := q.DeleteLike(ctx, like.ID)
		return false, err
	} else {
		_, err := q.CreateLike(ctx, arg)
		return true, err
	}
}
