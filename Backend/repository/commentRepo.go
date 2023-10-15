package repository

import (
	"gorm.io/gorm"
	"log"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
)

/* ====================================== */

type ICommentRepo interface {
	createAndPreload(comment *model.Comment) error
	GetByID(id uint) (*model.Comment, int)
	Create(comment *model.Comment) int
	GetByArticleID(articleID uint) ([]model.Comment, int64, int)
	Delete(id uint) int
	GetRootByArticleID(articleID uint) ([]model.Comment, int)
	GetRepliesByArticleID(articleID uint) ([]model.Comment, int)
	GetAllCount() (int64, int)
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
	db.Preload("User").Preload("RepliedUser").Where("id = ?", comment.ID).First(comment)
	return nil
}

// 检查评论是否存在
func (commentRepo *CommentRepo) GetByID(id uint) (*model.Comment, int) {
	var comment model.Comment
	err := db.Where("id = ?", id).First(&model.Comment{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errmsg.ERROR_COMMENT_NOT_EXIST
		}
		return nil, errmsg.ERROR
	}
	return &comment, errmsg.SUCCESS
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
	err := db.Preload("User").Preload("RepliedUser").
		Where("article_id = ?", articleID).
		Find(&comments).Count(&total).Error
	if err != nil {
		return nil, 0, errmsg.ERROR
	}
	return comments, total, errmsg.SUCCESS
}

// 获取谋篇文章的所有根评论
func (commentRepo *CommentRepo) GetRootByArticleID(articleID uint) ([]model.Comment, int) {
	var comments []model.Comment
	err := db.Preload("User").
		Where("article_id = ? AND parent_comment_id IS NULL", articleID).
		Order("likes DESC").
		Find(&comments).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return comments, errmsg.SUCCESS
}

// 获取谋篇文章的所有对根评论的回复
func (commentRepo *CommentRepo) GetRepliesByArticleID(articleID uint) ([]model.Comment, int) {
	var replies []model.Comment
	err := db.Preload("User").Preload("RepliedUser").
		Where("article_id = ? AND parent_comment_id IS NOT NULL", articleID).
		Find(&replies).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return replies, errmsg.SUCCESS
}

// 删除评论
func (commentRepo *CommentRepo) Delete(id uint) int {
	_, code := commentRepo.GetByID(id)
	if code != errmsg.SUCCESS {
		return code
	}
	err := db.Where("id = ?", id).Delete(&model.Comment{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func (commentRepo *CommentRepo) GetAllCount() (int64, int) {
	var total int64
	err := db.Model(&model.Comment{}).Select("id").Count(&total).Error
	if err != nil {
		log.Println("查询评论总数失败！", err)
		return 0, errmsg.ERROR
	}
	return total, errmsg.SUCCESS
}
