package controllers

// initializeRoutes sets up the routes for the server
func (s *Server) initializeRoutes() {
	// Create a group for the version 1 of the API
	v1 := s.Router.Group("/api/v1")
	{
		//ToDos routes
		v1.POST("/todos", s.CreateToDo)
		v1.PUT("/todos/:id", s.UpdateToDo)
		v1.DELETE("/todos/:id", s.DeleteToDo)
	}
}
