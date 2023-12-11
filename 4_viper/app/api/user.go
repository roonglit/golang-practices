package api

import (
	"database/sql"
	"golang101/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := model.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
	}

	result, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	insertedUserID, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	user, err := server.store.GetUser(ctx, insertedUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type listUsersRequest struct {
	Page    int64 `form:"page" binding:"required,min=1"`
	PerPage int64 `form:"per_page" binding:"required,min=5"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	println("enter here")
	var req listUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		println(req.Page)
		println(req.PerPage)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := model.ListUsersParams{
		Limit:  req.PerPage,
		Offset: (req.Page - 1) * req.PerPage,
	}

	users, err := server.store.ListUsers(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}
