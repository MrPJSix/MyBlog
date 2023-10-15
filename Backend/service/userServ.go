package service

import (
	"myblog.backend/model"
	"myblog.backend/repository"
	"myblog.backend/utils/errmsg"
)

/* ====================================== */

type IUserService interface {
	checkUserRight(requester *model.User, userID uint) bool
	CreateUser(user *model.User) int
	GetUserByID(id uint) (*model.User, int)
	GetUserList(pageSize, pageNum int) ([]model.User, int64, int)
	UpdateUserBasicInfo(requester *model.User, id uint, user *model.User) int
	DeleteUser(id uint) int
	CheckPassword(user *model.User) int
	GetUsersCount() (int64, int)
	GetAllCount() (int64, int)
}

// Todo UpdatePassword(username, password string) int
// Todo ResetPassword(username string) int
// Todo UpdateRole(username string) int

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService() *UserService {
	userRepo := repository.NewUserRepo()
	return &UserService{userRepo}
}

/* ====================================== */

func (us *UserService) checkUserRight(requester *model.User, userID uint) bool {
	if requester.Role == 1 {
		return true
	}
	if requester.ID == userID {
		return true
	}
	return false
}

// 新增用户
func (us *UserService) CreateUser(user *model.User) int {
	return us.userRepo.Create(user)
}

func (us *UserService) GetUserByID(id uint) (*model.User, int) {
	return us.userRepo.GetByID(id)
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
func (us *UserService) UpdateUserBasicInfo(requester *model.User, id uint, user *model.User) int {
	if !us.checkUserRight(requester, id) {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return us.userRepo.UpdateBasicInfo(id, user)
}

func (us *UserService) DeleteUser(id uint) int {
	return us.userRepo.Delete(id)
}

func (us *UserService) CheckPassword(user *model.User) int {
	return us.userRepo.CheckPassword(user)
}

func (us *UserService) GetUsersCount() (int64, int) {
	return us.userRepo.GetUsersCount()
}

func (us *UserService) GetAllCount() (int64, int) {
	return us.userRepo.GetAllCount()
}
