package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content   string  `gorm:"type:varchar(500);not null" json:"content"`
	UserID    uint    `json:"user_id"`
	User      User    `json:"user"`
	ArticleID uint    `json:"article_id"`
	Article   Article `json:"article"`
}
