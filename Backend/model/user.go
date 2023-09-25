package model

import (
	"gorm.io/gorm"
	"log"
	"myblog.backend/utils/securepw"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(25);not null;unique" json:"username" validate:"required,min=8,max=25" lable:"用户名"`
	Password string `gorm:"type:varchar(100);not null" json:"password" validate:"required,min=8,max=25" lable:"密码"`
	FullName string `gorm:"type:varchar(25);not null;unique" json:"full_name"`
	Bio      string `gorm:"type:varchar(100)" json:"bio"`
	Role     uint8  `gorm:"type:tinyint;default:2" json:"role" validate:"required, gte=2" lable:"角色码"`
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
