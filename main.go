package main

import (
	"github.com/gin-gonic/gin"
	"vitaliesvet.com/post-rest-app/persistent/db"
	"vitaliesvet.com/post-rest-app/persistent/repository"
	"vitaliesvet.com/post-rest-app/rest"
	"vitaliesvet.com/post-rest-app/service"
)

func main() {
	db := db.InitDB()
	repositories := repository.Init(db)
	services := service.Init(repositories)
	controllers := rest.Init(services)
	server := gin.Default()
	rest.RegisterRoutes(server, controllers)
	server.Run(":8080")
}
