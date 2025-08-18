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

func (r *PostRepositoryImpl) CountByTitle(title string) (int64, error) {
	var totalElements int64
	results := r.db.Model(&model.Post{}).Where("title = ?", title).Count(&totalElements)
	return totalElements, results.Error
}

func (r *PostRepositoryImpl) FindAll(pageNumber, pageSize int64) ([]model.Post, int64, error) {
	var post []model.Post
	offset := pageNumber * pageSize
	results := r.db.Offset(int(offset)).Limit(int(pageSize)).Find(&post)
	var totalElements int64
	r.db.Model(&model.Post{}).Count(&totalElements)
	return post, totalElements, results.Error
}
