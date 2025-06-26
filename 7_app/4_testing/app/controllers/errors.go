package controllers

import "github.com/gin-gonic/gin"

func BadRequestError(err error) gin.H {
	return gin.H{"error": "Bad Request", "message": err.Error()}
}

func InternalServerError(err error) gin.H {
	return gin.H{"error": "Internal Server Error", "message": err.Error()}
}
