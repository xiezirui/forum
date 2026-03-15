package controller

import (
	"forum_service/src/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PostController 帖子控制器
type PostController struct {
	postService *service.PostService
}

// NewPostController 创建帖子控制器实例
func NewPostController(db *gorm.DB) *PostController {
	return &PostController{
		postService: service.NewPostService(db),
	}
}

// GetPostList 获取帖子列表
// @Summary 获取帖子列表
// @Tags 帖子
// @Accept json
// @Produce json
// @Param currentPage query int false "当前页码"
// @Success 200 {object} object "成功"
// @Router /post/list [get]
func (pc *PostController) GetPostList(c *gin.Context) {
	// 获取当前页码，默认为1
	currentPage := 1
	if pageStr := c.Query("currentPage"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil && page > 0 {
			currentPage = page
		}
	}

	// 每页大小，默认为5
	pageSize := 5

	// 获取帖子列表
	postData, err := pc.postService.GetPostList(currentPage, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取帖子列表失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 返回帖子列表
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取帖子列表成功",
		"data":    postData,
	})
}

// PublishPostRequest 发布帖子请求结构
type PublishPostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Tag     int    `json:"tag"`
}

// PublishPost 发布帖子
// @Summary 发布帖子
// @Tags 帖子
// @Accept json
// @Produce json
// @Param post body PublishPostRequest true "帖子信息"
// @Success 200 {object} object "成功"
// @Router /post/publish [post]
func (pc *PostController) PublishPost(c *gin.Context) {
	var req PublishPostRequest
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

	// 从token中获取用户ID
	// 这里需要实现从token获取用户ID的逻辑
	// 暂时使用固定值，实际应该从token中解析
	userID := int64(1) // TODO: 从token中获取用户ID

	// 调用service层发布帖子
	post, err := pc.postService.PublishPost(req.Title, req.Content, userID, req.Tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "发布帖子失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 返回帖子信息
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "发布成功",
		"data":    post,
	})
}

// GetPostDetail 获取帖子详情
// @Summary 获取帖子详情
// @Tags 帖子
// @Accept json
// @Produce json
// @Param id path int true "帖子ID"
// @Success 200 {object} object "成功"
// @Router /post/detail/{id} [get]
func (pc *PostController) GetPostDetail(c *gin.Context) {
	// 获取帖子ID
	postIDStr := c.Param("id")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "帖子ID格式错误",
			"data":    nil,
		})
		return
	}

	// 调用service层获取帖子详情
	postDetail, err := pc.postService.GetPostDetail(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取帖子详情失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 返回帖子详情
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"postVo":       postDetail,
			"comments":     []interface{}{},
			"pictures":     []interface{}{},
			"likeStatus":   0,
			"collectStatus": 0,
		},
	})
}
