package repository

import (
	"gorm.io/gorm"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
)

/* ====================================== */

type ICommentRepo interface {
	createAndPreload(comment *model.Comment) error
	CheckByID(id uint) int
	Create(comment *model.Comment) int
	GetByArticleID(articleID uint) ([]model.Comment, int64, int)
	Delete(id uint) int
}

type CommentRepo struct {
}

func NewCommentRepo() *CommentRepo {
	return &CommentRepo{}
}

/* ====================================== */

func (cr *CommentRepo) createAndPreload(comment *model.Comment) error {
	if err := db.Create(comment).Error; err != nil {
		return err
	}
	return db.Preload("Article").Preload("User").Where("id = ?", comment.ID).First(comment).Error
}

// 检查评论是否存在
func (commentRepo *CommentRepo) CheckByID(id uint) int {
	err := db.Where("id = ?", id).First(&model.Comment{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errmsg.ERROR_COMMENT_NOT_EXIST
		}
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 新增评论
func (commentRepo *CommentRepo) Create(comment *model.Comment) int {
	err := commentRepo.createAndPreload(comment)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 获取谋篇文章的所有评论
func (commentRepo *CommentRepo) GetByArticleID(articleID uint) ([]model.Comment, int64, int) {
	var comments []model.Comment
	var total int64
	err := db.Preload("Article").Preload("User").
		Where("article_id = ?", articleID).
		Find(&comments).Count(&total).Error
	if err != nil {
		return nil, 0, errmsg.ERROR
	}
	return comments, total, errmsg.SUCCESS
}

// 删除评论
func (commentRepo *CommentRepo) Delete(id uint) int {
	code := commentRepo.CheckByID(id)
	if code != errmsg.SUCCESS {
		return code
	}
	err := db.Where("id = ?", id).Delete(&model.Comment{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
