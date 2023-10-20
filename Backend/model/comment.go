package model

import (
	"gorm.io/gorm"
	"log"
)

type Comment struct {
	gorm.Model
	Content         string `gorm:"type:varchar(500);not null;comment:评论内容" json:"content"`
	UserID          uint   `gorm:"comment:评论者ID" json:"user_id"`
	User            User   `json:"user"`
	ArticleID       uint   `gorm:"comment:评论所属文章ID" json:"article_id"`
	Article         Article
	RootCommentID   uint      `gorm:"index;comment:评论所属根评论ID" json:"root_comment_id"`
	RootComment     *Comment  `gorm:"foreignKey:RootCommentID" json:"root_comment"`
	ParentCommentID *uint     `gorm:"index;comment:该评论回复的评论ID" json:"parent_comment_id"`
	ParentComment   *Comment  `gorm:"foreignKey:ParentCommentID" json:"parent_comment"`
	RepliedUserID   *uint     `gorm:"comment:该评论回复的评论者ID" json:"replied_user_id"`
	RepliedUser     *User     `gorm:"foreignKey:RepliedUserID" json:"replied_user"`
	Replies         []Comment `gorm:"foreignKey:ParentCommentID;" json:"replies"`
	Likes           int       `gorm:"default:0;not null;comment:点赞数" json:"likes"'`
}

func (comment *Comment) BeforeCreate(tx *gorm.DB) error {
	log.Println("创建评论前的处理...")
	if comment.ParentCommentID != nil {
		var repliedComment Comment
		err := tx.Select("user_id", "root_comment_id").Where("id = ?", *comment.ParentCommentID).First(&repliedComment).Error
		if err != nil {
			log.Println("创建回复前出错了", err)
			return err
		}
		log.Println("查询成功", repliedComment.UserID, repliedComment.RootCommentID)
		comment.RepliedUserID = &repliedComment.UserID
		comment.RootCommentID = repliedComment.RootCommentID
	}
	return nil
}

func (comment *Comment) AfterCreate(tx *gorm.DB) error {
	err := tx.Model(&Article{}).Where("id = ?", comment.ArticleID).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).
		Error

	if comment.ParentCommentID == nil {
		comment.RootCommentID = comment.ID
		tx.Model(comment).Update("root_comment_id", comment.RootCommentID)
		var article Article
		tx.Select("user_id").Where("id = ?", comment.ArticleID).First(&article)
		notification := &Notification{
			ReceiverID: article.UserID,
			SenderID:   comment.UserID,
			ReplyID:    comment.ID,
			ArticleID:  comment.ArticleID,
			IsRead:     false,
		}
		tx.Create(notification)
	} else if comment.ParentCommentID != nil && comment.RepliedUserID != nil { // 给用户创建一条通知
		notification := &Notification{
			ReceiverID: *comment.RepliedUserID,
			SenderID:   comment.UserID,
			CommentID:  comment.ParentCommentID,
			ReplyID:    comment.ID,
			ArticleID:  comment.ArticleID,
			IsRead:     false,
		}
		tx.Create(notification)
	}

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
