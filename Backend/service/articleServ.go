package service

import (
	"myblog.backend/model"
	"myblog.backend/repository"
)

/* ====================================== */

type IArticleService interface {
	CreateArticle(article *model.Article) int
	GetArticleInfo(id uint) (model.Article, int)
	GetArticleList(pageSize, pageNum int) ([]model.Article, int64, int)
	GetListByTitle(title string, pageSize, pageNum int) ([]model.Article, int64, int)
	GetListByCategory(articleID uint, pageSize, pageNum int) ([]model.Article, int64, int)
	GetListByUser(userID uint, pageSize, pageNum int) ([]model.Article, int64, int)
	UpdateArticle(id uint, article *model.Article) int
	DeleteArticle(id uint) int
}

type ArticleService struct {
	articleRepo *repository.ArticleRepo
}

func NewArticleService() *ArticleService {
	articleRepo := repository.NewArticleRepo()
	return &ArticleService{articleRepo}
}

/* ====================================== */

func (as *ArticleService) CreateArticle(article *model.Article) int {
	return as.articleRepo.Create(article)
}

func (as *ArticleService) GetArticleInfo(id uint) (model.Article, int) {
	return as.articleRepo.GetInfo(id)
}

func (as *ArticleService) GetArticleList(pageSize, pageNum int) ([]model.Article, int64, int) {
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
	return as.articleRepo.GetList(pageSize, offset)
}

func (as *ArticleService) GetListByTitle(title string, pageSize, pageNum int) ([]model.Article, int64, int) {
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
	return as.articleRepo.GetListByTitle(title, pageSize, offset)
}

func (as *ArticleService) GetListByCategory(categoryID uint, pageSize, pageNum int) ([]model.Article, int64, int) {
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
	return as.articleRepo.GetListByCategory(categoryID, pageSize, offset)
}

func (as *ArticleService) GetListByUser(userID uint, pageSize, pageNum int) ([]model.Article, int64, int) {
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
	return as.articleRepo.GetListByUser(userID, pageSize, offset)
}

func (as *ArticleService) UpdateArticle(id uint, article *model.Article) int {
	return as.articleRepo.Update(id, article)
}

func (as *ArticleService) DeleteArticle(id uint) int {
	return as.articleRepo.Delete(id)
}
