package main

import "github.com/gin-gonic/gin"

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Search struct {
	Query string `form:"query" binding:"required"`
}

type UserUri struct {
	Name string `uri:"name" binding:"required"`
}

func main() {
	r := gin.Default()

	// JSON body parsing
	r.POST("/signup", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Welcome " + user.Name})
	})

	// Query parameters
	r.GET("/search", func(c *gin.Context) {
		var search Search
		if err := c.ShouldBindQuery(&search); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"query": search.Query})
	})

	// Uri parameters
	r.GET("/user/:name", func(c *gin.Context) {
		var user UserUri
		if err := c.ShouldBindUri(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Hello " + user.Name})
	})

	r.Run()
}
