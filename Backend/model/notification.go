package model

import (
	"time"
)

type Notification struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CreateAt   time.Time `json:"create_at"`
	ReceiverID uint      `gorm:"comment:接收者ID(被回复者)" json:"receiver_id"`
	Receiver   User      `gorm:"foreignKey:ReceiverID" json:"receiver"`
	SenderID   uint      `gorm:"comment:发送者ID(回复人)" json:"sender_id"`
	Sender     User      `gorm:"foreignKey:SenderID" json:"sender"`
	CommentID  *uint     `gorm:"comment:被回复的评论ID" json:"comment_id"`
	Comment    *Comment  `json:"comment"`
	ReplyID    uint      `gorm:"comment:回复ID" json:"reply_id"`
	Reply      Comment   `gorm:"foreignKey:ReplyID" json:"reply"`
	ArticleID  uint      `gorm:"comment:文章ID" json:"article_id"`
	Article    Article   `json:"article"`
	IsRead     bool      `gorm:"type:tinyint(1);not null;default:0;comment:是否已读" json:"is_read"`
}
