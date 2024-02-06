package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func randomLike(t *testing.T, userID primitive.ObjectID, targetID primitive.ObjectID) Like {
	arg := ToggleLikeParams{
		UserID:   userID,
		TargetID: targetID,
	}

	createdResult, deletedResult, err := testQueries.ToggleLike(testCtx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, createdResult)
	require.Empty(t, deletedResult)

	insertedID, ok := createdResult.InsertedID.(primitive.ObjectID)
	require.True(t, ok)
	require.NotEqual(t, primitive.NilObjectID, insertedID)

	like, err := testQueries.GetLike(testCtx, insertedID)
	require.NoError(t, err)
	require.NotEmpty(t, like)

	require.Equal(t, insertedID, like.ID)
	require.Equal(t, arg.UserID, like.UserID)
	require.Equal(t, arg.TargetID, like.TargetID)
	require.WithinDuration(t, time.Now(), like.CreatedAt, time.Second)

	return like
}

func TestToggleLike(t *testing.T) {
	user := randomUser(t)
	post := randomPost(t)
	randomLike(t, user.ID, post.ID)

	arg := ToggleLikeParams{
		UserID:   user.ID,
		TargetID: post.ID,
	}

	createdResult, deletedResult, err := testQueries.ToggleLike(testCtx, arg)
	require.NoError(t, err)
	require.Empty(t, createdResult)
	require.NotEmpty(t, deletedResult)

	createdResult, deletedResult, err = testQueries.ToggleLike(testCtx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, createdResult)
	require.Empty(t, deletedResult)
}

func TestGetLike(t *testing.T) {
	like1 := randomLike(t, primitive.NewObjectID(), primitive.NewObjectID())

	like2, err := testQueries.GetLike(testCtx, like1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, like2)

	require.Equal(t, like1.ID, like2.ID)
	require.Equal(t, like1.UserID, like2.UserID)
	require.Equal(t, like1.TargetID, like2.TargetID)
	require.WithinDuration(t, like1.CreatedAt, like2.CreatedAt, time.Second)
}

func TestListLikes(t *testing.T) {
	post := randomPost(t)
	n := 10
	for i := 0; i < n; i++ {
		randomLike(t, primitive.NewObjectID(), post.ID)
	}

	arg := ListLikesParams{
		TargetID: post.ID,
		Offset:   int64(n / 2),
		Limit:    int64(n / 2),
	}

	likes, err := testQueries.ListLikes(testCtx, arg)
	require.NoError(t, err)
	require.Len(t, likes, n/2)

	for _, l := range likes {
		require.NotEmpty(t, l)
	}
}

func TestCountLikes(t *testing.T) {
	post := randomPost(t)
	n := 10
	for i := 0; i < n; i++ {
		randomLike(t, primitive.NewObjectID(), post.ID)
	}

	nLikes, err := testQueries.CountLikes(testCtx, post.ID)
	require.NoError(t, err)
	require.NotZero(t, nLikes)
	require.EqualValues(t, n, nLikes)
}

func TestIsLiked(t *testing.T) {
	user := randomUser(t)
	post := randomPost(t)
	like1 := randomLike(t, user.ID, post.ID)

	arg := IsLikedParams{
		UserID:   user.ID,
		TargetID: post.ID,
	}

	like2, liked := testQueries.IsLiked(testCtx, arg)
	require.True(t, liked)
	require.NotEmpty(t, like2)
	require.Equal(t, like1, like2)

	_, _, err := testQueries.ToggleLike(testCtx, ToggleLikeParams{UserID: arg.UserID, TargetID: arg.TargetID})
	require.NoError(t, err)

	like3, liked := testQueries.IsLiked(testCtx, arg)
	require.False(t, liked)
	require.Empty(t, like3)
}
