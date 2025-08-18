package repository

import (
	"gorm.io/gorm"
	model "vitaliesvet.com/post-rest-app/persistent/db/model/post"
)

type Repositories struct {
	PostRepository PostRepository
}

type PostRepository interface {
	FindAll(pageNumber, pageSize int64) ([]model.Post, int64, error)
	CountByTitle(title string) (int64, error)
	Create(post model.Post) (uint, error)
}

func Init(db *gorm.DB) *Repositories {
	return &Repositories{
		PostRepository: NewPostRepository(db),
	}
}
