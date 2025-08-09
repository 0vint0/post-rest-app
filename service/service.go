package service

import (
	"vitaliesvet.com/post-rest-app/persistent/repository"
	"vitaliesvet.com/post-rest-app/rest/view"
)

type PostService interface {
	Create(p view.PostView) (view.PostView, error)
	FindAll() ([]view.PostView, error)
}

type Services struct {
	PostService PostService
}

func Init(r *repository.Repositories) *Services {
	return &Services{
		PostService: NewPostService(&r.PostRepository),
	}
}
