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

	store := db.NewStore(conn)

	createdUser, err := store.CreateUserWithAuditLog(ctx, db.CreateUserParams{
		Username: "John Doe2",
		Email:    "john@doe2.com",
		Password: "password123",
	})
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	fmt.Println("Created User:", createdUser)
}
