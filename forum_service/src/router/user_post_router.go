package router

import (
	"forum_service/src/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupUserPostRoutes 设置用户帖子相关路由
func SetupUserPostRoutes(r *gin.Engine, db *gorm.DB) {
	// 创建用户帖子控制器
	userPostController := controller.NewUserPostController(db)

	// 用户帖子路由
	r.GET("/user/userPost/:uid", userPostController.GetUserPosts)
}
