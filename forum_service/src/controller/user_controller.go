package controller

import (
	"forum_service/src/model"
	"forum_service/src/service"
	"net/http"

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
