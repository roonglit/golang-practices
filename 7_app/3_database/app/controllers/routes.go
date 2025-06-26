package controllers

func (s *Server) SetupRoutes() {
	s.Router.GET("/", s.Health)

	s.Router.POST("/todos", s.CreateTodo)
	s.Router.GET("/todos", s.GetTodos)
	s.Router.PUT("/todos/:id", s.UpdateTodo)
	s.Router.DELETE("/todos/:id", s.DestroyTodo)
}
