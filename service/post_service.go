package service

import (
	"github.com/samber/lo"
	model "vitaliesvet.com/post-rest-app/persistent/db/model/post"
	"vitaliesvet.com/post-rest-app/persistent/repository"
	"vitaliesvet.com/post-rest-app/rest/view"
)

type PostServiceImpl struct {
	postRepository *repository.PostRepository
}

func NewPostService(postRepository *repository.PostRepository) PostService {
	return &PostServiceImpl{
		postRepository: postRepository,
	}
}

func (s *PostServiceImpl) Create(p view.PostView) (view.PostView, error) {
	post := model.Post{
		Title: p.Title,
		Text:  p.Text,
	}

	id, error := (*s.postRepository).Create(post)
	p.ID = id
	return p, error
}

func (s *PostServiceImpl) FindAll() ([]view.PostView, error) {
	posts, error := (*s.postRepository).FindAll()
	return lo.Map(posts, mapToView), error
}

func mapToView(model model.Post, index int) view.PostView {
	return view.PostView{
		ID:    model.ID,
		Text:  model.Text,
		Title: model.Title,
	}
}
