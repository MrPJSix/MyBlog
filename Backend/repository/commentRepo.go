package repository

import (
	"context"
	"gorm.io/gorm"
	"log"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
	"myblog.backend/utils/rdsprefix"
	"strconv"
	"time"
)

/* ====================================== */

type ICommentRepo interface {
	createAndPreload(comment *model.Comment) error
	GetByID(id uint) (*model.Comment, int)
	Create(comment *model.Comment) int
	GetByArticleID(articleID uint) ([]model.Comment, int64, int)
	GetRootByArticleID(articleID uint) ([]model.Comment, int)
	GetRepliesByArticleID(articleID uint) ([]model.Comment, int)
	GetRepliesByRoot(rootCommentID uint, pageSize, offset int) ([]model.Comment, int)
	Delete(id uint) int
	GetAllCount() (int64, int)

	// 点赞功能
	UserIsLikedRds(commentID, userID uint) int // Deprecated: 用Redis太复杂
	UserIsLikedSQL(commentID, userID uint) (bool, int)
	IncreaseLikesRds(commentID, uesrID uint) int
	IncreaseLikes(commentID, uesrID uint) int
	DecreaseLikesRds(commentID, userID uint) int
	DecreaseLikes(commentID, userID uint) int
	SaveLikesToRedis(commentID uint) error
}

type CommentRepo struct {
}

func NewCommentRepo() *CommentRepo {
	return &CommentRepo{}
}

/* ====================================== */

func (cr *CommentRepo) createAndPreload(comment *model.Comment) error {
	if err := db.Create(comment).Error; err != nil {
		return err
	}
	db.Preload("User").Preload("RepliedUser").Where("id = ?", comment.ID).First(comment)
	return nil
}

// 检查评论是否存在
func (commentRepo *CommentRepo) GetByID(id uint) (*model.Comment, int) {
	var comment model.Comment
	err := db.Where("id = ?", id).First(&model.Comment{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errmsg.ERROR_COMMENT_NOT_EXIST
		}
		return nil, errmsg.ERROR
	}
	return &comment, errmsg.SUCCESS
}

