package repository

import (
	"gorm.io/gorm"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
)

/* ====================================== */

type IArticleRepo interface {
	CheckByTitle(title string) int
	Create(article *model.Article) int
	GetInfo(id uint) (model.Article, int)
	GetList(pageSize, offset int) ([]model.Article, int64, int)
	GetListByTitle(title string, pageSize, offset int) ([]model.Article, int64, int)
	GetListByCategory(categoryID uint, pageSize, offset int) ([]model.Article, int64, int)
	GetListByUser(userID uint, pageSize, offset int) ([]model.Article, int64, int)
	Update(id uint, article *model.Article) int
	Delete(id uint) int
}

type ArticleRepo struct{}

func NewArticleRepo() *ArticleRepo {
	return &ArticleRepo{}
}

/* ====================================== */

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
	err = db.Create(article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询单个文章
func (ar *ArticleRepo) GetInfo(id uint) (model.Article, int) {
	var art model.Article
	err = db.Where("id = ?", id).Preload("Category").Preload("User").First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	return art, errmsg.SUCCESS
}

// 查询文章列表
func (ar *ArticleRepo) GetList(pageSize, offset int) ([]model.Article, int64, int) {
	var articleList []model.Article
	var total int64

	err = db.Preload("Category").Preload("User").
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

	err = db.Preload("Category").Preload("User").
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

	err = db.Preload("Category").Preload("User").
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

	err = db.Preload("Category").
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
	var maps = make(map[string]interface{})
	maps["title"] = article.Title
	maps["category_id"] = article.CategoryID
	maps["desc"] = article.Desc
	maps["content"] = article.Content
	maps["img"] = article.Img

	err = db.Model(&model.Article{}).Where("id = ?", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func (ar *ArticleRepo) Delete(id uint) int {
	var art model.Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}