package rest

import (
	"github.com/gin-gonic/gin"
	"vitaliesvet.com/post-rest-app/service"
)

type PostControllerImpl struct {
	postService *service.PostService
}

func (p *PostControllerImpl) createPost(context *gin.Context) {

}

func (p *PostControllerImpl) getPosts(context *gin.Context) {

}

func NewPostController(postService *service.PostService) PostController {
	return &PostControllerImpl{
		postService: postService,
	}
}
