package main

import (
	"forum_service/src/model"
	"forum_service/src/router"
	"forum_service/src/service"
	"log"
	"strings"

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

	// JWT中间件
	JWTMiddleware := func(c *gin.Context) {
		// 获取Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		// 解析token
		token := strings.TrimSpace(authHeader)
		claims, err := service.ParseToken(token)
		if err != nil {
			c.Next()
			return
		}

		// 将用户ID存入上下文
		c.Set("user_id", claims.UserID)
		c.Next()
	}

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
	err = db.AutoMigrate(&model.Follow{})
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}
	err = db.AutoMigrate(&model.Like{})
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}
	err = db.AutoMigrate(&model.Post{})
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}
	err = db.AutoMigrate(&model.Collection{})
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	// 4. 设置静态文件服务
	r.Static("/uploads", "./uploads")

	// 5. 设置路由
	r.Use(JWTMiddleware)
	router.SetupUserRoutes(r, db)
	router.SetupProfileRoutes(r, db)
	router.SetupUserPostRoutes(r, db)
	router.SetupCollectionRoutes(r, db)
	router.SetupPostRoutes(r, db)

	// 6. 启动服务器
	// 监听并在 0.0.0.0:8089 上启动服务
	r.Run(":8089")
}
