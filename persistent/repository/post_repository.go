package repository

import (
	"gorm.io/gorm"
	model "vitaliesvet.com/post-rest-app/persistent/db/model/post"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func (r *PostRepositoryImpl) FindAll() []model.Post {
	return nil
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{
		db: db,
	}
}
