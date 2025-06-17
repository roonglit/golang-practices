package model

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	arg := CreateUserParams{
		Username: "tester1",
		Email:    "test1@example.com",
	}

	result, err := testQueries.CreateUser(ctx, arg)
	require.NoError(t, err)
	insertedUserID, err := result.LastInsertId()
	require.NoError(t, err)
	user, err := testQueries.GetUser(ctx, insertedUserID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
}
