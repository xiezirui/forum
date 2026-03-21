package db

import (
	"forum_service/src/model"
	"gorm.io/gorm"
)

// GetPostsByUserID 根据用户ID获取帖子列表
func GetPostsByUserID(db *gorm.DB, userID int64, currentPage, pageSize int) ([]*model.Post, int64, error) {
	var posts []*model.Post
	var total int64

	// 获取总数
	err := db.Model(&model.Post{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (currentPage - 1) * pageSize
	err = db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&posts).Error

	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

// GetPostByID 根据ID获取帖子
func GetPostByID(db *gorm.DB, postID int64) (*model.Post, error) {
	var post model.Post
	err := db.Where("id = ?", postID).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// GetPostList 获取帖子列表
func GetPostList(db *gorm.DB, currentPage, pageSize int) ([]*model.Post, int64, error) {
	var posts []*model.Post
	var total int64

	// 获取总数
	err := db.Model(&model.Post{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (currentPage - 1) * pageSize
	err = db.Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&posts).Error

	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

// CreatePost 创建帖子
func CreatePost(db *gorm.DB, post *model.Post) error {
	return db.Create(post).Error
}

// UpdatePostType 更新帖子类型（置顶/取消置顶）
func UpdatePostType(db *gorm.DB, postID int64, postType int) error {
	return db.Model(&model.Post{}).Where("id = ?", postID).Update("type", postType).Error
}

// UpdatePostStatus 更新帖子状态（精选/取消精选）
func UpdatePostStatus(db *gorm.DB, postID int64, status int) error {
	return db.Model(&model.Post{}).Where("id = ?", postID).Update("status", status).Error
}
