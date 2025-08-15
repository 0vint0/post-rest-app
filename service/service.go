package service

import (
	"vitaliesvet.com/post-rest-app/persistent/repository"
	"vitaliesvet.com/post-rest-app/rest/view"
)

type PostService[T view.Viewer] interface {
	Create(p view.PostView) (view.PostView, error)
	FindAll(p view.PagedRequest) (view.Paginated[T], error)
}

type Services struct {
	PostService PostService[view.PostView]
}

func Init(r *repository.Repositories) *Services {
	return &Services{
		PostService: NewPostService(r.PostRepository),
	}
}
