package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(24);not null"`
	Password string `json:"password" gorm:"type:varchar(24);not null"`
	Role     int    `json:"role" gorm:"type:int;default:2"`
}
