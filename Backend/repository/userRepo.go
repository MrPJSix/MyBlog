package repository

import (
	"context"
	"gorm.io/gorm"
	"log"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
	"myblog.backend/utils/rdsprefix"
	"myblog.backend/utils/securepw"
	"strconv"
	"strings"
	"time"
)

/* ====================================== */

type IUserRepo interface {
	CheckUsername(username string) int
	CheckUserID(id uint) int
	CheckFullName(fullName string) int
	Create(user *model.User) int
	GetByID(id uint) (*model.User, int)
	GetList(pageSize, offset int) ([]model.User, int64, int)
	UpdateBasicInfo(id uint, user *model.User) int
	Delete(id uint) int
	CheckPassword(user *model.User) int
	GetUsersCount() (int64, int)
	GetAllCount() (int64, int)

	// 关注功能
	UserIsFollowedRds(followerID, followedID uint) int // Deprecated: 用Redis太复杂
	UserIsFollowedSQL(followerID, followedID uint) (bool, int)
	FollowRds(followerID, followedID uint) int
	Follow(followerID, followedID uint) int
	UnFollowRds(followerID, followedID uint) int
	UnFollow(followerID, followedID uint) int
	SaveFollowsToRedis(followerID uint) error
	GetTheFollowed(userID uint, pageSize, offset int) ([]model.User, int)
	GetFans(userID uint, pageSize, offset int) ([]model.User, int)
	getUsersInfo(userIDs []string) ([]model.User, int)
	GetUsersByIDs([]uint) ([]model.User, int)

	// 获取发文章排行榜用户
	GetTop5Authors() ([]model.TopAuthor, int)

	// 获取有粉丝的用户
	GetTop5FollowedUsers() ([]model.TopFollowedUser, int)

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

func (ur *UserRepo) GetByID(id uint) (*model.User, int) {
	var user model.User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errmsg.ERROR_USER_NOT_EXIST
		}
		return nil, errmsg.ERROR
	}
	return &user, errmsg.SUCCESS
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
	err := db.Model(&model.User{}).Where("id = ?", id).Updates(maps).First(&user).Error
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "Duplicate entry") && strings.Contains(err.Error(), "for key 'full_name'") {
			return errmsg.ERROR_USER_FULLNAME_EXIST
		}
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func (ur *UserRepo) Delete(id uint) int {
	var user *model.User
	var code int
	user, code = ur.GetByID(id)
	if code != errmsg.SUCCESS {
		return code
	}
	err := db.Delete(user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 检查密码
func (ur *UserRepo) CheckPassword(user *model.User) int {
	if code := ur.CheckUsername(user.Username); code != errmsg.ERROR_USERNAME_USED {
		return code
	}
	inputPassword := user.Password
	db.Where("username = ?", user.Username).First(&user)
	if !securepw.CheckPasswordHash(user.Password, inputPassword) {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCESS
}

// 获取所有用户量
func (ur *UserRepo) GetUsersCount() (int64, int) {
	var total int64
	err := db.Model(&model.User{}).Select("id").Where("role <> ?", 1).Count(&total).Error
	if err != nil {
		log.Println("查询用户总数失败！", err)
		return 0, errmsg.ERROR
	}
	return total, errmsg.SUCCESS
}

// 获取用户和管理员总量
func (ur *UserRepo) GetAllCount() (int64, int) {
	var total int64
	err := db.Model(&model.User{}).Select("id").Count(&total).Error
	if err != nil {
		log.Println("查询用户和管理员总数失败！", err)
		return 0, errmsg.ERROR
	}
	return total, errmsg.SUCCESS
}

/* ------------ 关注功能 ------------ */
func (ur *UserRepo) UserIsFollowedRds(followerID, followedID uint) int {
	ctx := context.Background()
	redisKey := rdsprefix.UserFollowSet + strconv.Itoa(int(followerID))
	redisKeySync := rdsprefix.UserFollowSync + strconv.Itoa(int(followerID))
	exists, err := rdb.Exists(ctx, redisKey, redisKeySync).Result()
	if err != nil {
		log.Println("[Redis]无法确认set是否存在", err)
		return errmsg.REDIS_ERROR
	}
	if exists != 2 {
		return errmsg.REDIS_SET_NOT_EXISTS
	}
	isSyncing, err := rdb.Get(ctx, redisKeySync).Result()
	if err != nil {
		log.Println("[Redis]查询UserFollowSync错误", err)
		return errmsg.REDIS_ERROR
	}
	if isSyncing == "1" {
		return errmsg.REDIS_IS_SYNCING
	}
	followed, err := rdb.SIsMember(ctx, redisKey, followedID).Result()
	if err != nil {
		log.Println("[Redis]无法确认member是否在set中", err)
		return errmsg.REDIS_ERROR
	}
	if !followed {
		return errmsg.REDIS_SET_ISNOT_MEMBER
	}
	return errmsg.REDIS_SET_IS_MEMBER
}

func (ur *UserRepo) UserIsFollowedSQL(followerID, followedID uint) (bool, int) {
	var userFollows model.UserFollows
	err := db.Where("follower_id = ? AND followed_id = ?", followerID, followedID).
		First(&userFollows).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("[MySQL]UserIsFollowed查询错误", err)
		return false, errmsg.ERROR
	}
	if err == gorm.ErrRecordNotFound {
		return false, errmsg.SUCCESS
	}
	return true, errmsg.SUCCESS
}

func (ur *UserRepo) FollowRds(followerID, followedID uint) int {
	ctx := context.Background()
	pipe := rdb.TxPipeline()
	setKey := rdsprefix.UserFollowSet + strconv.Itoa(int(followerID))
	followListKey := rdsprefix.UserFollowList + strconv.Itoa(int(followerID))
	pipe.SAdd(ctx, setKey, followedID)
	pipe.LPush(ctx, followListKey, followedID)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func (ur *UserRepo) Follow(followerID, followedID uint) int {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.User{}).Where("id = ?", followerID).
			UpdateColumn("follows", gorm.Expr("follows + ?", 1)).Error
		if err != nil {
			log.Printf("用户%d增加关注数出错\n", followerID)
			log.Println(err)
			return err
		}
		err = tx.Model(&model.User{}).Where("id = ?", followedID).
			UpdateColumn("fans", gorm.Expr("fans + ?", 1)).Error
		if err != nil {
			log.Printf("用户%d增加粉丝数出错\n", followerID)
			log.Println(err)
			return err
		}
		err = tx.Create(&model.UserFollows{FollowerID: followerID, FollowedID: followedID}).Error
		if err != nil {
			log.Printf("%d-%d创建关注记录出错\n", followerID, followedID)
			log.Println(err)
			return err
		}
		return nil
	})
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func (ur *UserRepo) UnFollowRds(followerID, followedID uint) int {
	ctx := context.Background()
	pipe := rdb.TxPipeline()
	setKey := rdsprefix.UserFollowSet + strconv.Itoa(int(followerID))
	followListKey := rdsprefix.UserFollowList + strconv.Itoa(int(followerID))
	pipe.SRem(ctx, setKey, followedID)
	pipe.LRem(ctx, followListKey, 0, followedID)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func (ur *UserRepo) UnFollow(followerID, followedID uint) int {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.User{}).Where("id = ?", followerID).
			UpdateColumn("follows", gorm.Expr("follows - ?", 1)).Error
		if err != nil {
			log.Printf("用户%d减少关注数出错\n", followerID)
			log.Println(err)
			return err
		}
		err = tx.Model(&model.User{}).Where("id = ?", followedID).
			UpdateColumn("fans", gorm.Expr("fans - ?", 1)).Error
		if err != nil {
			log.Printf("用户%d减少粉丝数出错\n", followerID)
			log.Println(err)
			return err
		}
		err = tx.Delete(&model.UserFollows{FollowerID: followerID, FollowedID: followedID}).Error
		if err != nil {
			log.Printf("%d-%d删除关注记录出错\n", followerID, followedID)
			log.Println(err)
			return err
		}
		return nil
	})
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func (ur *UserRepo) SaveFollowsToRedis(followerID uint) error {
	ctx := context.Background()
	var follows []model.UserFollows
	err := db.Select("followed_id").
		Where("follower_id = ?", followerID).
		Order("created_at DESC").
		Find(&follows).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("[MySQL]查询用户%v的关注出错\n", followerID)
		log.Println(err)
		return err
	}
	// 分批添加到Redis
	syncKey := rdsprefix.UserFollowSync + strconv.Itoa(int(followerID))
	followSetKey := rdsprefix.UserFollowSet + strconv.Itoa(int(followerID))
	followListKey := rdsprefix.UserFollowList + strconv.Itoa(int(followerID))
	batchSize := 1000
	rdb.Set(ctx, syncKey, "1", 0)
	rdb.SAdd(ctx, followSetKey, "0")
	rdb.LPush(ctx, followListKey, "0")
	for i := 0; i < len(follows); i += batchSize {
		end := i + batchSize
		if end > len(follows) {
			end = len(follows)
		}
		followedIDs := make([]interface{}, end-i)
		for j := i; j < end; j++ {
			followedIDs[j-i] = follows[j].FollowedID
		}
		pipe := rdb.TxPipeline()
		pipe.SAdd(ctx, followSetKey, followedIDs...)
		pipe.LPush(ctx, followSetKey, followedIDs...)
		_, err = pipe.Exec(ctx)
		if err != nil {
			log.Printf("[Redis]第%d轮同步用户%d的关注到Redis失败", i/batchSize+1, followerID)
			log.Println(err)
			return err
		}
	}
	pipe := rdb.TxPipeline()
	pipe.Expire(ctx, followSetKey, 7*24*time.Hour)
	pipe.Expire(ctx, followListKey, 7*24*time.Hour)
	pipe.Set(ctx, syncKey, "0", 7*24*time.Hour)
	_, err = pipe.Exec(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (ur *UserRepo) GetTheFollowedRds(userID uint, pageSize, offset int) ([]model.User, int) {
	ctx := context.Background()
	syncKey := rdsprefix.UserFollowSync + strconv.Itoa(int(userID))
	followListKey := rdsprefix.UserFollowList + strconv.Itoa(int(userID))
	exists, err := rdb.Exists(ctx, followListKey, syncKey).Result()
	if err != nil {
		log.Println("[Redis]无法确认list是否存在", err)
		return nil, errmsg.REDIS_ERROR
	}
	if exists == 0 {
		return nil, errmsg.REDIS_LIST_NOT_EXISTS
	}
	isSyncing, err := rdb.Get(ctx, syncKey).Result()
	if err != nil {
		log.Println("[Redis]查询UserStarSync错误", err)
		return nil, errmsg.REDIS_ERROR
	}
	if isSyncing == "1" {
		return nil, errmsg.REDIS_IS_SYNCING
	}
	followedIDs, err := rdb.LRange(ctx, followListKey, int64(offset), int64(pageSize-1)).Result()
	if err != nil {
		return nil, errmsg.ERROR
	}
	return ur.getUsersInfo(followedIDs)
}

func (ur *UserRepo) GetTheFollowed(userID uint, pageSize, offset int) ([]model.User, int) {
	var followedIDs []string
	err := db.Model(model.UserFollows{}).
		Select("followed_id").
		Where("follower_id = ?", userID).
		Order("created_at DESC").
		Limit(pageSize).Offset(offset).
		Find(&followedIDs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return ur.getUsersInfo(followedIDs)
}

func (ur *UserRepo) GetFans(userID uint, pageSize, offset int) ([]model.User, int) {
	var followerIDs []string
	log.Println("[MySQL]查询粉丝数开始...")
	err := db.Model(model.UserFollows{}).
		Select("follower_id").
		Where("followed_id = ?", userID).
		Order("created_at DESC").
		Limit(pageSize).Offset(offset).
		Find(&followerIDs).Error
	log.Println("[MySQL]查询粉丝数结束...")
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return ur.getUsersInfo(followerIDs)
}
func (ur *UserRepo) getUsersInfo(userIDs []string) ([]model.User, int) {
	var users []model.User
	err := db.Where("id IN ?", userIDs).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return users, errmsg.SUCCESS
}

/* ------------ 关注功能 ------------ */

// 获取发文章排行榜用户
func (ur *UserRepo) GetTop5Authors() ([]model.TopAuthor, int) {
	var topAuthors []model.TopAuthor
	log.Println("[MySQL]查询Top10发文量作者开始...")
	sql := `
		SELECT user.id, user.full_name, user.avatar_url, COUNT(user_id) AS articles_count
		FROM user
		JOIN article a on user.id = a.user_id
		GROUP BY user_id
		HAVING articles_count > 0
		ORDER BY articles_count DESC
		LIMIT 5`
	err := db.Raw(sql).Scan(&topAuthors).Error
	log.Println("[MySQL]查询Top10发文量作者开始...")
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)
		return nil, errmsg.ERROR
	}
	return topAuthors, errmsg.SUCCESS
}

// 获取有粉丝的用户
func (ur *UserRepo) GetTop5FollowedUsers() ([]model.TopFollowedUser, int) {
	var topUsers []model.TopFollowedUser
	log.Println("[MySQL]查询高粉丝量作者开始...")
	sql := `
		SELECT id, full_name, avatar_url, fans
		FROM user
		WHERE fans > 0
		ORDER BY fans DESC
		LIMIT 5`
	err := db.Raw(sql).Scan(&topUsers).Error
	log.Println("[MySQL]查询高粉丝量作者结束...")
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)
		return nil, errmsg.ERROR
	}
	return topUsers, errmsg.SUCCESS
}

func (ur *UserRepo) GetUsersByIDs(authorsIDs []uint) ([]model.User, int) {
	return nil, 0
}
