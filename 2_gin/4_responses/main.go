package main

import "github.com/gin-gonic/gin"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()

	// gin.H response
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	// json object response
	r.GET("/user", func(c *gin.Context) {
		user := User{
			ID:    1,
			Name:  "Alice",
			Email: "alice@example.com",
		}
		c.JSON(200, user)
	})

	// xml response
	r.GET("/xml", func(c *gin.Context) {
		user := User{
			ID:    2,
			Name:  "Bob",
			Email: "bob@example.com",
		}
		c.XML(200, user)
	})

	r.Run()
}
