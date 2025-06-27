package controllers

import "learning/app/middlewares"

func (s *Server) SetupRoutes() {
	s.Router.GET("/", s.Health)

	s.Router.POST("/users/signup", s.CreateUser)
	s.Router.POST("/users/login", s.LoginUser)

	authRoutes := s.Router.Group("/").Use(middlewares.RequireUser(s.TokenMaker))
	authRoutes.POST("/todos", s.CreateTodo)
	authRoutes.GET("/todos", s.GetTodos)
	authRoutes.PUT("/todos/:id", s.UpdateTodo)
	authRoutes.DELETE("/todos/:id", s.DestroyTodo)
}
