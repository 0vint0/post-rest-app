package repository

import (
	"gorm.io/gorm"
	model "vitaliesvet.com/post-rest-app/persistent/db/model/post"
)

type Repositories struct {
	PostRepository PostRepository
}

type PostRepository interface {
	FindAll() ([]model.Post, error)
	Create(post model.Post) (uint, error)
}

func Init(db *gorm.DB) *Repositories {
	return &Repositories{
		PostRepository: NewPostRepository(db),
	}
}
