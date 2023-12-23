package db

import (
	"testing"
	"time"

	"github.com/DMV-Nicolas/robotgram/util"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func randomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomUsername(),
		HashedPassword: util.RandomPassword(16),
		FullName:       util.RandomUsername(),
		Email:          util.RandomEmail(),
		Avatar:         "avatar.png",
		Gender:         "male",
	}

	result, err := testQueries.CreateUser(testCtx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	require.True(t, ok)
	require.NotEqual(t, primitive.NilObjectID, insertedID)

	user, err := testQueries.GetUser(testCtx, "_id", insertedID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, insertedID, user.ID)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Avatar, user.Avatar)
	require.Equal(t, arg.Gender, user.Gender)
	require.Empty(t, user.Description)
	require.WithinDuration(t, time.Now(), user.CreatedAt, time.Second)

	return user
}

func TestCreateUser(t *testing.T) {
	randomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := randomUser(t)

	user2, err := testQueries.GetUser(testCtx, "_id", user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Avatar, user2.Avatar)
	require.Equal(t, user1.Gender, user2.Gender)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)

	user3, err := testQueries.GetUser(testCtx, "_id", primitive.NewObjectID())
	require.Error(t, err)
	require.Empty(t, user3)

	user4, err := testQueries.GetUser(testCtx, "arroz", "arroz")
	require.Error(t, err)
	require.Empty(t, user4)

	user5, err := testQueries.GetUser(testCtx, "_id", "invalid-object-id")
	require.Error(t, err)
	require.Empty(t, user5)
}

func TestListUsers(t *testing.T) {
	n := 10
	for i := 0; i < n; i++ {
		randomUser(t)
	}

	arg := ListUsersParams{
		Offset: int64(n / 2),
		Limit:  int64(n / 2),
	}

	users, err := testQueries.ListUsers(testCtx, arg)
	require.NoError(t, err)
	require.Len(t, users, n/2)

	for _, u := range users {
		require.NotEmpty(t, u)
	}
}

func TestUpdateUser(t *testing.T) {
	user1 := randomUser(t)

	arg := UpdateUserParams{
		ID:             user1.ID,
		HashedPassword: util.RandomPassword(20),
		FullName:       util.RandomUsername(),
		Description:    util.RandomPassword(100),
		Gender:         "female",
		Avatar:         "other-avatar.png",
	}

	result, err := testQueries.UpdateUser(testCtx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.EqualValues(t, 1, result.MatchedCount)
	require.EqualValues(t, 1, result.ModifiedCount)

	user2, err := testQueries.GetUser(testCtx, "_id", user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.NotEqual(t, user1.HashedPassword, user2.HashedPassword)
	require.NotEqual(t, user1.FullName, user2.FullName)
	require.NotEqual(t, user1.Avatar, user2.Avatar)
	require.NotEqual(t, user1.Gender, user2.Gender)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := randomUser(t)

	result, err := testQueries.DeleteUser(testCtx, user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.EqualValues(t, 1, result.DeletedCount)

	user2, err := testQueries.GetUser(testCtx, "_id", user1.ID)
	require.Error(t, err)
	require.EqualError(t, mongo.ErrNoDocuments, err.Error())
	require.Empty(t, user2)
}
