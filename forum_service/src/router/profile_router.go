package router

import (
	"forum_service/src/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupProfileRoutes 设置用户资料相关路由
func SetupProfileRoutes(r *gin.Engine, db *gorm.DB) {
	// 创建用户资料控制器
	profileController := controller.NewProfileController(db)

	// 用户资料路由
	r.GET("/user/profile/:uid", profileController.GetProfile)

	// 关注相关路由
	r.POST("/follow", profileController.Follow)
	r.POST("/unfollow", profileController.Unfollow)
}
