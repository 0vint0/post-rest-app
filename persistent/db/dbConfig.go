package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	model "vitaliesvet.com/post-rest-app/persistent/db/model/post"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("post.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error ", err)
		panic("failed to connect database")
	}
	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Post{})
}
