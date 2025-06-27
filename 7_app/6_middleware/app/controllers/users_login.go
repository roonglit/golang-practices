package controllers

import (
	"learning/app/models"
	"learning/app/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func (s *Server) LoginUser(c *gin.Context) {
	var loginRequest LoginUserRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, BadRequestError(err))
		return
	}

	user, err := s.Store.GetUserByUsername(c, loginRequest.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, NotFoundError(err))
			return
		}
		c.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}

	err = util.CheckPassword(loginRequest.Password, user.PasswordHash)
	if err != nil {
		c.JSON(http.StatusUnauthorized, UnauthorizedError(err))
		return
	}

	accessToken, _, err := s.TokenMaker.CreateToken(
		user.Username,
		s.Config.AccessTokenDuration,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InternalServerError(err))
		return
	}

	c.JSON(http.StatusOK, loginResponse(accessToken, user))
}

func loginResponse(accessToken string, user models.User) Login {
	return Login{
		AccessToken: accessToken,
		User: User{
			ID:        int(user.ID),
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Time.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		},
	}
}
