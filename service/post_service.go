package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
	"vitaliesvet.com/post-rest-app/c_validator"
	"vitaliesvet.com/post-rest-app/error_handler"
	"vitaliesvet.com/post-rest-app/mapper"
	model "vitaliesvet.com/post-rest-app/persistent/db/model/post"
	"vitaliesvet.com/post-rest-app/persistent/repository"
	"vitaliesvet.com/post-rest-app/rest/view"
)

type PostServiceImpl struct {
	validator      *validator.Validate
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) PostService[view.PostView] {
	v := validator.New()
	c_validator.RegisterCustomValidators(v, postRepository)
	return &PostServiceImpl{
		validator:      v,
		postRepository: postRepository,
	}
}

func (s *PostServiceImpl) Create(p view.PostView) (view.PostView, error) {
	if err := s.validator.Struct(p); err != nil {
		if verr := error_handler.NewValidationError(err); verr != nil {
			return view.PostView{}, verr // return typed validation error
		}
		return view.PostView{}, err
	}

	post := model.Post{
		Title: p.Title,
		Text:  p.Text,
	}

	id, error := s.postRepository.Create(post)
	p.ID = id
	return p, error
}

func (s *PostServiceImpl) FindAll(p view.PagedRequest) (view.Paginated[view.PostView], error) {
	posts, totalElements, error := s.postRepository.FindAll(p.PageNumber, p.PageSize)
	postViews := lo.Map(posts, mapper.MapToPostView)
	return view.Paginated[view.PostView]{
		Items:      postViews,
		PageNumber: p.PageNumber,
		PageSize:   p.PageSize,
		TotalSize:  totalElements,
	}, error
}
