package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateCommentParams struct {
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	TargetID primitive.ObjectID `json:"target_id" bson:"target_id"`
	Content  string             `json:"content" bson:"content"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (*mongo.InsertOneResult, error) {
	comment := Comment{
		ID:        primitive.NewObjectID(),
		UserID:    arg.UserID,
		TargetID:  arg.TargetID,
		Content:   arg.Content,
		CreatedAt: time.Now(),
	}

	coll := q.db.Collection("comments")
	result, err := coll.InsertOne(ctx, comment)

	return result, err
}

func (q *Queries) GetComment(ctx context.Context, id primitive.ObjectID) (Comment, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	opts := options.FindOne()

	var comment Comment
	coll := q.db.Collection("comments")
	err := coll.FindOne(ctx, filter, opts).Decode(&comment)

	return comment, err
}

/*type ListCommentsParams struct {
	TargetID primitive.ObjectID `json:"target_id" bson:"target_id"`
	Offset   int64              `json:"offset" bson:"limit"`
	Limit    int64              `json:"limit" bson:"limit"`
}

func (q *Queries) ListComments(ctx context.Context, arg ListCommentsParams) ([]Comment, error) {
	filter := bson.D{primitive.E{Key: "target_id", Value: arg.TargetID}}

	var comments []Comment
	coll := q.db.Collection("comments")
	cursor, err := coll.Find(ctx, filter, options.Find().SetSkip(arg.Offset).SetLimit(arg.Limit))
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var comment Comment
		err = cursor.Decode(&comment)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}*/
