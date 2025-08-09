package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vitaliesvet.com/post-rest-app/rest/view"
	"vitaliesvet.com/post-rest-app/service"
)

type PostControllerImpl struct {
	postService *service.PostService
}

func (p *PostControllerImpl) createPost(context *gin.Context) {
	var post view.PostView
	error := context.ShouldBindJSON(&post)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse provided values!"})
		return
	}
	post, error = (*p.postService).Create(post)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create post!", "error": error})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Post Created", "post": post})
}

func (p *PostControllerImpl) getPosts(context *gin.Context) {
	posts, error := (*p.postService).FindAll()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch posts!", "error": error})
		return
	}
	context.JSON(http.StatusOK, posts)

}

func NewPostController(postService *service.PostService) PostController {
	return &PostControllerImpl{
		postService: postService,
	}
}
