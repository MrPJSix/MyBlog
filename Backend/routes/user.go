package routes

import (
	"github.com/gin-gonic/gin"
	v1 "myblog.backend/api/v1"
)

func InitUserRouter(group *gin.RouterGroup) {
	// 分类模块的路由接口
	cateController := v1.NewCategoryController()
	group.GET("category/:id", cateController.GetCategoryInfo)
	group.GET("categories", cateController.GetCategoryList)

	// 文章模块的路由接口
	artController := v1.NewArticleController()
	group.POST("article", artController.CreateArticle)
	group.GET("articles/user/:id", artController.GetListByUser)
	group.GET("articles/category/:id", artController.GetListByCategory)
	group.GET("article/:id", artController.GetArticleInfo)
	group.GET("articles", artController.GetArticleList)
	group.PUT("article/:id", artController.UpdateSelfArticle)
	group.DELETE("article/:id", artController.DeleteSelfArticle)

	// 用户模块的路由接口
	userController := v1.NewUserController()
	group.GET("users", userController.GetUserList)
	group.PUT("user/:id", userController.UpdateUserBasicInfo)

	// 评论模块的路由接口
	commentController := v1.NewCommentController()
	group.POST("comment", commentController.CreateComment)
	group.GET("comment/article/:id", commentController.GetCommentsByArticleID)
	group.DELETE("comment/:id", commentController.DeleteComment)
}
