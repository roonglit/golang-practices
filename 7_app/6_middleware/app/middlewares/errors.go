package middlewares

import "github.com/gin-gonic/gin"

func AbortError(err error) gin.H {
	return gin.H{"error": "Abort", "message": err.Error()}
}
