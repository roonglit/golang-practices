package controllers

import (
	"learning/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetTodos(c *gin.Context) {
	todos, err := s.Store.GetTodos(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}

	c.JSON(http.StatusOK, getTodosResponse(todos))
}

func getTodosResponse(todos []models.Todo) []Todo {
	var response []Todo
	for _, todo := range todos {
		response = append(response, Todo{
			ID:          int(todo.ID),
			Title:       todo.Title,
			Description: todo.Description.String,
			Completed:   &todo.Completed.Bool,
			CreatedAt:   todo.CreatedAt.Time.Format("2006-01-02 15:04:05"),
			UpdatedAt:   todo.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		})
	}
	return response
}
