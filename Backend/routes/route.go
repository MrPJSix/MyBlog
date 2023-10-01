package routes

import (
	"github.com/gin-gonic/gin"
	"myblog.backend/config"
	"myblog.backend/middleware/cors"
)

func InitRouter() {
	gin.SetMode(config.AppMode)

	router := gin.Default()
	router.Use(cors.Cors())

	adminRouter := router.Group("admin")
	userRouter := router.Group("")

	InitAdminRouter(adminRouter)
	InitUserRouter(userRouter)

	router.Run(config.HttpPort)
}
