package repository

import (
	"gorm.io/gorm"
	"log"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
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

	err := db.Preload("Category").Preload("User").
		Limit(pageSize).Offset(offset).
		Where("category_id = ?", categoryID).
		Find(&cateArtList).Count(&total).Error
	if err != nil {
		return cateArtList, 0, errmsg.ERROR_CATE_NOT_EXIST
	}
	return cateArtList, total, errmsg.SUCCESS
}

// 通过用户查询文章列表
func (ar *ArticleRepo) GetListByUser(userID uint, pageSize, offset int) ([]model.Article, int64, int) {
	var articles []model.Article
	var total int64

	err := db.Preload("Category").
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
