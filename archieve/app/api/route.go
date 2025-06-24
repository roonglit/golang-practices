package api

func (server *Server) routesDraw() {
	router := server.router
	router.GET("/", server.getIndex)
	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.listUsers)
}
