package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
	CategoryID   uint   `gorm:"not null" json:"category_id"`
	Category     Category
	UserID       uint `gorm:"not null" json:"user_id"`
	User         User
}
