package db

import (
	"testing"

	"github.com/DMV-Nicolas/robotgram/util"
	"github.com/stretchr/testify/require"
)

func TestUsernameTaken(t *testing.T) {
	user1 := randomUser(t)

	arg := CreateUserParams{
		Username:       user1.Username,
		HashedPassword: util.RandomPassword(16),
		FullName:       util.RandomUsername(),
		Email:          util.RandomEmail(),
		Avatar:         "avatar.png",
		Gender:         "male",
	}

	result, err := testQueries.CreateUser(testCtx, arg)
	require.Error(t, err)
	require.EqualError(t, ErrUsernameTaken, err.Error())
	require.Empty(t, result)
}

func TestEmailTaken(t *testing.T) {
	user1 := randomUser(t)

	arg := CreateUserParams{
		Username:       util.RandomUsername(),
		HashedPassword: util.RandomPassword(16),
		FullName:       util.RandomUsername(),
		Email:          user1.Email,
		Avatar:         "avatar.png",
		Gender:         "male",
	}

	result, err := testQueries.CreateUser(testCtx, arg)
	require.Error(t, err)
	require.EqualError(t, ErrEmailTaken, err.Error())
	require.Empty(t, result)
}

func TestDuplicatedLike(t *testing.T) {
	post := randomPost(t)
	like := randomLike(t, post.ID)

	arg := CreateLikeParams{
		UserID:   like.UserID,
		TargetID: like.TargetID,
	}

	result, err := testQueries.CreateLike(testCtx, arg)
	require.Error(t, err)
	require.EqualError(t, ErrDuplicatedLike, err.Error())
	require.Empty(t, result)
}
