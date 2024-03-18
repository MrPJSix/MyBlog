package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title        string  `gorm:"type:varchar(100);not null;comment:标题" json:"title"`
	Content      string  `gorm:"type:longtext;comment:内容" json:"content"`
	ContentType  string  `gorm:"type:char(1);not null;default:'m';check:content_type IN ('h','m');comment:内容类型(HTML/Markdown)" json:"content_type"`
	Img          *string `gorm:"type:varchar(100);comment:封面链接" json:"img"`
	CommentCount int     `gorm:"type:int;not null;default:0;comment:评论数" json:"comment_count"`
	ReadCount    int     `gorm:"type:int;not null;default:0;comment:阅读量" json:"read_count"`
	CategoryID   uint    `gorm:"not null;comment:所属分类ID" json:"category_id"`
	Category     Category
	UserID       uint `gorm:"not null;comment:作者ID" json:"user_id"`
	User         User
	Likes        int `gorm:"default:0;comment:点赞数" json:"likes"`
	Stars        int `gorm:"default:0;comment:收藏数" json:"stars"`
}

type ArticleLike struct {
	ArticleID uint `gorm:"primaryKey"`
	UserID    uint `gorm:"primaryKey"`
}

func (article *Article) BeforeSave(_ *gorm.DB) error {
	if article.ContentType != "h" && article.ContentType != "m" {
		article.ContentType = "m"
	}
	return nil
}
