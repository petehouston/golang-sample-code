package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PersonParams struct {
	Name    string `form:"name"`
	Age     int    `form:"age"`
	IsHired bool   `form:"isHired"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		name := c.Query("name")
		age := c.DefaultQuery("age", "40")
		isHired := c.DefaultQuery("isHired", "false")

		c.String(http.StatusOK, "Name: %s | Age: %s | isHired: %s", name, age, isHired)
	})

	router.GET("/bind", func(c *gin.Context) {
		var params PersonParams

		if c.Bind(&params) == nil {
			c.String(http.StatusOK, "Name: %s | Age: %d | isHired: %t | Department Ids: %s", params.Name, params.Age, params.IsHired)
		}
	})

	router.Run()
}
