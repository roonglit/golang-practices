package controllers

type Resource struct {
	ID int32 `uri:"id" binding:"required"`
}
