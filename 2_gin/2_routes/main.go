package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{"message": "Hello " + name})
	})

	r.GET("/search", func(c *gin.Context) {
		query := c.Query("q")
		c.JSON(200, gin.H{"query": query})
	})

	r.Run() // Default: localhost:8080
}
