package controller

import (
	"fmt"
	"forum_service/src/model"
	"forum_service/src/service"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserController 用户控制器
type UserController struct {
	userService *service.UserService
}

// NewUserController 创建用户控制器实例
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		userService: service.NewUserService(db),
	}
}

// RegisterRequest 用户注册请求结构
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
	Gender   int    `json:"gender"`
}

// LoginRequest 用户登录请求结构
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ChangePasswordRequest 修改密码请求结构
type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6"`
}

// Register 用户注册
// @Summary 用户注册
// @Tags 用户
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "用户注册信息"
// @Success 200 {object} object "注册成功"
// @Router /user/register [post]
func (uc *UserController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service层注册用户
	user, err := uc.userService.Register(req.Username, req.Password, req.Email, req.Gender)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "注册失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 返回用户信息（不包含密码）
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
		"data": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"gender":   user.Gender,
		},
	})
}

// AddUser 添加新用户
// @Summary 添加新用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param user body model.User true "用户信息"
// @Success 200 {object} object "成功"
// @Router /user/add [post]
func (uc *UserController) AddUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service层添加用户
	newUser, err := uc.userService.Register(user.Username, user.Password, user.Email, user.Gender)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "添加用户失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "添加用户成功",
		"data": gin.H{
			"id":       newUser.ID,
			"username": newUser.Username,
			"email":    newUser.Email,
			"gender":   newUser.Gender,
		},
	})
}

// Login 用户登录
// @Summary 用户登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param user body LoginRequest true "用户登录信息"
// @Success 200 {object} object "登录成功"
// @Router /user/login [post]
func (uc *UserController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service层登录
	user, err := uc.userService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户名或密码错误",
			"data":    nil,
		})
		return
	}

	// 检查用户状态
	if user.State != 1 {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "账户已被禁用",
			"data":    nil,
		})
		return
	}

	// 生成JWT token
	token, err := uc.userService.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "生成token失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 返回用户信息和token
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"data": gin.H{
			"token": token,
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"email":    user.Email,
				"gender":   user.Gender,
				"avatar":   user.Avatar,
				"rid":      user.RID,
			},
		},
	})
}

// GetInfo 获取用户信息
// @Summary 获取用户信息
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {object} object "成功"
// @Router /user/getInfo [post]
func (uc *UserController) GetInfo(c *gin.Context) {
	// 从请求头获取token
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权，请先登录",
			"data":    nil,
		})
		return
	}

	// 调用service层获取用户信息
	user, err := uc.userService.GetInfo(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "获取用户信息失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取用户信息成功",
		"data": gin.H{
			"id":           user.ID,
			"username":     user.Username,
			"email":        user.Email,
			"gender":       user.Gender,
			"avatar":       user.Avatar,
			"rid":          user.RID,
			"createTime":   user.CreateTime.Format("2006-01-02"),
			"createTimeStr": user.CreateTime.Format("2006-01-02"),
		},
	})
}

// ChangePassword 修改密码
// @Summary 修改密码
// @Tags 用户
// @Accept json
// @Produce json
// @Param request body ChangePasswordRequest true "修改密码请求"
// @Success 200 {object} object "修改成功"
// @Router /user/changePassword [post]
func (uc *UserController) ChangePassword(c *gin.Context) {
	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 从请求头获取token
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权，请先登录",
			"data":    nil,
		})
		return
	}

	// 调用service层获取用户信息
	user, err := uc.userService.GetInfo(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "获取用户信息失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 调用service层修改密码
	if err := uc.userService.ChangePassword(user.ID, req.OldPassword, req.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "修改密码失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 返回成功
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "修改密码成功",
		"data":    nil,
	})
}

// UploadAvatar 上传头像
// @Summary 上传头像
// @Tags 用户
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "头像文件"
// @Success 200 {object} object "上传成功"
// @Router /user/avatar [post]
func (uc *UserController) UploadAvatar(c *gin.Context) {
	// 从请求头获取token
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权，请先登录",
			"data":    nil,
		})
		return
	}

	// 调用service层获取用户信息
	user, err := uc.userService.GetInfo(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "获取用户信息失败: " + err.Error(),
			"data":    nil,
		})
		return
	}
	fmt.Printf("用户信息: ID=%d, Avatar=%s\n", user.ID, user.Avatar)

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "获取文件失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 检查文件大小（限制为1MB）
	if file.Size > 1*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "文件大小不能超过1MB",
			"data":    nil,
		})
		return
	}

	// 检查文件类型
	contentType := file.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "只支持jpg和png格式的图片",
			"data":    nil,
		})
		return
	}

	// 删除旧头像
	if user.Avatar != "" {
		oldAvatarPath := "." + user.Avatar
		fmt.Printf("旧头像路径: %s\n", oldAvatarPath)
		if _, err := os.Stat(oldAvatarPath); err == nil {
			fmt.Printf("删除旧头像: %s\n", oldAvatarPath)
			if err := os.Remove(oldAvatarPath); err != nil {
				fmt.Printf("删除旧头像失败: %v\n", err)
			}
		} else {
			fmt.Printf("旧头像文件不存在: %v\n", err)
		}
	} else {
		fmt.Println("用户没有旧头像")
	}

	// 生成文件名
	filename := fmt.Sprintf("avatar_%d_%d%s", user.ID, time.Now().Unix(), filepath.Ext(file.Filename))

	// 创建上传目录
	uploadDir := "./uploads/avatars"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建上传目录失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 保存文件
	filepath := filepath.Join(uploadDir, filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "保存文件失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 更新用户头像
	avatarURL := fmt.Sprintf("/uploads/avatars/%s", filename)
	if err := uc.userService.UpdateAvatar(user.ID, avatarURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新头像失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 返回成功
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "上传头像成功",
		"data": gin.H{
			"avatar": avatarURL,
		},
	})
}
