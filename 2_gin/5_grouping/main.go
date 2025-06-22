package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to the admin dashboard!"})
		})
		adminGroup.GET("/settings", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Admin settings page"})
		})
	}

	r.Run() // Default: localhost:8080
}
