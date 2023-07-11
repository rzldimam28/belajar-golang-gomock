package main

import (
	"belajar-golang-mock/db"
	"belajar-golang-mock/handler"
	"belajar-golang-mock/repository"
	"belajar-golang-mock/service"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	cancellationTimeOut = time.Second * 10
)

func main() {

	db := db.New()
	repository := repository.New()
	service := service.New(repository, db, cancellationTimeOut)
	handler := handler.New(service)

	r := gin.Default()
	r.GET("/ping", handler.Ping)
	r.GET("/cats", handler.ListAllCats)
	r.POST("/cats", handler.CreateCat)

	r.Run(":8899")

}