package routes

import (
	"github.com/gin-gonic/gin"
	v1 "myblog.backend/api/v1"
	"myblog.backend/middleware/auth"
)

func InitAdminRouter(group *gin.RouterGroup) {
	// 分类模块的路由接口
	cateController := v1.NewCategoryController()
	artController := v1.NewArticleController()
	adminController := v1.NewAdminController()
	commentController := v1.NewCommentController()

	group.POST("login", adminController.AdminLogin)

	group.Use(auth.JwtAuth(), auth.AdminAuthorization())
	{
		// 分类模块的路由接口
		group.POST("category", cateController.CreatePrimaryCategory)
		group.POST("category/:id/sub", cateController.CreateSecondaryCategory)
		group.GET("category/:id", cateController.GetCategoryInfo)
		group.GET("categories/primary", cateController.GetPrimaryCategories)
		group.GET("category/:id/subs", cateController.GetSecondaryCategories)
		group.PUT("category/:id", cateController.UpdateCategory)
		group.DELETE("category/:id", cateController.DeleteCategory)

		// 文章模块的路由接口
		group.POST("article", artController.CreateArticle)
		group.GET("user/:id/articles", artController.GetListByUser)
		group.GET("category/:id/articles", artController.GetListByCategory)
		group.GET("article/:id", artController.GetArticleInfo)
		group.GET("articles", artController.GetArticleList)
		group.PUT("article/:id", artController.UpdateArticle)
		group.DELETE("article/:id", artController.DeleteArticle)
		group.GET("articles/count", artController.GetAllArticlesCount)

		// 用户模块的路由接口
		group.POST("superuser", adminController.CreateAdmin)
		group.GET("user/:id", adminController.GetUserInfo)
		group.GET("users", adminController.GetUserList)
		group.PUT("user/:id", adminController.UpdateUserBasicInfo)
		group.DELETE("user/:id", adminController.DeleteUser)
		group.GET("users/count", adminController.GetUsersCount)
		group.GET("users/all", adminController.GetAllCount)

		// 评论模块的路由接口
		group.POST("article/:id/comment", commentController.CreateCommentToArticle)
		group.GET("article/:id/comments", commentController.GetCommentsByArticleID)
		group.DELETE("comment/:id", commentController.DeleteComment)
		group.GET("comments/count", commentController.GetAllCommentsCount)
	}
}
