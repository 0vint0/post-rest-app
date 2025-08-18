package rest

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"vitaliesvet.com/post-rest-app/error_handler"
	"vitaliesvet.com/post-rest-app/rest/view"
	"vitaliesvet.com/post-rest-app/service"
)

type PostControllerImpl struct {
	postService *service.PostService[view.PostView]
}

func (p *PostControllerImpl) createPost(context *gin.Context) {
	var post view.PostView
	error := context.ShouldBindJSON(&post)
	if !handleError(context, error, "Could not parse provided values!") {
		return
	}
	post, error = (*p.postService).Create(post)
	if !handleError(context, error) {
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Post Created", "post": post})
}

func (p *PostControllerImpl) getPosts(context *gin.Context) {

	pageNumber, error := strconv.ParseInt(context.DefaultQuery("pageNumber", "0"), 10, 64)

	if error != nil {
		fmt.Println("Error ", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could parse pageNumber!"})
		return
	}

	pageSize, error := strconv.ParseInt(context.DefaultQuery("pageSize", "10"), 10, 64)

	if error != nil {
		fmt.Println("Error ", error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could parse pageSize!"})
		return
	}

	paginatedPosts, error := (*p.postService).FindAll(view.NewPagedRequest(pageNumber, pageSize))
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch posts!", "error": error})
		return
	}
	context.JSON(http.StatusOK, paginatedPosts)

}

func NewPostController(postService *service.PostService[view.PostView]) PostController {
	return &PostControllerImpl{
		postService: postService,
	}
}

func handleError(context *gin.Context, error error, message ...string) bool {
	if error != nil {
		if message != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": message})
			return false
		}
		var vErr *error_handler.ValidationError
		if errors.As(error, &vErr) {
			context.JSON(http.StatusPreconditionFailed, gin.H{ // 412
				"error":   "validation_failed",
				"details": vErr.Fields,
			})
			return false
		}
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal",
			"message": "something went wrong",
			"details": error,
		})

	}
	return true
}
