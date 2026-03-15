package router

import (
	"forum_service/src/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupUserRoutes 设置用户相关路由
func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {
	// 创建用户控制器
	userController := controller.NewUserController(db)

	// 用户路由分组
	userGroup := r.Group("/user")
	{
		// 用户注册
		userGroup.POST("/register", userController.Register)

		// 添加用户
		userGroup.POST("/add", userController.AddUser)

		// 用户登录
		userGroup.POST("/login", userController.Login)

		// 获取用户信息
		userGroup.GET("/getInfo", userController.GetInfo)

		// 上传头像
		userGroup.POST("/avatar", userController.UploadAvatar)

		// 修改密码
		userGroup.POST("/resetPass", userController.ChangePassword)
	}
}
