package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	Ping(c *gin.Context)
	ListAllCats(c *gin.Context)
	CreateCat(c *gin.Context)
}