package service

import (
	"log"
	"myblog.backend/model"
	"myblog.backend/repository"
	"myblog.backend/utils/errmsg"
)

/* ====================================== */

type IArticleService interface {
	checkUserRight(requester *model.User, authorID uint) bool
	CreateArticle(article *model.Article) int
	GetArticleInfo(id uint) (*model.Article, int)
	GetArticleList(pageSize, pageNum int) ([]model.Article, int)
	GetAllArticlesCount() (int64, int)
	GetListByTitle(title string, pageSize, pageNum int) ([]model.Article, int)
	GetListByCategory(articleID uint, pageSize, pageNum int) ([]model.Article, int)
	GetArticlesCountByCategory(categoryID uint) (int64, int)
	GetListByUser(userID uint, pageSize, pageNum int) ([]model.Article, int)
	GetArticlesCountByUser(userID uint) (int64, int)
	UpdateArticle(requester *model.User, id uint, article *model.Article) int
	DeleteArticle(requester *model.User, id uint) int
	UserIsLiked(articleID, userID uint) (bool, int)
	likeSQLToRedis(articleID uint) // Deprecated: 用Redis太复杂
	UserLikesArticle(articleID, userID uint) int
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

func (as *ArticleService) GetArticleList(pageSize, pageNum int) ([]model.Article, int) {
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

func (as *ArticleService) GetAllArticlesCount() (int64, int) {
	return as.articleRepo.GetAllCount()
}

func (as *ArticleService) GetListByTitle(title string, pageSize, pageNum int) ([]model.Article, int) {
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

func (as *ArticleService) GetListByCategory(categoryID uint, pageSize, pageNum int) ([]model.Article, int) {
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

func (as *ArticleService) GetArticlesCountByCategory(categoryID uint) (int64, int) {
	return as.articleRepo.GetCountByCategory(categoryID)
}

func (as *ArticleService) GetListByUser(userID uint, pageSize, pageNum int) ([]model.Article, int) {
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

func (as *ArticleService) GetArticlesCountByUser(userID uint) (int64, int) {
	return as.articleRepo.GetCountByUser(userID)
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

func (as *ArticleService) UserIsLiked(articleID, userID uint) (bool, int) {
	//code = as.articleRepo.UserIsLikedRds(articleID, userID)
	//if code == errmsg.REDIS_SET_IS_MEMBER {
	//	return true, errmsg.SUCCESS
	//} else if code == errmsg.REDIS_SET_ISNOT_MEMBER {
	//	return false, errmsg.SUCCESS
	//} else if code == errmsg.REDIS_SET_NOT_EXISTS {
	//	go as.likeSQLToRedis(articleID)
	//}
	return as.articleRepo.UserIsLikedSQL(articleID, userID)
}

// Deprecated: 用Redis太复杂
func (as *ArticleService) likeSQLToRedis(articleID uint) {
	err := as.articleRepo.SaveLikesToRedis(articleID)
	if err != nil {
		log.Println("文章点赞加载到Redis出错", articleID, err)
	} else {
		log.Println("文章点赞加载到Redis成功", articleID)
	}
}

func (as *ArticleService) UserLikesArticle(articleID, userID uint) int {
	isLiked, code := as.UserIsLiked(articleID, userID)
	if code != errmsg.SUCCESS {
		return code
	}
	if isLiked {
		code = as.articleRepo.DecreaseLikes(articleID, userID)
	} else {
		code = as.articleRepo.IncreaseLikes(articleID, userID)
	}
	return code
}
