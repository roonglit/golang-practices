package main

import (
	"context"
	"fmt"
	"learning/db"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	connectionString := "postgres://postgres:postgres@localhost:5432/learning?sslmode=disable"
	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		fmt.Println("Unable to connect to database:", err)
		return
	}

	queries := db.New(conn)

	// create a new user
	createdUser, err := queries.CreateUser(ctx, db.CreateUserParams{
		Username: "John Doe",
		Email:    "john@doe.com",
		Password: "password123",
	})
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	fmt.Println(createdUser)

	// get the user we just inserted
	user, err := queries.GetUser(ctx, createdUser.ID)
	if err != nil {
		fmt.Println("Error getting user:", err)
		return
	}
	fmt.Println("User:", user)

	// update the user
	updatedUser, err := queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:       user.ID,
		Username: "Jane Doe",
		Email:    "jane@doe.com",
	})
	if err != nil {
		fmt.Println("Error updating user:", err)
		return
	}
	fmt.Println("Updated User:", updatedUser)
}
