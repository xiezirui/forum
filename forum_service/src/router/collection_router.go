package router

import (
	"forum_service/src/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupCollectionRoutes 设置收藏相关路由
func SetupCollectionRoutes(r *gin.Engine, db *gorm.DB) {
	// 创建收藏控制器
	collectionController := controller.NewCollectionController(db)

	// 收藏路由
	r.GET("/collection/:uid", collectionController.GetCollections)
}
