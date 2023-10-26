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

type IArticleRepo interface {
	createAndPreload(article *model.Article) error
	updateAndPreload(article *model.Article) error
	CheckByTitle(title string) int
	Create(article *model.Article) int
	GetInfo(id uint) (*model.Article, int)
	GetList(pageSize, offset int) ([]model.Article, int64, int)
	GetListByTitle(title string, pageSize, offset int) ([]model.Article, int64, int)
	GetListByCategory(categoryID uint, pageSize, offset int) ([]model.Article, int64, int)
	GetListByUser(userID uint, pageSize, offset int) ([]model.Article, int64, int)
	Update(id uint, article *model.Article) int
	Delete(id uint) int
	IncreaseReadCount(id uint)
	GetAllCount() (int64, int)
	UserIsLikedRds(articleID, userID uint) int // Deprecated: 用Redis太复杂
	UserIsLikedSQL(articleID, userID uint) (bool, int)
	IncreaseLikes(articleID, uesrID uint) int
	DecreaseLikes(articleID, userID uint) int
	SaveLikesToRedis(articleID uint) error // Deprecated: 用Redis太复杂
}

type ArticleRepo struct{}

func NewArticleRepo() *ArticleRepo {
	return &ArticleRepo{}
}

/* ====================================== */

// 创建并且预加载
func (ar *ArticleRepo) createAndPreload(article *model.Article) error {
	if err := db.Create(article).Error; err != nil {
		return err
	}
	return db.Preload("Category").Preload("User").Where("id = ?", article.ID).First(article).Error
}

// 更新并且预加载
func (ar *ArticleRepo) updateAndPreload(article *model.Article) error {
	var maps = make(map[string]interface{})
	maps["title"] = article.Title
	maps["category_id"] = article.CategoryID
	maps["content"] = article.Content

	err := db.Model(&model.Article{}).Where("id = ?", article.ID).Updates(&maps).Error
	if err != nil {
		return err
	}
	return db.Preload("Category").Preload("User").Where("id = ?", article.ID).First(article).Error
}

// 检查标题是否已存在
func (ar *ArticleRepo) CheckByTitle(title string) int {
	var article model.Article
	err := db.Where("title = ?", title).First(&article).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errmsg.SUCCESS
		}
		return errmsg.ERROR
	}
	return errmsg.ERROR_ARTICLE_TITLE_EXIST
}

