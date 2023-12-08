package api

import (
	"golang101/app/model"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *model.Store
	router *gin.Engine
}

func NewServer(store *model.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/", server.getIndex)
	router.POST("/users", server.createUser)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
