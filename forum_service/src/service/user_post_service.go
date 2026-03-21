package service

import (
	"forum_service/src/db"
	"forum_service/src/model"
	"gorm.io/gorm"
)

// UserPostService 用户帖子服务
type UserPostService struct {
	db *gorm.DB
}

// NewUserPostService 创建用户帖子服务实例
func NewUserPostService(database *gorm.DB) *UserPostService {
	return &UserPostService{db: database}
}

// GetUserPosts 获取用户帖子列表
func (s *UserPostService) GetUserPosts(userID int64, currentPage, pageSize int) (*UserPostData, error) {
	// 获取用户信息
	var user model.User
	err := s.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	// 获取帖子列表
	posts, total, err := db.GetPostsByUserID(s.db, userID, currentPage, pageSize)
	if err != nil {
		return nil, err
	}

	// 构建 PostVo 列表
	var postVos []*PostVo
	for _, post := range posts {
		// 设置格式化的创建时间
		post.CreateTimeStr = post.CreatedAt.Format("2006-01-02 15:04:05")

		// 获取帖子作者信息
		user, err := db.GetUserByID(s.db, post.UserID)
		if err != nil {
			continue
		}

		// 构建 PostVo
		postVo := &PostVo{
			Post: post,
			User: user,
		}
		postVos = append(postVos, postVo)
	}

	// 构造返回数据
	userPostData := &UserPostData{
		User: &user,
		Pagination: &Pagination{
			CurrentPage: currentPage,
			PageSize:    pageSize,
			Total:       total,
			Records:     postVos,
		},
	}

	return userPostData, nil
}

// UserPostData 用户帖子数据结构
type UserPostData struct {
	User       *model.User   `json:"user"`
	Pagination *Pagination   `json:"pagination"`
}

// Pagination 分页数据结构
type Pagination struct {
	CurrentPage int       `json:"currentPage"`
	PageSize    int       `json:"pageSize"`
	Total       int64     `json:"total"`
	Records     []*PostVo `json:"records"`
}
