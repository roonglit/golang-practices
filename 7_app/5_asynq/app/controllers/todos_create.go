package controllers

import (
	"fmt"
	"learning/app/models"
	"learning/app/tasks"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) CreateTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, BadRequestError(err))
		return
	}

	createdTodo, err := s.Store.CreateTodo(c.Request.Context(), createTodoParams(todo))
	if err != nil {
		fmt.Println("Error creating todo:", err)
		c.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}

	if err := s.sendEmail(todo.Title, todo.Description); err != nil {
		c.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}

	c.JSON(http.StatusCreated, todoResponse(createdTodo))
}

func createTodoParams(todo Todo) models.CreateTodoParams {
	return models.CreateTodoParams{
		Title:       todo.Title,
		Description: pgtype.Text{String: todo.Description, Valid: true},
		Completed:   pgtype.Bool{Bool: *todo.Completed, Valid: true},
		CreatedAt:   pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
	}
}

func todoResponse(todo models.Todo) Todo {
	return Todo{
		ID:          int(todo.ID),
		Title:       todo.Title,
		Description: todo.Description.String,
		Completed:   &todo.Completed.Bool,
		CreatedAt:   todo.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:   todo.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
	}
}

func (s *Server) sendEmail(title string, description string) error {
	task, err := tasks.NewEmailTask(
		"your-email@gmail.com",
		fmt.Sprintf("New to do list %s", title),
		fmt.Sprintf("Title: %s \nDescription: %s\n\nThis is a test email sent from Golang.", title, description),
	)
	if err != nil {
		fmt.Printf("could not create task: %v", err)
		return err
	}

	s.AsynqClient.Enqueue(task)
	if err != nil {
		fmt.Printf("could not enqueue task: %v", err)
		return err
	}

	return nil
}
