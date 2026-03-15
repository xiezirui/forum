package db

import (
	"forum_service/src/model"
	"gorm.io/gorm"
)

// GetCollectionsByUserID 根据用户ID获取收藏列表
func GetCollectionsByUserID(db *gorm.DB, userID int64, currentPage, pageSize int) ([]*model.Collection, int64, error) {
	var collections []*model.Collection
	var total int64

	// 获取总数
	err := db.Model(&model.Collection{}).Where("user_id = ? AND status = 1", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (currentPage - 1) * pageSize
	err = db.Where("user_id = ? AND status = 1", userID).
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&collections).Error
	if err != nil {
		return nil, 0, err
	}

	// 手动关联查询帖子信息
	for _, collection := range collections {
		var post model.Post
		err := db.Where("id = ?", collection.PostID).First(&post).Error
		if err == nil {
			collection.Post = &post
		}
	}

	return collections, total, nil
}

// CreateCollection 创建收藏记录
func CreateCollection(db *gorm.DB, collection *model.Collection) error {
	return db.Create(collection).Error
}

// DeleteCollection 删除收藏记录
func DeleteCollection(db *gorm.DB, userID, postID int64) error {
	return db.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&model.Collection{}).Error
}

// GetCollection 获取收藏记录
func GetCollection(db *gorm.DB, userID, postID int64) (*model.Collection, error) {
	var collection model.Collection
	err := db.Where("user_id = ? AND post_id = ? AND status = 1", userID, postID).First(&collection).Error
	if err != nil {
		return nil, err
	}
	return &collection, nil
}
