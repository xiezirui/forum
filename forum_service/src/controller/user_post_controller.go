package controller

import (
	"forum_service/src/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserPostController 用户帖子控制器
type UserPostController struct {
	userPostService *service.UserPostService
}

// NewUserPostController 创建用户帖子控制器实例
func NewUserPostController(db *gorm.DB) *UserPostController {
	return &UserPostController{
		userPostService: service.NewUserPostService(db),
	}
}

// GetUserPosts 获取用户帖子列表
// @Summary 获取用户帖子列表
// @Tags 用户
// @Accept json
// @Produce json
// @Param uid path int true "用户ID"
// @Param currentPage query int false "当前页码"
// @Success 200 {object} object "成功"
// @Router /user/userPost/{uid} [get]
func (upc *UserPostController) GetUserPosts(c *gin.Context) {
	// 获取用户ID
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "用户ID格式错误",
			"data":    nil,
		})
		return
	}

	// 获取当前页码，默认为1
	currentPage := 1
	if pageStr := c.Query("currentPage"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil && page > 0 {
			currentPage = page
		}
	}

	// 每页大小，默认为5
	pageSize := 5

	// 获取用户帖子列表
	userPostData, err := upc.userPostService.GetUserPosts(uid, currentPage, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取用户帖子列表失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 返回用户帖子列表
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取用户帖子列表成功",
		"data":    userPostData,
	})
}
