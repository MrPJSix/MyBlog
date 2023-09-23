package repository

import (
	"gorm.io/gorm"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
)

func CheckUserName(userName string) int {
	var user model.User
	err = db.Where("username = ?", userName).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errmsg.ERROR_USER_NOT_EXIST
		}
		return errmsg.ERROR
	}
	return errmsg.ERROR_USERNAME_USED
}

func AddUser(user *model.User) int {
	code := CheckUserName(user.Username)
	if code == errmsg.ERROR_USERNAME_USED || code == errmsg.ERROR {
		return code
	}
	err = db.Create(user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
