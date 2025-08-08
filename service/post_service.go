package service

import "vitaliesvet.com/post-rest-app/persistent/repository"

type PostServiceImpl struct {
	postRepository *repository.PostRepository
}

func NewPostService(postRepository *repository.PostRepository) PostService {
	return PostServiceImpl{
		postRepository: postRepository,
	}
}
