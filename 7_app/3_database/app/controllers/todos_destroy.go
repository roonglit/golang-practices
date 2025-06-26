package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) DestroyTodo(c *gin.Context) {
	var resource Resource
	if err := c.ShouldBindUri(&resource); err != nil {
		c.JSON(http.StatusBadRequest, BadRequestError(err))
		return
	}

	err := s.Store.DeleteTodo(c.Request.Context(), resource.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}

	c.Status(204) // No Content
}
