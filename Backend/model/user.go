package model

import (
	"gorm.io/gorm"
	"log"
	"myblog.backend/utils/securepw"
	"time"
)

type User struct {
	gorm.Model
	Username  string  `gorm:"type:varchar(25);not null;unique;comment:用户名/账号" json:"username"`
	Password  string  `gorm:"type:varchar(100);not null;comment:密码" json:"password"`
	FullName  string  `gorm:"type:varchar(25);unique;comment:昵称" json:"full_name"`
	Bio       string  `gorm:"type:varchar(100);comment:个人简介" json:"bio"`
	Role      uint8   `gorm:"type:tinyint;default:2;comment:角色码(1:管理员; 2:普通用户)" json:"role"`
	AvatarURL *string `gorm:"type:varchar(100);comment:头像URL" json:"avatar_url"`
	Follows   int     `gorm:"default:2;comment:关注数" json:"follows"`
	Fans      int     `gorm:"default:2;comment:粉丝数" json:"fans"`
}

type UserArticleStar struct {
	UserID    uint `gorm:"primaryKey;comment:用户ID"`
	ArticleID uint `gorm:"primaryKey;comment:收藏的文章ID"`
	CreatedAt time.Time
}

type UserFollows struct {
	FollowerID   uint `gorm:"primaryKey;comment:关注者ID（粉丝）" `
	Follower     User `gorm:"foreignKey:FollowerID"`
	FollowedID   uint `gorm:"primaryKey;comment:被关注者ID（博主）"`
	FlllowedUser User `gorm:"foreignKey:FollowedID"`
	CreatedAt    time.Time
}

type TopAuthor struct {
	ID            uint    `json:"id"`
	FullName      string  `json:"full_name"`
	AvatarURL     *string `json:"avatar_url"`
	ArticlesCount int     `json:"articles_count"`
}

type TopFollowedUser struct {
	ID        uint    `json:"id"`
	FullName  string  `json:"full_name"`
	AvatarURL *string `json:"avatar_url"`
	Fans      int     `json:"fans"`
}

// 密码加密 & 权限控制
func (user *User) BeforeCreate(_ *gorm.DB) error {
	var err error
	user.Password, err = securepw.HashPassword(user.Password)
	if err != nil {
		log.Println("An error occured when turn password into hash pw: ", err)
		return err
	}
	return nil
}

func (user *User) BeforeDelete(tx *gorm.DB) error {
	var err error
	err = tx.Model(user).Update("full_name", nil).Error
	if err != nil {
		log.Println("An error occured when set a deleted user's full_name as null: ", err)
		return err
	}
	return err
}
