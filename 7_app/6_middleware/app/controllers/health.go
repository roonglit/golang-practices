package controllers

import "github.com/gin-gonic/gin"

func (s *Server) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
