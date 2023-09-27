package routes

import (
	"github.com/gin-gonic/gin"
	v1 "myblog.backend/api/v1"
	"myblog.backend/middleware/auth"
)

func InitUserRouter(group *gin.RouterGroup) {
	cateController := v1.NewCategoryController()
	artController := v1.NewArticleController()
	userController := v1.NewUserController()
	commentController := v1.NewCommentController()

	group.POST("login", userController.Login)
	group.POST("register", userController.Register)
	group.GET("user/:id", userController.GetUserInfo)

	group.GET("category/:id", cateController.GetCategoryInfo)
	group.GET("categories", cateController.GetCategoryList)

	group.GET("articles/user/:id", artController.GetListByUser)
	group.GET("articles/category/:id", artController.GetListByCategory)
	group.GET("article/:id", artController.GetArticleInfo)
	group.GET("articles", artController.GetArticleList)

	group.GET("comment/article/:id", commentController.GetCommentsByArticleID)

	group.Use(auth.JwtAuth())
	{
		group.PUT("user/:id", userController.UpdateUserBasicInfo)

		group.POST("article", artController.CreateArticle)
		group.PUT("article/:id", artController.UpdateArticle)
		group.DELETE("article/:id", artController.DeleteArticle)

		group.POST("comment", commentController.CreateComment)
		group.DELETE("comment/:id", commentController.DeleteComment)
	}

}
