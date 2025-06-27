package controllers

import (
	"learning/app/models"
	"learning/app/util/token"
	"learning/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Router     *gin.Engine
	Config     *config.Config
	Store      models.Store
	TokenMaker token.Maker
}

type Option func(*Server)

func New(opts ...Option) *Server {
	router := gin.Default()

	tokenMaker, err := token.NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		log.Error().Msgf("failed to create token maker: %v", err)
		return nil
	}

	s := &Server{
		Router:     router,
		TokenMaker: tokenMaker,
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

func SetTokenMaker(maker token.Maker) Option {
	return func(s *Server) {
		s.TokenMaker = maker
	}
}

func (s *Server) Run() {
	s.Router.Run(s.Config.ServerAddress)
}
