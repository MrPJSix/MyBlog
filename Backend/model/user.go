package model

import (
	"gorm.io/gorm"
	"log"
	"myblog.backend/utils/securepw"
)

type User struct {
	gorm.Model
	Username  string  `gorm:"type:varchar(25);not null;unique;comment:用户名/账号" json:"username"`
	Password  string  `gorm:"type:varchar(100);not null;comment:密码" json:"password"`
	FullName  string  `gorm:"type:varchar(25);unique;comment:昵称" json:"full_name"`
	Bio       string  `gorm:"type:varchar(100);comment:个人简介" json:"bio"`
	Role      uint8   `gorm:"type:tinyint;default:2;comment:角色码(1:管理员; 2:普通用户)" json:"role"`
	AvatarURL *string `gorm:"type:varchar(100);comment:头像URL" json:"avatar_url"`
}

type UserArticleStar struct {
	UserID    uint `gorm:"primaryKey" json:"user_id"`
	ArticleID uint `gorm:"primaryKey" json:"article_id"`
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
