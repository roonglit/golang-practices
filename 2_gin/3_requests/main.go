package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

type RequestHeader struct {
	Authorization string `header:"Authorization" binding:"required"` // Map "Authorization" header
	UserAgent     string `header:"User-Agent" binding:"required"`    // Map "User-Agent" header
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

	// Request headers
	r.GET("/header", func(c *gin.Context) {
		var headers RequestHeader

		// Bind header values to the struct
		if err := c.ShouldBindHeader(&headers); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Respond with the extracted header values
		c.JSON(http.StatusOK, gin.H{
			"Authorization": headers.Authorization,
			"User-Agent":    headers.UserAgent,
		})
	})

	r.Run()
}
