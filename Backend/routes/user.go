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
	group.GET("user/:id/the/follows", userController.GetOtherTheFollowed)
	group.GET("user/:id/the/fans", userController.GetOtherFans)
	group.GET("user/top/articles", userController.GetTop5Authors)
	group.GET("user/top/fans", userController.GetTop5FollowedUsers)

	group.GET("category/:id", cateController.GetCategoryInfo)
	group.GET("categories/primary", cateController.GetPrimaryCategories)
	group.GET("category/:id/subs", cateController.GetSecondaryCategories)

	group.GET("user/:id/articles", artController.GetListByUser)
	group.GET("category/:id/articles", artController.GetListByCategory)
	group.GET("article/:id", artController.GetArticleInfo)
	group.GET("articles", artController.GetArticleList)
	group.GET("title/articles", artController.GetListByTitle)
	group.GET("articles/count", artController.GetAllArticlesCount)
	group.GET("title/articles/count", artController.GetCountByTitle)
	group.GET("category/:id/articles/count", artController.GetCountByCategory)
	group.GET("user/:id/articles/count", artController.GetCountByUser)

	group.GET("article/:id/comments", commentController.GetRootCommentsByArticleID)
	group.GET("comment/:id/replies", commentController.GetRepliesByRootComment)

	group.Use(auth.JwtAuth())
	{
		group.GET("user/self/profile", userController.GetSelfProfile)
		group.PUT("user/self/profile", userController.UpdateSelfBasicInfo)
		group.POST("user/avatar", userController.UpLoadAvatar)
		group.GET("user/:id/isfollowed", userController.UserIsFollowed)
		group.POST("user/:id/follow", userController.UserFollow)
		group.GET("user/self/follows", userController.GetSelfTheFollowed)
		group.GET("user/self/fans", userController.GetSelfFans)

		group.POST("article", artController.CreateArticle)
		group.PUT("article/:id", artController.UpdateArticle)
		group.DELETE("article/:id", artController.DeleteArticle)
		group.GET("article/:id/isliked", artController.UserIsLiked)
		group.POST("article/:id/like", artController.UserLikesArticle)
		group.GET("article/:id/isstared", artController.UserIsStared)
		group.POST("article/:id/star", artController.UserStarsArticle)
		group.GET("user/star/articles", artController.GetStaredArticles)
		group.GET("user/star/articles/count", artController.GetStaredArtCount)

		group.POST("article/:id/comment", commentController.CreateCommentToArticle)
		group.POST("article/:id/comment/:cid/reply", commentController.CreateReply)
		group.DELETE("comment/:id", commentController.DeleteComment)
		group.GET("comment/:id/isliked", commentController.UserIsLiked)
		group.POST("comment/:id/like", commentController.UserLikesComment)

		group.GET("user/notification/unread", notifController.GetUnReadNotifsByReciver)
		group.GET("user/notification/read", notifController.GetReadNotifsByReciver)
		group.PUT("user/notifications/unread", notifController.MarkAsReadNotifs)
		group.DELETE("user/notifications/read", notifController.DeletReadNotifs)
	}
}
