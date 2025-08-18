package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"vitaliesvet.com/post-rest-app/rest/middleware"
)

func RegisterRoutes(server *gin.Engine, controllers Controllers) {
	middleware.PrometheusInit()

	// Prometheus metrics endpoint
	server.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// Middleware to track request metrics
	server.Use(middleware.TrackMetrics())

	server.POST("/posts", controllers.PostController.createPost) //GET, POST, PUT, PATCH , DELETE ...
	server.GET("/posts/", controllers.PostController.getPosts)
}
