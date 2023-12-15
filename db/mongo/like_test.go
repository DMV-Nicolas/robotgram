package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func randomLike(t *testing.T, postID primitive.ObjectID) Like {
	user := randomUser(t)
	arg := CreateLikeParams{
		UserID: user.ID,
		PostID: postID,
	}

	result, err := testQueries.CreateLike(testCtx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	require.True(t, ok)
	require.NotEqual(t, primitive.NilObjectID, insertedID)

	like, err := testQueries.GetLike(testCtx, insertedID)
	require.NoError(t, err)
	require.NotEmpty(t, like)

	require.Equal(t, insertedID, like.ID)
	require.Equal(t, arg.UserID, like.UserID)
	require.Equal(t, arg.PostID, like.PostID)
	require.WithinDuration(t, time.Now(), like.CreatedAt, time.Second)

	return like
}

func TestCreateLike(t *testing.T) {
	post := randomPost(t)
	randomLike(t, post.ID)
}

func TestGetLike(t *testing.T) {
	post := randomPost(t)
	like1 := randomLike(t, post.ID)

	like2, err := testQueries.GetLike(testCtx, like1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, like2)

	require.Equal(t, like1.ID, like2.ID)
	require.Equal(t, like1.UserID, like2.UserID)
	require.Equal(t, like1.PostID, like2.PostID)
	require.WithinDuration(t, like1.CreatedAt, like2.CreatedAt, time.Second)
}

func TestListLikes(t *testing.T) {
	post := randomPost(t)
	n := 10
	for i := 0; i < n; i++ {
		randomLike(t, post.ID)
	}

	arg := ListLikesParams{
		PostID: post.ID,
		Limit:  10,
	}

	likes, err := testQueries.ListLikes(testCtx, arg)
	require.NoError(t, err)
	require.Len(t, likes, n)

	for _, l := range likes {
		require.NotEmpty(t, l)
	}
}

func TestDeleteLike(t *testing.T) {
	post := randomPost(t)
	like1 := randomLike(t, post.ID)

	result, err := testQueries.DeleteLike(testCtx, like1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.EqualValues(t, 1, result.DeletedCount)

	like2, err := testQueries.GetLike(testCtx, like1.ID)
	require.Error(t, err)
	require.EqualError(t, mongo.ErrNoDocuments, err.Error())
	require.Empty(t, like2)
}
