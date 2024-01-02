package db

import (
	"testing"
	"time"

	"github.com/DMV-Nicolas/robotgram/backend/util"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func randomComment(t *testing.T, userID, targetID primitive.ObjectID) Comment {
	arg := CreateCommentParams{
		UserID:   userID,
		TargetID: targetID,
		Content:  util.RandomString(25),
	}

	result, err := testQueries.CreateComment(testCtx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	require.True(t, ok)
	require.NotEqual(t, primitive.NilObjectID, insertedID)

	comment, err := testQueries.GetComment(testCtx, insertedID)
	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, insertedID, comment.ID)
	require.Equal(t, arg.UserID, comment.UserID)
	require.Equal(t, arg.TargetID, comment.TargetID)
	require.Equal(t, arg.Content, comment.Content)
	require.WithinDuration(t, time.Now(), comment.CreatedAt, time.Second)

	return comment
}

func TestCreateComment(t *testing.T) {
	randomComment(t, primitive.NewObjectID(), primitive.NewObjectID())
}

func TestGetComment(t *testing.T) {
	comment1 := randomComment(t, primitive.NewObjectID(), primitive.NewObjectID())

	comment2, err := testQueries.GetComment(testCtx, comment1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, comment2)

	require.Equal(t, comment1.ID, comment2.ID)
	require.Equal(t, comment1.UserID, comment2.UserID)
	require.Equal(t, comment1.TargetID, comment2.TargetID)
	require.Equal(t, comment1.Content, comment2.Content)
	require.WithinDuration(t, comment1.CreatedAt, comment2.CreatedAt, time.Second)
}