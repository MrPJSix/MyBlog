package service

import (
	"myblog.backend/model"
	"myblog.backend/repository"
)

/* ====================================== */

type IUserService interface {
	CreateUser(user *model.User) int
	GetUserList(pageSize, pageNum int) ([]model.User, int64, int)
	UpdateUserBasicInfo(id uint, user *model.User) int
	DeleteUser(id uint) int
	CheckPassword(username, password string) int
	// Todo UpdatePassword(username, password string) int
	// Todo ResetPassword(username string) int
	// Todo UpdateRole(username string) int
}

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService() *UserService {
	userRepo := repository.NewUserRepo()
	return &UserService{userRepo}
}

/* ====================================== */

// 新增用户
func (us *UserService) CreateUser(user *model.User) int {
	return us.userRepo.Create(user)
}

// 查询用户列表（分页）
func (us *UserService) GetUserList(pageSize, pageNum int) ([]model.User, int64, int) {
	var offset int
	if pageSize >= 100 {
		pageSize = 100
	} else if pageSize <= 0 {
		pageSize = -1
	}
	if pageNum <= 0 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	return us.userRepo.GetList(pageSize, offset)
}

// 编辑用户基础信息(仅限于用户名称、个人简介)
func (us *UserService) UpdateUserBasicInfo(id uint, user *model.User) int {
	return us.userRepo.UpdateBasicInfo(id, user)
}

func (us *UserService) DeleteUser(id uint) int {
	return us.userRepo.Delete(id)
}

func (us *UserService) CheckPassword(username, password string) int {
	return us.userRepo.CheckPassword(username, password)
}
