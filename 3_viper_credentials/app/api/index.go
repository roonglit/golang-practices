package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getIndex(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "hello")
}
