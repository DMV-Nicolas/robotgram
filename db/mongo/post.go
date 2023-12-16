package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreatePostParams struct {
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Images      []string           `json:"images" bson:"images"`
	Videos      []string           `json:"videos" bson:"videos"`
	Description string             `json:"description" bson:"description"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (*mongo.InsertOneResult, error) {
	post := Post{
		ID:          primitive.NewObjectID(),
		UserID:      arg.UserID,
		Images:      arg.Images,
		Videos:      arg.Videos,
		Description: arg.Description,
		CreatedAt:   time.Now(),
	}

	coll := q.db.Collection("posts")
	result, err := coll.InsertOne(ctx, post)

	return result, err
}

func (q *Queries) GetPost(ctx context.Context, key string, value any) (Post, error) {
	filter := bson.D{primitive.E{Key: key, Value: value}}
	opts := options.FindOne()

	var post Post
	coll := q.db.Collection("posts")
	err := coll.FindOne(ctx, filter, opts).Decode(&post)

	return post, err
}

func (q *Queries) ListPosts(ctx context.Context, limit int) ([]Post, error) {
	filter := bson.D{}

	var posts []Post
	coll := q.db.Collection("posts")
	cursor, err := coll.Find(ctx, filter, options.Find().SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var post Post
		err = cursor.Decode(&post)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

type UpdatePostParams struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Images      []string           `json:"images" bson:"images"`
	Videos      []string           `json:"videos" bson:"videos"`
	Description string             `json:"description" bson:"description"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": arg.ID}
	update := bson.M{
		"$set": bson.M{
			"images":      arg.Images,
			"videos":      arg.Videos,
			"description": arg.Description,
		},
	}

	coll := q.db.Collection("posts")
	result, err := coll.UpdateOne(ctx, filter, update)

	return result, err
}

func (q *Queries) DeletePost(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}

	coll := q.db.Collection("posts")
	result, err := coll.DeleteOne(ctx, filter)

	return result, err
}
