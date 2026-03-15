package controller

import (
	"forum_service/src/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CollectionController 收藏控制器
type CollectionController struct {
	collectionService *service.CollectionService
}

// NewCollectionController 创建收藏控制器实例
func NewCollectionController(db *gorm.DB) *CollectionController {
	return &CollectionController{
		collectionService: service.NewCollectionService(db),
	}
}

// GetCollections 获取用户收藏列表
// @Summary 获取用户收藏列表
// @Tags 收藏
// @Accept json
// @Produce json
// @Param uid path int true "用户ID"
// @Param currentPage query int false "当前页码"
// @Success 200 {object} object "成功"
// @Router /collection/{uid} [get]
func (cc *CollectionController) GetCollections(c *gin.Context) {
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

	// 获取用户收藏列表
	collectionData, err := cc.collectionService.GetCollections(uid, currentPage, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取用户收藏列表失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 返回用户收藏列表
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取用户收藏列表成功",
		"data":    collectionData,
	})
}
