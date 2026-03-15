package service

import (
	"forum_service/src/db"
	"forum_service/src/model"

	"gorm.io/gorm"
)

// PostService 帖子服务
type PostService struct {
	db *gorm.DB
}

// NewPostService 创建帖子服务实例
func NewPostService(db *gorm.DB) *PostService {
	return &PostService{
		db: db,
	}
}

// PostVo 帖子视图对象
type PostVo struct {
	Post *model.Post `json:"post"`
	User *model.User `json:"user"`
}

// PostListData 帖子列表数据
type PostListData struct {
	Records     []*PostVo `json:"records"`
	CurrentPage int       `json:"currentPage"`
	Total       int64     `json:"total"`
	PageSize    int       `json:"pageSize"`
}

// GetPostList 获取帖子列表
func (ps *PostService) GetPostList(currentPage, pageSize int) (*PostListData, error) {
	// 获取帖子列表
	posts, total, err := db.GetPostList(ps.db, currentPage, pageSize)
	if err != nil {
		return nil, err
	}

	// 构建PostVo列表
	var postVos []*PostVo
	for _, post := range posts {
		// 设置格式化的创建时间
		post.CreateTimeStr = post.CreatedAt.Format("2006-01-02 15:04:05")

		// 获取帖子作者信息
		user, err := db.GetUserByID(ps.db, post.UserID)
		if err != nil {
			continue
		}

		// 构建PostVo
		postVo := &PostVo{
			Post: post,
			User: user,
		}
		postVos = append(postVos, postVo)
	}

	// 构建返回数据
	postListData := &PostListData{
		Records:     postVos,
		CurrentPage: currentPage,
		Total:       total,
		PageSize:    pageSize,
	}

	return postListData, nil
}

// PublishPost 发布帖子
func (ps *PostService) PublishPost(title, content string, userID int64, tag int) (*model.Post, error) {
	post := &model.Post{
		Title:      title,
		Content:    content,
		UserID:     userID,
		CategoryID: int64(tag),
		Status:     1,
		Type:       0,
		LikeCount:  0,
		Views:      0,
	}
	err := db.CreatePost(ps.db, post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// GetPostDetail 获取帖子详情
func (ps *PostService) GetPostDetail(postID int64) (*PostVo, error) {
	// 获取帖子信息
	post, err := db.GetPostByID(ps.db, postID)
	if err != nil {
		return nil, err
	}

	// 设置格式化的创建时间
	post.CreateTimeStr = post.CreatedAt.Format("2006-01-02 15:04:05")

	// 获取帖子作者信息
	user, err := db.GetUserByID(ps.db, post.UserID)
	if err != nil {
		return nil, err
	}

	// 构建PostVo
	postVo := &PostVo{
		Post: post,
		User: user,
	}

	return postVo, nil
}
