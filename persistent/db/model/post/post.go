package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"size:255;not null"`
	Text  string `gorm:"type:text;not null"`
}
