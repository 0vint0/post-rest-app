package rest

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, controllers Controllers) {
	server.POST("/posts", controllers.PostController.createPost) //GET, POST, PUT, PATCH , DELETE ...
	server.GET("/posts/", controllers.PostController.getPosts)
}
