package repository

import (
	"gorm.io/gorm"
	model "vitaliesvet.com/post-rest-app/persistent/db/model/post"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{
		db: db,
	}
}

func (r *PostRepositoryImpl) Create(post model.Post) (uint, error) {
	result := r.db.Create(&post)
	return post.ID, result.Error
}

func (r *PostRepositoryImpl) FindAll() ([]model.Post, error) {
	var post []model.Post
	results := r.db.Find(&post)
	return post, results.Error
}
