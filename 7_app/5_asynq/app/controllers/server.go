package controllers

import (
	"learning/app/models"
	"learning/config"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

type Server struct {
	Router      *gin.Engine
	Config      *config.Config
	Store       models.Store
	AsynqClient *asynq.Client
}

type Option func(*Server)

func New(opts ...Option) *Server {
	router := gin.Default()

	s := &Server{
		Router: router,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.SetupRoutes()

	return s
}

func SetConfig(config *config.Config) Option {
	return func(s *Server) {
		s.Config = config
	}
}

func SetStore(store models.Store) Option {
	return func(s *Server) {
		s.Store = store
	}
}

func SetAsynqClient(client *asynq.Client) Option {
	return func(s *Server) {
		s.AsynqClient = client
	}
}

func (s *Server) Run() {
	s.Router.Run(s.Config.ServerAddress)
}
