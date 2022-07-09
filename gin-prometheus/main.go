package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	router := gin.Default()

	requestCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "custom_http_request_counter",
			Help: "Number of HTTP requests",
		})
	prometheus.Register(requestCounter)

	router.GET("/", func(context *gin.Context) {
		// increase the counter value
		requestCounter.Inc()

		context.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.Run(":8080")
}
