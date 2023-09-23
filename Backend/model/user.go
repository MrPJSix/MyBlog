package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(25);not null;unique" json:"username" validate:"required,min=8,max=25" lable:"用户名"`
	Password string `gorm:"type:varchar(25);not null" json:"password" validate:"required,min=8,max=25" lable:"密码"`
	FullName string `gorm:"type:varchar(15);not null;unique" json:"fullName"`
	Bio      string `gorm:"type:varchar(50)" json:"bio"`
	Role     int    `gorm:"type:int;default:2" json:"role" validate:"required, gte=2" lable:"角色码"`
	Articles []Article
}
