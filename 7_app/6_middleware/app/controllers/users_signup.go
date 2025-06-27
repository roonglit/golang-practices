package controllers

import (
	"learning/app/models"
	"learning/app/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) CreateUser(c *gin.Context) {
	var user CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, BadRequestError(err))
		return
	}

	createdUser, err := s.Store.CreateUser(c, createUserParams(user))
	if err != nil {
		c.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}

	c.JSON(http.StatusCreated, userResponse(createdUser))
}

func createUserParams(user CreateUserRequest) models.CreateUserParams {
	passwordHash, _ := util.HashPassword(user.Password)
	return models.CreateUserParams{
		Username:     user.Username,
		PasswordHash: passwordHash,
		Email:        user.Email,
		CreatedAt:    pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		UpdatedAt:    pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
	}
}

func userResponse(user models.User) User {
	return User{
		ID:        int(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
	}
}
