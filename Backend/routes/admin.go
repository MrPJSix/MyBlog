package routes

import (
	"github.com/gin-gonic/gin"
	v1 "myblog.backend/api/v1"
)

func InitAdminRouter(group *gin.RouterGroup) {
	// 分类模块的路由接口
	group.POST("category", v1.AddCategory)
	group.GET("category/:id", v1.GetCateInfo)
	group.GET("categories", v1.GetCate)
	group.PUT("category/:id", v1.EditCate)
	group.DELETE("category/:id", v1.DeleteCate)

	// 文章模块的路由接口
	group.POST("article", v1.AddArticle)
	group.GET("article/category/:id", v1.GetCateArt)
	group.GET("article/:id", v1.GetArtInfo)
	group.GET("articles", v1.GetArtList)
	group.PUT("article/:id", v1.EditArt)
	group.DELETE("article/:id", v1.DeleteArt)

	// 用户模块的路由接口
	group.POST("user", v1.AddUser)
	group.GET("users", v1.GetUsers)
	group.PUT("user/:id", v1.EditUser)
	group.DELETE("user/:id", v1.DeleteUser)
}
