package routers

import (
	"github.com/kimcodell/SH-MH-diary/server/services"

	"github.com/gin-gonic/gin"
)

func ConnectPostRouter(router *gin.Engine) {
	postRouter := router.Group("/post")

	postRouter.GET("/", services.GetAllPosts)
	postRouter.GET("/:id", services.GetPostById)
	postRouter.POST("/", services.CreatePost)
	postRouter.PUT("/:id", services.UpdatePostById)
	postRouter.DELETE("/:id", services.DeletePostById)
	postRouter.GET("/users", services.GetUsersPosts)
}
