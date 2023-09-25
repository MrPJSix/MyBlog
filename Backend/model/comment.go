package model

import (
	"gorm.io/gorm"
	"log"
)

type Comment struct {
	gorm.Model
	Content   string  `gorm:"type:varchar(500);not null" json:"content"`
	UserID    uint    `json:"user_id"`
	User      User    `json:"user"`
	ArticleID uint    `json:"article_id"`
	Article   Article `json:"article"`
}

func (comment *Comment) AfterCreate(tx *gorm.DB) error {
	err := tx.Model(&Article{}).Where("id = ?", comment.ArticleID).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).
		Error
	if err != nil {
		log.Println("文章评论计数添加失败！", err)
		return err
	}
	return nil
}

func (comment *Comment) AfterDelete(tx *gorm.DB) error {
	err := tx.Model(&Article{}).Where("id = ?", comment.ArticleID).
		UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).
		Error
	if err != nil {
		log.Println("文章评论数减少失败！", err)
		return err
	}
	return nil
}
