package token

import (
	"testing"
	"time"

	"github.com/DMV-Nicolas/tinygram/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoToken(t *testing.T) {
	userID := util.RandomID()
	duration := time.Minute
	issuedAt := time.Now()
	expiresAt := time.Now().Add(duration)

	maker, err := NewPasetoMaker(util.RandomPassword(32))
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	token, payload, err := maker.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, userID, payload.UserID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiresAt, payload.ExpiresAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	userID := util.RandomID()
	duration := -time.Minute

	maker, err := NewPasetoMaker(util.RandomPassword(32))
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	token, payload, err := maker.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.Empty(t, payload)
}

func TestWrongToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomPassword(32))
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	payload, err := maker.VerifyToken("wrong-token")
	require.Error(t, err)
	require.Empty(t, payload)
}

func TestWrongSymmetricKey(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomPassword(31))
	require.Error(t, err)
	require.Empty(t, maker)
}
