package repository

import (
	"gorm.io/gorm"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
)

// 新增文章
func CreateArt(data *model.Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类下的所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]model.Article, int, int64) {
	var cateArtList []model.Article
	var total int64

	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArtList).Error
	db.Model(&cateArtList).Where("cid = ?", id).Count(&total)
	if err != nil {
		return cateArtList, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}

// 查询单个文章
func GetArtInfo(id int) (model.Article, int) {
	var art model.Article
	err := db.Where("id = ?", id).Preload("Category").First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	return art, errmsg.SUCCESS
}

// 查询文章列表
func GetArt(pageSize int, pageNum int) ([]model.Article, int, int64) {
	var articleList []model.Article
	var total int64

	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at desc").Find(&articleList).Error
	db.Model(&articleList).Count(&total)
	if err != nil {
		return articleList, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// 搜索文章标题
func SearchArt(title string, pageSize int, pageNum int) ([]model.Article, int, int64) {
	var articleList []model.Article
	var total int64

	err := db.Select("article.id, title, img, created_at, updated_at, `desc`, comment_count, read_count, Category.Name").
		Order("Created_At DESC").Joins("Category").Where("title LIKE ?", "%"+title+"%").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	db.Model(&articleList).Where("title LIKE ?", title+"%").Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// 编辑文章
func EditArt(id int, data *model.Article) int {
	var art model.Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.CID
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := db.Model(&art).Where("id = ?", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func DeleteArt(id int) int {
	var art model.Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