// 添加文章
func (ar *ArticleRepo) Create(article *model.Article) int {
	code := ar.CheckByTitle(article.Title)
	if code != errmsg.SUCCESS {
		return code
	}
	err := ar.createAndPreload(article)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询单个文章
func (ar *ArticleRepo) GetInfo(id uint) (*model.Article, int) {
	var art model.Article
	err := db.Where("id = ?", id).Preload("Category").Preload("User").First(&art).Error
	if err != nil {
		return nil, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return &art, errmsg.SUCCESS
}

// 查询文章列表
func (ar *ArticleRepo) GetList(pageSize, offset int) ([]model.Article, int64, int) {
	var articleList []model.Article
	var total int64

	err := db.Preload("Category").Preload("User").
		Order("created_at desc").
		Limit(pageSize).Offset(offset).
		Find(&articleList).Count(&total).Error
	if err != nil {
		return articleList, 0, errmsg.ERROR
	}
	return articleList, total, errmsg.SUCCESS
}

// 通过文章标题查询文章列表
func (ar *ArticleRepo) GetListByTitle(title string, pageSize, offset int) ([]model.Article, int64, int) {
	var articleList []model.Article
	var total int64

	err := db.Preload("Category").Preload("User").
		Order("created_at DESC").
		Where("title LIKE ?", "%"+title+"%").
		Limit(pageSize).Offset(offset).
		Find(&articleList).Count(&total).Error
	if err != nil {
		return nil, 0, errmsg.ERROR
	}
	return articleList, total, errmsg.SUCCESS
}

// 通过分类名查询文章列表
func (ar *ArticleRepo) GetListByCategory(categoryID uint, pageSize, offset int) ([]model.Article, int64, int) {
	var cateArtList []model.Article
	var total int64

	var category model.Category
	db.Preload("SubCategories").Where("id = ?", categoryID).First(&category)
	if category.ParentID != nil {
		err := db.Preload("Category").Preload("User").
			Order("created_at desc").
			Limit(pageSize).Offset(offset).
			Where("category_id = ?", categoryID).
			Find(&cateArtList).Count(&total).Error
		if err != nil {
			return cateArtList, 0, errmsg.ERROR_CATE_NOT_EXIST
		}
		return cateArtList, total, errmsg.SUCCESS
	} else {
		var cids []int
		for _, sub := range category.SubCategories {
			cids = append(cids, sub.ID)
		}
		err := db.Preload("Category").Preload("User").
			Order("created_at desc").
			Limit(pageSize).Offset(offset).
			Where("category_id IN ?", cids).
			Find(&cateArtList).Count(&total).Error
		if err != nil {
			return cateArtList, 0, errmsg.ERROR_CATE_NOT_EXIST
		}
		return cateArtList, total, errmsg.SUCCESS
	}
}

// 通过用户查询文章列表
func (ar *ArticleRepo) GetListByUser(userID uint, pageSize, offset int) ([]model.Article, int64, int) {
	var articles []model.Article
	var total int64

	err := db.Preload("Category").
		Order("created_at desc").
		Limit(pageSize).Offset(offset).
		Where("user_id = ?", userID).
		Find(&articles).Count(&total).Error
	if err != nil {
		return nil, 0, errmsg.ERROR
	}
	return articles, total, errmsg.SUCCESS
}

// 编辑文章
func (ar *ArticleRepo) Update(id uint, article *model.Article) int {
	article.ID = id
	err := ar.updateAndPreload(article)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func (ar *ArticleRepo) Delete(id uint) int {
	var art model.Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 增加浏览量
func (ar *ArticleRepo) IncreaseReadCount(id uint) {
	db.Model(&model.Article{}).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
}

// 获取所有文章总量
func (ar *ArticleRepo) GetAllCount() (int64, int) {
	var total int64
	err := db.Model(&model.Article{}).Select("id").Count(&total).Error
	if err != nil {
		log.Println("查询文章总数失败！", err)
		return 0, errmsg.ERROR
	}
	return total, errmsg.SUCCESS
}

// Deprecated: 用Redis太复杂
// 查看用户是否已经点赞(Redis)
func (ar *ArticleRepo) UserIsLikedRds(articleID, userID uint) int {
	exists, err := rdb.Exists(context.Background(), rdsprefix.ArticleLikeSet+strconv.Itoa(int(articleID))).Result()
	if err != nil {
		log.Println("[Redis]无法确认set是否存在", err)
		return errmsg.REDIS_ERROR
	}
	if exists != 1 {
		return errmsg.REDIS_SET_NOT_EXISTS
	}
	liked, err := rdb.SIsMember(context.Background(), rdsprefix.ArticleLikeSet+strconv.Itoa(int(articleID)), userID).Result()
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
func (ar *ArticleRepo) UserIsLikedSQL(articleID, userID uint) (bool, int) {
	var articleLike model.ArticleLike
	err := db.Where("article_id = ? AND user_id = ?", articleID, userID).First(&articleLike).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("[MySQL]查询出错", err)
		return false, errmsg.ERROR
	}
	if err == gorm.ErrRecordNotFound {
		return false, errmsg.SUCCESS
	}
	return true, errmsg.SUCCESS
}

func (ar *ArticleRepo) IncreaseLikes(articleID, userID uint) int {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.Article{}).Where("id = ?", articleID).
			UpdateColumn("likes", gorm.Expr("likes + ?", 1)).Error
		if err != nil {
			log.Printf("文章%d增加点赞数出错\n", articleID, err)
			return err
		}
		err = tx.Create(&model.ArticleLike{ArticleID: articleID, UserID: userID}).Error
		if err != nil {
			log.Printf("文章%d增加点赞记录出错\n", articleID, err)
			return err
		}
		return nil
	})
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func (ar *ArticleRepo) DecreaseLikes(articleID, userID uint) int {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.Article{}).Where("id = ?", articleID).
			UpdateColumn("likes", gorm.Expr("likes - ?", 1)).Error
		if err != nil {
			log.Printf("文章%d减少点赞数出错\n", articleID, err)
			return err
		}
		err = tx.Delete(&model.ArticleLike{ArticleID: articleID, UserID: userID}).Error
		if err != nil {
			log.Printf("文章%d减少点赞记录出错\n", articleID, err)
			return err
		}
		return nil
	})
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// Deprecated: 用Redis太复杂
func (ar *ArticleRepo) SaveLikesToRedis(articleID uint) error {
	var ctx = context.Background()
	// 从数据库中查询点赞某篇文章的所有用户
	var likes []model.ArticleLike
	err := db.Where("article_id = ?", articleID).Find(&likes).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("[MySQL]查询文章%v的点赞用户出错\n", articleID)
		log.Println(err)
		return err
	}
	// 分批添加到 Redis
	batchSize := 1000 // 选择一个合适的批处理大小
	redisKey := rdsprefix.ArticleLikeSet + strconv.Itoa(int(articleID))
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
			log.Printf("[Redis]添加文章%v的点赞用户到Set出错\n", articleID)
			log.Println(err)
			return err
		}
	}
	// 为Redis集合设置过期时间(一周)
	_, err = rdb.Expire(ctx, redisKey, 7*24*time.Hour).Result()
	if err != nil {
		log.Println("[Redis]添加文章%v的点赞用户到Set出错\n")
		return err
	}

	return nil
}
