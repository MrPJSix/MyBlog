package repository

import (
	"gorm.io/gorm"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
	"myblog.backend/utils/securepw"
)

/* ====================================== */

type IUserRepo interface {
	CheckUsername(username string) int
	CheckUserID(id uint) int
	CheckFullName(fullName string) int
	Create(user *model.User) int
	GetList(pageSize, offset int) ([]model.User, int64, int)
	UpdateBasicInfo(id uint, user *model.User) int
	Delete(id uint) int
	CheckPassword(username, password string) int
	// Todo UpdatePassword(username, password string) int
	// Todo ResetPassword(username string) int
	// Todo UpdateRole(username string) int
}

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

/* ====================================== */

// 检查用户名是否存在
func (ur *UserRepo) CheckUsername(username string) int {
	var user model.User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errmsg.ERROR_USER_NOT_EXIST
		}
		return errmsg.ERROR
	}
	return errmsg.ERROR_USERNAME_USED
}

// 检查用户ID是否存在
func (ur *UserRepo) CheckUserID(id uint) int {
	var user model.User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errmsg.ERROR_USER_NOT_EXIST
		}
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 检查用户名称是否存在
func (ur *UserRepo) CheckFullName(fullName string) int {
	var user model.User
	err := db.Where("full_name = ?", fullName).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errmsg.SUCCESS
		}
		return errmsg.ERROR
	}
	return errmsg.ERROR_USER_FULLNAME_EXIST
}

// 新增用户
func (ur *UserRepo) Create(user *model.User) int {
	code := ur.CheckUsername(user.Username)
	if code == errmsg.ERROR_USERNAME_USED || code == errmsg.ERROR {
		return code
	}
	code = ur.CheckFullName(user.FullName)
	if code == errmsg.ERROR_USER_FULLNAME_EXIST || code == errmsg.ERROR {
		return code
	}

	err := db.Create(user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func (ur *UserRepo) GetList(pageSize, offset int) ([]model.User, int64, int) {
	var users []model.User
	var total int64
	err := db.Limit(pageSize).Offset(offset).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errmsg.ERROR
	}
	return users, total, errmsg.SUCCESS
}

// 编辑用户基础信息(仅限于用户名称、个人简介)
func (ur *UserRepo) UpdateBasicInfo(id uint, user *model.User) int {
	var maps = make(map[string]interface{})
	maps["full_name"] = user.FullName
	maps["bio"] = user.Bio
	err := db.Model(&model.User{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func (ur *UserRepo) Delete(id uint) int {
	var user model.User
	code := ur.CheckUserID(id)
	if code != errmsg.SUCCESS {
		return code
	}
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 检查密码
func (ur *UserRepo) CheckPassword(username, password string) int {
	var user model.User
	db.Where("username = ?", username).First(&user)
	if !securepw.CheckPasswordHash(user.Password, password) {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCESS
}