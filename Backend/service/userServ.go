package service

import (
	"log"
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
	UpdateSelfBasicInfo(id uint, user *model.User) int
	DeleteUser(id uint) int
	CheckPassword(user *model.User) int
	GetUsersCount() (int64, int)
	GetAllCount() (int64, int)

	// 关注功能
	UserIsFollowed(followerID, followedID uint) (bool, int)
	followSQLToRedis(followerID uint)
	UserFollow(followerID, followedID uint) int
	GetTheFollowed(userID uint, pageSize, pageNum int) ([]model.User, int)
	GetFans(userID uint, pageSize, pageNum int) ([]model.User, int)

	//
	GetTop5Authors() ([]model.TopAuthor, int)
	GetTop5FollowedUsers() ([]model.TopFollowedUser, int)
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

// 编辑用户基础信息(仅限于用户名称、个人简介，管理员使用)
func (us *UserService) UpdateUserBasicInfo(requester *model.User, id uint, user *model.User) int {
	if !us.checkUserRight(requester, id) {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return us.userRepo.UpdateBasicInfo(id, user)
}

// 编辑用户个人信息(仅限于用户名称、个人简介，个人更新使用)
func (us *UserService) UpdateSelfBasicInfo(id uint, user *model.User) int {
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

func (us *UserService) UserIsFollowed(followerID, followedID uint) (bool, int) {
	var code int
	code = us.userRepo.UserIsFollowedRds(followerID, followedID)
	if code == errmsg.REDIS_SET_IS_MEMBER {
		return true, errmsg.SUCCESS
	} else if code == errmsg.REDIS_SET_ISNOT_MEMBER {
		return false, errmsg.SUCCESS
	} else if code == errmsg.REDIS_SET_NOT_EXISTS {
		go us.followSQLToRedis(followerID)
	} else if code == errmsg.REDIS_IS_SYNCING {
	}
	return us.userRepo.UserIsFollowedSQL(followerID, followedID)
}
func (us *UserService) followSQLToRedis(followerID uint) {
	err := us.userRepo.SaveFollowsToRedis(followerID)
	if err != nil {
		log.Println("用户关注加载到Redis出错", followerID, err)
	} else {
		log.Println("用户收藏加载到Redis成功", followerID)
	}
}
func (us *UserService) UserFollow(followerID, followedID uint) int {
	rdsCode := us.userRepo.UserIsFollowedRds(followerID, followedID)
	if rdsCode == errmsg.REDIS_SET_IS_MEMBER {
		go us.userRepo.UnFollow(followerID, followedID)
		return us.userRepo.UnFollowRds(followerID, followedID)
	} else if rdsCode == errmsg.REDIS_SET_ISNOT_MEMBER {
		go us.userRepo.Follow(followerID, followedID)
		return us.userRepo.FollowRds(followerID, followedID)
	} else if rdsCode == errmsg.REDIS_IS_SYNCING {
		followed, code := us.userRepo.UserIsFollowedSQL(followerID, followedID)
		if code != errmsg.SUCCESS {
			return code
		}
		if followed {
			go us.userRepo.UnFollowRds(followerID, followedID)
			code = us.userRepo.UnFollow(followerID, followedID)
		} else {
			go us.userRepo.FollowRds(followerID, followedID)
			code = us.userRepo.Follow(followerID, followedID)
		}
		return code
	}
	followed, code := us.UserIsFollowed(followerID, followedID)
	if code != errmsg.SUCCESS {
		return code
	}
	if followed {
		code = us.userRepo.UnFollow(followerID, followedID)
	} else {
		code = us.userRepo.Follow(followerID, followedID)
	}
	return code
}
func (us *UserService) GetTheFollowed(userID uint, pageSize, pageNum int) ([]model.User, int) {
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
	var users []model.User
	var code int
	users, code = us.userRepo.GetTheFollowedRds(userID, pageSize, offset)
	if code != errmsg.SUCCESS {
		users, code = us.userRepo.GetTheFollowed(userID, pageSize, offset)
	}
	return users, code
}
func (us *UserService) GetFans(userID uint, pageSize, pageNum int) ([]model.User, int) {
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
	return us.userRepo.GetFans(userID, pageSize, offset)
}

func (us *UserService) GetTop5Authors() ([]model.TopAuthor, int) {
	return us.userRepo.GetTop5Authors()
}

// 获取有粉丝的用户
func (us *UserService) GetTop5FollowedUsers() ([]model.TopFollowedUser, int) {
	return us.userRepo.GetTop5FollowedUsers()
}
