package main

import (
	"forum_service/src/model"
	"forum_service/src/router"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1. 创建一个默认的 Gin 引擎 (包含 Logger 和 Recovery 中间件)
	r := gin.Default()

	// 2. 添加CORS中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 2. 初始化数据库连接
	dsn := "root:123456@tcp(127.0.0.1:3306)/forum_server?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 3. 自动迁移数据库表
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	// 4. 设置路由
	router.SetupUserRoutes(r, db)

	// 5. 启动服务器
	// 监听并在 0.0.0.0:8089 上启动服务
	r.Run(":8089")
}
