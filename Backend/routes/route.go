package routes

import (
	"github.com/gin-gonic/gin"
	"myblog.backend/config"
)

func InitRouter() {
	gin.SetMode(config.AppMode)

	router := gin.Default()

	adminRouter := router.Group("admin")
	userRouter := router.Group("")

	InitAdminRouter(adminRouter)
	InitUserRouter(userRouter)

	router.Run(config.HttpPort)
}
