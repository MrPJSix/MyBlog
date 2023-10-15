package service

import (
	"myblog.backend/model"
	"myblog.backend/repository"
	"myblog.backend/utils/errmsg"
)

/* ====================================== */

type IArticleService interface {
	checkUserRight(requester *model.User, authorID uint) bool
	CreateArticle(article *model.Article) int
	GetArticleInfo(id uint) (*model.Article, int)
	GetArticleList(pageSize, pageNum int) ([]model.Article, int64, int)
	GetListByTitle(title string, pageSize, pageNum int) ([]model.Article, int64, int)
	GetListByCategory(articleID uint, pageSize, pageNum int) ([]model.Article, int64, int)
	GetListByUser(userID uint, pageSize, pageNum int) ([]model.Article, int64, int)
	UpdateArticle(requester *model.User, id uint, article *model.Article) int
	DeleteArticle(requester *model.User, id uint) int
	GetAllArticlesCount() (int64, int)
}

type ArticleService struct {
	articleRepo *repository.ArticleRepo
}

func NewArticleService() *ArticleService {
	articleRepo := repository.NewArticleRepo()
	return &ArticleService{articleRepo}
}

/* ====================================== */

func (as *ArticleService) checkUserRight(requester *model.User, authorID uint) bool {
	if requester.Role == 1 {
		return true
	}
	if requester.ID == authorID {
		return true
	}
	return false
}

func (as *ArticleService) CreateArticle(article *model.Article) int {
	return as.articleRepo.Create(article)
}

func (as *ArticleService) GetArticleInfo(id uint) (*model.Article, int) {
	article, code := as.articleRepo.GetInfo(id)
	if code != errmsg.SUCCESS {
		return nil, code
	}
	as.articleRepo.IncreaseReadCount(id)
	return article, code
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

func (as *ArticleService) UpdateArticle(requester *model.User, id uint, article *model.Article) int {
	if !as.checkUserRight(requester, article.UserID) {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return as.articleRepo.Update(id, article)
}

func (as *ArticleService) DeleteArticle(requester *model.User, id uint) int {
	article, code := as.articleRepo.GetInfo(id)
	if code != errmsg.SUCCESS {
		return code
	}
	if !as.checkUserRight(requester, article.UserID) {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return as.articleRepo.Delete(id)
}

func (as *ArticleService) GetAllArticlesCount() (int64, int) {
	return as.articleRepo.GetAllCount()
}
