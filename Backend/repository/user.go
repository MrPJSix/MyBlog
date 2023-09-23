package repository

import (
	"gorm.io/gorm"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
	"myblog.backend/utils/securepw"
)

// 检查用户名是否存在
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

func CheckUserID(userID int) int {
	var user model.User
	err = db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errmsg.ERROR_USER_NOT_EXIST
		}
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 新增用户
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

// 查询用户列表
func GetUsers(pageSize, pageNum int) ([]model.User, int64) {
	var users []model.User
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// 编辑用户个人信息
func EditUser(id int, data *model.User) int {
	var user model.User
	var maps = make(map[string]interface{})
	maps["full_name"] = data.FullName
	maps["bio"] = data.Bio
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// Todo 修改用户密码

// Todo 重置用户密码

// Todo 修改用户权限

// 删除用户
func DeleteUser(id int) int {
	var user model.User
	code := CheckUserID(id)
	if code != errmsg.SUCCESS {
		return code
	}
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 登录验证
func CheckLogin(username, password string) int {
	var user model.User
	db.Where("username = ?", username).First(&user)
	if !securepw.CheckPasswordHash(user.Password, password) {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCESS
}
