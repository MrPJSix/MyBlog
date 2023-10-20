package model

import (
	"time"
)

type Notification struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CreateAt   time.Time `json:"create_at"`
	ReceiverID uint      `gorm:"" json:"receiver_id"`
	Receiver   User      `gorm:"foreignKey:ReceiverID" json:"receiver"`
	SenderID   uint      `gorm:"" json:"sender_id"`
	Sender     User      `gorm:"foreignKey:SenderID" json:"sender"`
	CommentID  *uint     `gorm:"" json:"comment_id"`
	Comment    Comment   `json:"comment"`
	ReplyID    uint      `gorm:"" json:"reply_id"`
	Reply      Comment   `gorm:"foreignKey:ReplyID" json:"reply"`
	ArticleID  uint      `gorm:"" json:"article_id"`
	Article    Article   `json:"article"`
	IsRead     bool      `gorm:"type:tinyint(1);not null;default:0" json:"is_read"`
}