// 新增评论
func (commentRepo *CommentRepo) Create(comment *model.Comment) int {
	err := commentRepo.createAndPreload(comment)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 获取谋篇文章的所有评论
func (commentRepo *CommentRepo) GetByArticleID(articleID uint) ([]model.Comment, int64, int) {
	var comments []model.Comment
	var total int64
	err := db.Preload("User").Preload("RepliedUser").
		Where("article_id = ?", articleID).
		Find(&comments).Count(&total).Error
	if err != nil {
		return nil, 0, errmsg.ERROR
	}
	return comments, total, errmsg.SUCCESS
}

// 获取谋篇文章的所有根评论
func (commentRepo *CommentRepo) GetRootByArticleID(articleID uint) ([]model.Comment, int) {
	var comments []model.Comment
	err := db.Preload("User").
		Where("article_id = ? AND parent_comment_id IS NULL", articleID).
		Order("likes DESC").
		Find(&comments).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return comments, errmsg.SUCCESS
}

// 获取谋篇文章的所有对根评论的回复
func (commentRepo *CommentRepo) GetRepliesByArticleID(articleID uint) ([]model.Comment, int) {
	var replies []model.Comment
	err := db.Preload("User").Preload("RepliedUser").
		Where("article_id = ? AND parent_comment_id IS NOT NULL", articleID).
		Find(&replies).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return replies, errmsg.SUCCESS
}

// 获取根评论的所有回复
func (commentRepo *CommentRepo) GetRepliesByRoot(rootCommentID uint, pageSize, offset int) ([]model.Comment, int) {
	var replies []model.Comment
	err := db.Preload("User").Preload("RepliedUser").
		Where("root_comment_id = ? AND parent_comment_id IS NOT NULL", rootCommentID).
		Limit(pageSize).Offset(offset).
		Find(&replies).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return replies, errmsg.SUCCESS
}

// 删除评论
func (commentRepo *CommentRepo) Delete(id uint) int {
	_, code := commentRepo.GetByID(id)
	if code != errmsg.SUCCESS {
		return code
	}
	err := db.Where("id = ?", id).Delete(&model.Comment{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 获取评论总数
func (commentRepo *CommentRepo) GetAllCount() (int64, int) {
	var total int64
	err := db.Model(&model.Comment{}).Select("id").Count(&total).Error
	if err != nil {
		log.Println("查询评论总数失败！", err)
		return 0, errmsg.ERROR
	}
	return total, errmsg.SUCCESS
}

// 查看用户是否已经点赞(Redis)
func (cr *CommentRepo) UserIsLikedRds(commentID, userID uint) int {
	ctx := context.Background()
	redisKey := rdsprefix.CommentLikeSet + strconv.Itoa(int(commentID))
	redisKeySync := rdsprefix.CommentLikeSync + strconv.Itoa(int(commentID))
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
		log.Println("[Redis]查询CommentSync错误", err)
		return errmsg.REDIS_ERROR
	}
	if isSyncing == "1" {
		return errmsg.REDIS_IS_SYNCING
	}
	liked, err := rdb.SIsMember(ctx, rdsprefix.CommentLikeSet+strconv.Itoa(int(commentID)), userID).Result()
	if err != nil {
		log.Println("[Redis]无法确认member是否在set中", err)
		return errmsg.REDIS_ERROR
	}
	if !liked {
		return errmsg.REDIS_SET_ISNOT_MEMBER
	}
	return errmsg.REDIS_SET_IS_MEMBER
}

// 查看用户是否已经点赞(MySQL)
func (cr *CommentRepo) UserIsLikedSQL(commentID, userID uint) (bool, int) {
	var commentLike model.CommentLike
	err := db.Where("comment_id = ? AND user_id = ?", commentID, userID).First(&commentLike).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("[MySQL]查询出错", err)
		return false, errmsg.ERROR
	}
	if err == gorm.ErrRecordNotFound {
		return false, errmsg.SUCCESS
	}
	return true, errmsg.SUCCESS
}

func (cr *CommentRepo) IncreaseLikes(commentID, userID uint) int {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.Comment{}).Where("id = ?", commentID).
			UpdateColumn("likes", gorm.Expr("likes + ?", 1)).Error
		if err != nil {
			log.Printf("评论%d增加点赞数出错\n", commentID)
			log.Println(err)
			return err
		}
		err = tx.Create(&model.CommentLike{CommentID: commentID, UserID: userID}).Error
		if err != nil {
			log.Printf("评论%d增加点赞记录出错\n", commentID)
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

func (cr *CommentRepo) DecreaseLikes(commentID, userID uint) int {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.Comment{}).Where("id = ?", commentID).
			UpdateColumn("likes", gorm.Expr("likes - ?", 1)).Error
		if err != nil {
			log.Printf("评论%d减少点赞数出错\n", commentID)
			log.Println(err)
			return err
		}
		err = tx.Delete(&model.CommentLike{CommentID: commentID, UserID: userID}).Error
		if err != nil {
			log.Printf("评论%d减少点赞记录出错\n", commentID)
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

func (cr *CommentRepo) IncreaseLikesRds(commentID, userID uint) int {
	_, err := rdb.SAdd(context.Background(), rdsprefix.CommentLikeSet+strconv.Itoa(int(commentID)), userID).Result()
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func (cr *CommentRepo) DecreaseLikesRds(commentID, userID uint) int {
	res, err := rdb.SRem(context.Background(), rdsprefix.CommentLikeSet+strconv.Itoa(int(commentID)), userID).Result()
	if err != nil || res == 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 用Redis太复杂
func (cr *CommentRepo) SaveLikesToRedis(commentID uint) error {
	var ctx = context.Background()
	// 从数据库中查询点赞某篇评论的所有用户
	var likes []model.CommentLike
	err := db.Where("comment_id = ?", commentID).Find(&likes).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("[MySQL]查询评论%v的点赞用户出错\n", commentID)
		log.Println(err)
		return err
	}
	// 分批添加到 Redis
	redisKeySync := rdsprefix.CommentLikeSync + strconv.Itoa(int(commentID))
	batchSize := 1000 // 选择一个合适的批处理大小
	redisKey := rdsprefix.CommentLikeSet + strconv.Itoa(int(commentID))
	rdb.Set(ctx, redisKeySync, "1", 0)
	rdb.SAdd(ctx, redisKey, "0")
	for i := 0; i < len(likes); i += batchSize {
		end := i + batchSize
		if end > len(likes) {
			end = len(likes)
		}
		userIDs := make([]interface{}, end-i)
		for j := i; j < end; j++ {
			userIDs[j-i] = likes[j].UserID
		}
		_, err = rdb.SAdd(ctx, redisKey, userIDs...).Result()
		if err != nil {
			log.Printf("[Redis]添加评论%v的点赞用户到Set出错\n", commentID)
			log.Println(err)
			return err
		}
	}
	// 为Redis集合设置过期时间(一周)
	pipe := rdb.TxPipeline()
	pipe.Expire(ctx, redisKey, 7*24*time.Hour)
	pipe.Set(ctx, redisKeySync, "0", 7*24*time.Hour)
	_, err = pipe.Exec(ctx)
	if err != nil {
		log.Printf("[Redis]添加评论%v的点赞用户到Set设置过期时限时出错\n", commentID)
		return err
	}

	return nil
}
