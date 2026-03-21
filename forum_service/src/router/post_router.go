package router

import (
	"forum_service/src/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupPostRoutes 设置帖子相关路由
func SetupPostRoutes(r *gin.Engine, db *gorm.DB) {
	postController := controller.NewPostController(db)

	// 帖子路由组
	postGroup := r.Group("/post")
	{
		// 获取帖子列表
		postGroup.GET("/list", postController.GetPostList)
		// 获取帖子详情
		postGroup.GET("/detail/:id", postController.GetPostDetail)
		// 发布帖子
		postGroup.POST("/publish", postController.PublishPost)
		// 设置帖子置顶
		postGroup.GET("/top", postController.SetPostTop)
		// 设置帖子精选
		postGroup.GET("/wonderful", postController.SetPostWonderful)
	}
}
