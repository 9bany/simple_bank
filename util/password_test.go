package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashesPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashesPassword)

	err = CheckPassword(password, hashesPassword)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashesPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashesPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashesPassword2)
	require.NotEqual(t, hashesPassword, hashesPassword2)
}