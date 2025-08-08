package rest

import (
	"github.com/gin-gonic/gin"
	"vitaliesvet.com/post-rest-app/service"
)

type Controllers struct {
	PostController PostController
}

type PostController interface {
	createPost(context *gin.Context)

	getPosts(context *gin.Context)
}

func Init(services *service.Services) Controllers {
	return Controllers{
		PostController: NewPostController(&services.PostService),
	}
}
