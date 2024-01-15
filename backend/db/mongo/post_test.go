package db

import (
	"testing"
	"time"

	"github.com/DMV-Nicolas/robotgram/backend/util"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func randomPost(t *testing.T) Post {
	user := randomUser(t)
	arg := CreatePostParams{
		UserID:      user.ID,
		Images:      util.RandomImages(3),
		Description: util.RandomPassword(100),
	}

	result, err := testQueries.CreatePost(testCtx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	require.True(t, ok)
	require.NotEqual(t, primitive.NilObjectID, insertedID)

	post, err := testQueries.GetPost(testCtx, "_id", insertedID)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, insertedID, post.ID)
	require.Equal(t, arg.UserID, post.UserID)
	require.Equal(t, arg.Images, post.Images)
	require.Equal(t, arg.Description, post.Description)
	require.WithinDuration(t, time.Now(), post.CreatedAt, time.Second)

	return post
}

func TestCreatePost(t *testing.T) {
	randomPost(t)
}

func TestGetPost(t *testing.T) {
	post1 := randomPost(t)

	post2, err := testQueries.GetPost(testCtx, "_id", post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.UserID, post2.UserID)
	require.Equal(t, post1.Images, post2.Images)
	require.Equal(t, post1.Description, post2.Description)
	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
}

func TestListPosts(t *testing.T) {
	n := 10
	for i := 0; i < n; i++ {
		randomPost(t)
	}

	arg := ListPostsParams{
		Offset: int64(n / 2),
		Limit:  int64(n / 2),
	}

	posts, err := testQueries.ListPosts(testCtx, arg)
	require.NoError(t, err)
	require.Len(t, posts, n/2)

	for _, p := range posts {
		require.NotEmpty(t, p)
	}
}

func TestUpdatePost(t *testing.T) {
	post1 := randomPost(t)

	arg := UpdatePostParams{
		ID:          post1.ID,
		Images:      util.RandomImages(5),
		Description: util.RandomPassword(200),
	}

	result, err := testQueries.UpdatePost(testCtx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.EqualValues(t, 1, result.MatchedCount)
	require.EqualValues(t, 1, result.ModifiedCount)

	post2, err := testQueries.GetPost(testCtx, "_id", post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.UserID, post2.UserID)
	require.NotEqual(t, post1.Images, post2.Images)
	require.NotEqual(t, post1.Description, post2.Description)
	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
}

func TestDeletePost(t *testing.T) {
	post1 := randomPost(t)

	result, err := testQueries.DeletePost(testCtx, post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.EqualValues(t, 1, result.DeletedCount)

	post2, err := testQueries.GetPost(testCtx, "_id", post1.ID)
	require.Error(t, err)
	require.EqualError(t, mongo.ErrNoDocuments, err.Error())
	require.Empty(t, post2)
}
