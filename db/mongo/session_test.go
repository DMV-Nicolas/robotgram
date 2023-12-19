package db

import (
	"testing"
	"time"

	"github.com/DMV-Nicolas/tinygram/util"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func randomSession(t *testing.T) Session {
	user := randomUser(t)
	arg := CreateSessionParams{
		ID:           primitive.NewObjectID(),
		UserID:       user.ID,
		RefreshToken: util.RandomString(30),
		UserAgent:    util.RandomString(10),
		ClientIP:     util.RandomString(10),
		IsBlocked:    false,
		ExpiresAt:    time.Now().Add(time.Minute),
	}

	result, err := testQueries.CreateSession(testCtx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	require.True(t, ok)
	require.NotEqual(t, primitive.NilObjectID, insertedID)

	session, err := testQueries.GetSession(testCtx, insertedID)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	require.Equal(t, insertedID, session.ID)
	require.Equal(t, arg.UserID, session.UserID)
	require.Equal(t, arg.RefreshToken, session.RefreshToken)
	require.Equal(t, arg.UserAgent, session.UserAgent)
	require.Equal(t, arg.ClientIP, session.ClientIP)
	require.Equal(t, arg.IsBlocked, session.IsBlocked)
	require.WithinDuration(t, arg.ExpiresAt, session.ExpiresAt, time.Second)

	return session
}

func TestCreateSession(t *testing.T) {
	randomSession(t)
}

func TestGetSession(t *testing.T) {
	session1 := randomSession(t)

	session2, err := testQueries.GetSession(testCtx, session1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, session2)

	require.Equal(t, session1.ID, session2.ID)
	require.Equal(t, session1.UserID, session2.UserID)
	require.Equal(t, session1.RefreshToken, session2.RefreshToken)
	require.Equal(t, session1.UserAgent, session2.UserAgent)
	require.Equal(t, session1.ClientIP, session2.ClientIP)
	require.Equal(t, session1.IsBlocked, session2.IsBlocked)

	require.WithinDuration(t, session1.ExpiresAt, session2.ExpiresAt, time.Second)

	session3, err := testQueries.GetSession(testCtx, primitive.NewObjectID())
	require.Error(t, err)
	require.Empty(t, session3)
}

func TestDeleteSession(t *testing.T) {
	session1 := randomSession(t)

	result, err := testQueries.DeleteSession(testCtx, session1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.EqualValues(t, 1, result.DeletedCount)

	session2, err := testQueries.GetSession(testCtx, session1.ID)
	require.Error(t, err)
	require.EqualError(t, mongo.ErrNoDocuments, err.Error())
	require.Empty(t, session2)
}

func TestBlockSession(t *testing.T) {
	session1 := randomSession(t)

	result, err := testQueries.BlockSession(testCtx, session1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.EqualValues(t, 1, result.MatchedCount)
	require.EqualValues(t, 1, result.ModifiedCount)

	session2, err := testQueries.GetSession(testCtx, session1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, session2)

	require.True(t, session2.IsBlocked)
}
