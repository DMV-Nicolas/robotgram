package api

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	db "github.com/DMV-Nicolas/tinygram/db/mongo"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func requireBodyMatchInsertOneResult(t *testing.T, body *bytes.Buffer, result *mongo.InsertOneResult) {
	bodyResult := new(mongo.InsertOneResult)
	err := json.NewDecoder(body).Decode(bodyResult)
	require.NoError(t, err)
	require.NotEmpty(t, bodyResult)

	bodyInsertedID, err := primitive.ObjectIDFromHex(bodyResult.InsertedID.(string))
	require.NoError(t, err)

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	require.True(t, ok)

	require.Equal(t, insertedID, bodyInsertedID)
}

func requireBodyMatchUpdateResult(t *testing.T, body *bytes.Buffer, result *mongo.UpdateResult) {
	bodyResult := new(mongo.UpdateResult)
	err := json.NewDecoder(body).Decode(bodyResult)
	require.NoError(t, err)
	require.NotEmpty(t, bodyResult)

	require.Equal(t, result.MatchedCount, bodyResult.MatchedCount)
	require.Equal(t, result.ModifiedCount, bodyResult.ModifiedCount)
	require.Equal(t, result.UpsertedCount, bodyResult.UpsertedCount)
	require.EqualValues(t, result.UpsertedID, bodyResult.UpsertedID)
}

func requireBodyMatchDeleteResult(t *testing.T, body *bytes.Buffer, result *mongo.DeleteResult) {
	bodyResult := new(mongo.DeleteResult)
	err := json.NewDecoder(body).Decode(bodyResult)
	require.NoError(t, err)
	require.NotEmpty(t, bodyResult)

	require.Equal(t, result.DeletedCount, bodyResult.DeletedCount)
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	bodyResult := new(db.User)
	err := json.NewDecoder(body).Decode(bodyResult)
	require.NoError(t, err)
	require.NotEmpty(t, bodyResult)

	require.Equal(t, bodyResult.ID, user.ID)
	require.Equal(t, bodyResult.Username, user.Username)
	require.Equal(t, bodyResult.Email, user.Email)
	require.Equal(t, bodyResult.Description, user.Description)
	require.Equal(t, bodyResult.FullName, user.FullName)
	require.Equal(t, bodyResult.Avatar, user.Avatar)

	require.WithinDuration(t, bodyResult.CreatedAt, user.CreatedAt, time.Second)
}

func requireBodyMatchUsers(t *testing.T, body *bytes.Buffer, users []db.User) {
	bodyResult := make([]db.User, 0, len(users))
	err := json.NewDecoder(body).Decode(&bodyResult)
	require.NoError(t, err)
	require.NotEmpty(t, bodyResult)

	require.Len(t, bodyResult, len(users))

	for i := range bodyResult {
		require.Equal(t, users[i].ID, bodyResult[i].ID)
		require.Equal(t, users[i].Username, bodyResult[i].Username)
		require.Equal(t, users[i].FullName, bodyResult[i].FullName)
		require.Equal(t, users[i].Avatar, bodyResult[i].Avatar)
		require.Equal(t, users[i].Gender, bodyResult[i].Gender)
		require.WithinDuration(t, users[i].CreatedAt, bodyResult[i].CreatedAt, time.Second)
	}
}

func requireBodyMatchPost(t *testing.T, body *bytes.Buffer, post db.Post) {
	bodyResult := new(db.Post)
	err := json.NewDecoder(body).Decode(bodyResult)
	require.NoError(t, err)
	require.NotEmpty(t, bodyResult)

	require.Equal(t, bodyResult.ID, post.ID)
	require.Equal(t, bodyResult.UserID, post.UserID)
	require.Equal(t, bodyResult.Images, post.Images)
	require.Equal(t, bodyResult.Videos, post.Videos)
	require.Equal(t, bodyResult.Description, post.Description)

	require.WithinDuration(t, bodyResult.CreatedAt, post.CreatedAt, time.Second)
}

func requireBodyMatchPosts(t *testing.T, body *bytes.Buffer, posts []db.Post) {
	bodyResult := make([]db.Post, 0, len(posts))
	err := json.NewDecoder(body).Decode(&bodyResult)
	require.NoError(t, err)
	require.NotEmpty(t, bodyResult)

	require.Len(t, bodyResult, len(posts))

	for i := range bodyResult {
		require.Equal(t, posts[i].ID, bodyResult[i].ID)
		require.Equal(t, posts[i].UserID, bodyResult[i].UserID)
		require.Equal(t, posts[i].Images, bodyResult[i].Images)
		require.Equal(t, posts[i].Videos, bodyResult[i].Videos)
		require.Equal(t, posts[i].Description, bodyResult[i].Description)
		require.WithinDuration(t, posts[i].CreatedAt, bodyResult[i].CreatedAt, time.Second)
	}
}
