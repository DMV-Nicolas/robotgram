package db

import (
	"testing"
	"time"

	"github.com/DMV-Nicolas/tinygram/util"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func randomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomUsername(),
		HashedPassword: util.RandomPassword(16),
		FullName:       util.RandomUsername(),
		Email:          util.RandomEmail(),
		Gender:         "male",
	}

	user, err := testQueries.CreateUser(testCtx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Gender, user.Gender)

	require.Empty(t, user.Description)
	require.NotEqual(t, primitive.NilObjectID, user.ID)
	require.WithinDuration(t, time.Now(), user.CreatedAt, time.Second)

	return user
}

func TestCreateUser(t *testing.T) {
	randomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := randomUser(t)

	user2, err := testQueries.GetUser(testCtx, user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Gender, user2.Gender)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestListUsers(t *testing.T) {
	n := 10
	for i := 0; i < n; i++ {
		randomUser(t)
	}

	users, err := testQueries.ListUsers(testCtx, 10)
	require.NoError(t, err)
	require.Len(t, users, n)

	for _, u := range users {
		require.NotEmpty(t, u)
	}

}
