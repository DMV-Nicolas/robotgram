package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ToggleLikeParams struct {
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	TargetID primitive.ObjectID `json:"target_id" bson:"target_id"`
}

func (q *Queries) ToggleLike(ctx context.Context, arg ToggleLikeParams) (*mongo.InsertOneResult, *mongo.DeleteResult, error) {
	like, liked, err := q.IsLiked(ctx, IsLikedParams{TargetID: arg.TargetID, UserID: arg.UserID})

	if err != nil {
		return nil, nil, err
	}

	if liked {
		// delete like
		filter := bson.M{"_id": like.ID}

		coll := q.db.Collection("likes")
		result, err := coll.DeleteOne(ctx, filter)

		return nil, result, err
	} else {
		// create like
		like := Like{
			ID:        primitive.NewObjectID(),
			UserID:    arg.UserID,
			TargetID:  arg.TargetID,
			CreatedAt: time.Now(),
		}

		coll := q.db.Collection("likes")
		result, err := coll.InsertOne(ctx, like)

		return result, nil, err
	}
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

type IsLikedParams struct {
	TargetID primitive.ObjectID `json:"target_id" bson:"target_id"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
}

func (q *Queries) IsLiked(ctx context.Context, arg IsLikedParams) (Like, bool, error) {
	filter := bson.D{
		primitive.E{Key: "user_id", Value: arg.UserID},
		primitive.E{Key: "target_id", Value: arg.TargetID},
	}
	opts := options.FindOne()

	var like Like
	coll := q.db.Collection("likes")
	err := coll.FindOne(ctx, filter, opts).Decode(&like)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Like{}, false, nil
		} else {
			return Like{}, false, err
		}
	}

	return like, true, nil
}
