package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomPassword(16)
	hashedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	hashedPassword2, err := HashPassword(password)
	require.NotEqual(t, hashedPassword1, hashedPassword2)

	err = CheckPassword(password, hashedPassword1)
	require.NoError(t, err)

	wrongPassword := RandomPassword(8) + ":)"
	err = CheckPassword(wrongPassword, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	tooLongPassword := RandomPassword(100)
	hashedPassword3, err := HashPassword(tooLongPassword)
	require.Error(t, err)
	require.EqualError(t, bcrypt.ErrPasswordTooLong, err.Error())
	require.Empty(t, hashedPassword3)
}
