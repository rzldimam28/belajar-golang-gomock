package handler

import (
	"belajar-golang-mock/constant"
	"belajar-golang-mock/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.Service
}

func New(service service.Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong!")
}

func (h *handler) ListAllCats(c *gin.Context) {
	cats, err := h.service.ListAllCats(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H {
			"success": false,
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"success": true,
		"data": cats,
	})	
}

func (h *handler) CreateCat(c *gin.Context) {
  var request constant.Request
	err := c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusRequestEntityTooLarge, gin.H {
			"success": false,
			"data": nil,
		})
		return
	}

	id, err := h.service.CreateNewCat(c.Request.Context(), request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H {
			"success": false,
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"success": true,
		"data": gin.H {
			"id": id,
		},
	})
}