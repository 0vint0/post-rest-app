package service

import "vitaliesvet.com/post-rest-app/persistent/repository"

type PostService interface{}

type Services struct {
	PostService PostService
}

func Init(r *repository.Repositories) *Services {
	return &Services{
		PostService: NewPostService(&r.PostRepository),
	}
}
