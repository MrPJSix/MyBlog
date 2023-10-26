package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title        string  `gorm:"type:varchar(100);not null;comment:标题" json:"title"`
	Content      string  `gorm:"type:longtext;comment:内容" json:"content"`
	Img          *string `gorm:"type:varchar(100);comment:封面链接" json:"img"`
	CommentCount int     `gorm:"type:int;not null;default:0;comment:评论数" json:"comment_count"`
	ReadCount    int     `gorm:"type:int;not null;default:0;comment:阅读量" json:"read_count"`
	CategoryID   uint    `gorm:"not null;comment:所属分类ID" json:"category_id"`
	Category     Category
	UserID       uint `gorm:"not null;comment:作者ID" json:"user_id"`
	User         User
	Likes        int `gorm:"default:0;comment:点赞数" json:"likes"`
}

type ArtileLike struct {
	ArticleID uint `gorm:"primaryKey" json:"article_id"`
	UserID    uint `gorm:"primaryKey" json:"user_id"`
}
