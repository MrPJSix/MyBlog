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
	notifController := v1.NewNotificationController()

	group.POST("login", userController.Login)
	group.POST("register", userController.Register)
	group.GET("user/:id", userController.GetUserInfo)

	group.GET("category/:id", cateController.GetCategoryInfo)
	group.GET("categories/primary", cateController.GetPrimaryCategories)
	group.GET("category/:id/subs", cateController.GetSecondaryCategories)

	group.GET("user/:id/articles", artController.GetListByUser)
	group.GET("category/:id/articles", artController.GetListByCategory)
	group.GET("article/:id", artController.GetArticleInfo)
	group.GET("articles", artController.GetArticleList)

	group.GET("article/:id/comments", commentController.GetRootCommentsByArticleID)
	group.GET("comment/:id/replies", commentController.GetRepliesByRootComment)

	group.Use(auth.JwtAuth())
	{
		group.GET("user/self/profile", userController.GetSelfProfile)
		group.PUT("user/self/profile", userController.UpdateSelfBasicInfo)
		group.POST("user/avatar", userController.UpLoadAvatar)

		group.POST("article", artController.CreateArticle)
		group.PUT("article/:id", artController.UpdateArticle)
		group.DELETE("article/:id", artController.DeleteArticle)
		group.GET("article/:id/isliked", artController.UserIsLiked)
		group.POST("article/:id/like", artController.UserLikesArticle)

		group.POST("article/:id/comment", commentController.CreateCommentToArticle)
		group.POST("article/:id/comment/:cid/reply", commentController.CreateReply)
		group.DELETE("comment/:id", commentController.DeleteComment)

		group.GET("user/notification/unread", notifController.GetUnReadNotifsByReciver)
		group.GET("user/notification/read", notifController.GetReadNotifsByReciver)
		group.PUT("user/notifications/unread", notifController.MarkAsReadNotifs)
		group.DELETE("user/notifications/read", notifController.DeletReadNotifs)
	}
}
