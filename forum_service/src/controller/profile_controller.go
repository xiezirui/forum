package controller

import (
	"forum_service/src/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProfileController 用户资料控制器
type ProfileController struct {
	profileService *service.ProfileService
}

// NewProfileController 创建用户资料控制器实例
func NewProfileController(db *gorm.DB) *ProfileController {
	return &ProfileController{
		profileService: service.NewProfileService(db),
	}
}

// GetProfile 获取用户资料
// @Summary 获取用户资料
// @Tags 用户
// @Accept json
// @Produce json
// @Param uid path int true "用户ID"
// @Success 200 {object} object "成功"
// @Router /user/profile/{uid} [get]
func (pc *ProfileController) GetProfile(c *gin.Context) {
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

	// 获取当前登录用户ID
	var currentUserID int64 = 0
	if userID, exists := c.Get("user_id"); exists {
		currentUserID = userID.(int64)
	}

	// 获取用户资料
	profileData, err := pc.profileService.GetProfile(uid, currentUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取用户资料失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 返回用户资料
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取用户资料成功",
		"data":    profileData,
	})
}

// FollowRequest 关注请求结构
type FollowRequest struct {
	EntityType int   `json:"entityType" binding:"required"`
	EntityID   int64 `json:"entityId" binding:"required"`
}

// Follow 关注用户
// @Summary 关注用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param request body FollowRequest true "关注请求"
// @Success 200 {object} object "成功"
// @Router /follow [post]
func (pc *ProfileController) Follow(c *gin.Context) {
	var req FollowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 获取当前登录用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权，请先登录",
			"data":    nil,
		})
		return
	}

	// 关注用户
	err := pc.profileService.Follow(userID.(int64), req.EntityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "关注失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 获取更新后的粉丝数
	followerCount, err := pc.profileService.GetFollowerCount(req.EntityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取粉丝数失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 返回成功
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "关注成功",
		"data": gin.H{
			"followerCount": followerCount,
			"hasFollowed":   true,
		},
	})
}

// Unfollow 取消关注用户
// @Summary 取消关注用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param request body FollowRequest true "取消关注请求"
// @Success 200 {object} object "成功"
// @Router /unfollow [post]
func (pc *ProfileController) Unfollow(c *gin.Context) {
	var req FollowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 获取当前登录用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权，请先登录",
			"data":    nil,
		})
		return
	}

	// 取消关注用户
	err := pc.profileService.Unfollow(userID.(int64), req.EntityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "取消关注失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 获取更新后的粉丝数
	followerCount, err := pc.profileService.GetFollowerCount(req.EntityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取粉丝数失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 返回成功
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "取消关注成功",
		"data": gin.H{
			"followerCount": followerCount,
			"hasFollowed":   false,
		},
	})
}

// GetFollowerCount 获取粉丝数
func (pc *ProfileController) GetFollowerCount(userID int64) (int64, error) {
	return pc.profileService.GetFollowerCount(userID)
}
