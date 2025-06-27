package controllers

import (
	"fmt"
	"learning/app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) UpdateTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, BadRequestError(err))
		return
	}

	var resource Resource
	if err := c.ShouldBindUri(&resource); err != nil {
		c.JSON(http.StatusBadRequest, BadRequestError(err))
		return
	}
	updatedTodo, err := s.Store.UpdateTodo(c.Request.Context(), updateTodoParams(resource.ID, todo))
	if err != nil {
		fmt.Println("Error updating todo:", err)
		c.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}

	c.JSON(http.StatusOK, todoResponse(updatedTodo))
}

func updateTodoParams(id int32, todo Todo) models.UpdateTodoParams {
	return models.UpdateTodoParams{
		ID:          id,
		Title:       todo.Title,
		Description: pgtype.Text{String: todo.Description, Valid: true},
		Completed:   pgtype.Bool{Bool: *todo.Completed, Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
	}
}
