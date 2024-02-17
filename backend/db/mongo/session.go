package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateSessionParams struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	RefreshToken string             `json:"refresh_token" bson:"refresh_token"`
	UserAgent    string             `json:"user_agent" bson:"user_agent"`
	ClientIP     string             `json:"client_ip" bson:"client_ip"`
	IsBlocked    bool               `json:"is_blocked" bson:"is_blocked"`
	ExpiresAt    time.Time          `json:"expires_at" bson:"expires_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (*mongo.InsertOneResult, error) {
	session := Session(arg)

	coll := q.db.Collection("sessions")
	result, err := coll.InsertOne(ctx, session)

	return result, err
}

func (q *Queries) GetSession(ctx context.Context, id primitive.ObjectID) (Session, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	opts := options.FindOne()

	var session Session
	coll := q.db.Collection("sessions")
	err := coll.FindOne(ctx, filter, opts).Decode(&session)

	return session, err
}

func (q *Queries) DeleteSession(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}

	coll := q.db.Collection("sessions")
	result, err := coll.DeleteOne(ctx, filter)

	return result, err
}

func (q *Queries) BlockSession(ctx context.Context, id primitive.ObjectID) (*mongo.UpdateResult, error) {
	update := bson.M{
		"$set": bson.M{
			"is_blocked": true,
		},
	}

	coll := q.db.Collection("sessions")
	result, err := coll.UpdateByID(ctx, id, update)

	return result, err
}
