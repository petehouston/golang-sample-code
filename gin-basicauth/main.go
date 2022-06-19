package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Homepage")
	})

	router.GET("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "secret",
	}), func(context *gin.Context) {
		context.String(http.StatusOK, "Welcome to admin dashboard!")
	})

	router.Run()
}
