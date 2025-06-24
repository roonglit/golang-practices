package model

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListPostsOfUser(t *testing.T) {
	ctx := context.Background()

	user2 := genUser("tester2", "test2@example.com")
	post1 := genPost("First article", "Hello first article", user2.ID)
	post2 := genPost("Second article", "Hello second article", user2.ID)

	user3 := genUser("tester3", "test3@example.com")
	post3 := genPost("Third article", "Hello third article", user3.ID)

	posts, err := testQueries.ListPostsOfUser(ctx, user2.ID)

	require.NoError(t, err)
	require.NotEmpty(t, posts)

	require.Equal(t, 2, len(posts))
	require.Equal(t, post1.Title, posts[0].Title)
	require.Equal(t, post2.Title, posts[1].Title)

	posts, err = testQueries.ListPostsOfUser(ctx, user3.ID)
	require.NoError(t, err)
	require.NotEmpty(t, posts)

	require.Equal(t, 1, len(posts))
	require.Equal(t, post3.Title, posts[0].Title)
}

func genUser(username, email string) User {
	ctx := context.Background()

	// create a user
	userParams := CreateUserParams{
		Username: username,
		Email:    email,
	}

	result, _ := testQueries.CreateUser(ctx, userParams)
	userId, _ := result.LastInsertId()
	user, _ := testQueries.GetUser(ctx, userId)

	return user
}

func genPost(title, content string, userId int64) Post {
	ctx := context.Background()

	postParams := CreatePostParams{
		Title:    title,
		Content:  content,
		AuthorID: userId,
	}

	result, _ := testQueries.CreatePost(ctx, postParams)
	postId, _ := result.LastInsertId()
	post, _ := testQueries.GetPost(ctx, postId)

	return post
}
