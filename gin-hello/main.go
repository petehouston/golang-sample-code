package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello world!")
	})
	err := router.Run()
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}
}
