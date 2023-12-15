package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreatePostParams struct {
	Owner       string   `json:"owner" bson:"owner"`
	Images      []string `json:"images" bson:"images"`
	Videos      []string `json:"videos" bson:"videos"`
	Description string   `json:"description" bson:"description"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	post := Post{
		ID:          primitive.NewObjectID(),
		Owner:       arg.Owner,
		Images:      arg.Images,
		Videos:      arg.Videos,
		Description: arg.Description,
		CreatedAt:   time.Now(),
	}

	coll := q.db.Collection("posts")
	_, err := coll.InsertOne(ctx, post)

	return post, err
}

func (q *Queries) GetPost(ctx context.Context, id primitive.ObjectID) (Post, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
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

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) error {
	filter := bson.M{"_id": arg.ID}
	update := bson.M{
		"$set": bson.M{
			"images":      arg.Images,
			"videos":      arg.Videos,
			"description": arg.Description,
		},
	}

	coll := q.db.Collection("posts")
	_, err := coll.UpdateOne(ctx, filter, update)

	return err
}

func (q *Queries) DeletePost(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	coll := q.db.Collection("posts")
	_, err := coll.DeleteOne(ctx, filter)

	return err
}
