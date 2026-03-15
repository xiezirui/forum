package db

import (
	"forum_service/src/model"
	"gorm.io/gorm"
)

// CreateLike 创建点赞记录
func CreateLike(db *gorm.DB, like *model.Like) error {
	return db.Create(like).Error
}

// DeleteLike 删除点赞记录
func DeleteLike(db *gorm.DB, userID, entityID int64, entityType int) error {
	return db.Where("user_id = ? AND entity_id = ? AND entity_type = ?", userID, entityID, entityType).Delete(&model.Like{}).Error
}

// GetLike 获取点赞记录
func GetLike(db *gorm.DB, userID, entityID int64, entityType int) (*model.Like, error) {
	var like model.Like
	err := db.Where("user_id = ? AND entity_id = ? AND entity_type = ? AND status = 1", userID, entityID, entityType).First(&like).Error
	if err != nil {
		return nil, err
	}
	return &like, nil
}

// GetLikeCount 获取点赞数
func GetLikeCount(db *gorm.DB, entityID int64, entityType int) (int64, error) {
	var count int64
	err := db.Model(&model.Like{}).Where("entity_id = ? AND entity_type = ? AND status = 1", entityID, entityType).Count(&count).Error
	return count, err
}

// GetUserLikeCount 获取用户获得的点赞总数
func GetUserLikeCount(db *gorm.DB, userID int64) (int64, error) {
	var count int64
	// 获取用户所有帖子ID
	var postIDs []int64
	err := db.Model(&model.Post{}).Where("user_id = ?", userID).Pluck("id", &postIDs).Error
	if err != nil {
		return 0, err
	}

	// 获取用户所有评论ID
	var commentIDs []int64
	err = db.Model(&model.Comment{}).Where("user_id = ?", userID).Pluck("id", &commentIDs).Error
	if err != nil {
		return 0, err
	}

	// 获取帖子点赞数
	var postLikeCount int64
	if len(postIDs) > 0 {
		err = db.Model(&model.Like{}).Where("entity_id IN ? AND entity_type = 1 AND status = 1", postIDs).Count(&postLikeCount).Error
		if err != nil {
			return 0, err
		}
	}

	// 获取评论点赞数
	var commentLikeCount int64
	if len(commentIDs) > 0 {
		err = db.Model(&model.Like{}).Where("entity_id IN ? AND entity_type = 2 AND status = 1", commentIDs).Count(&commentLikeCount).Error
		if err != nil {
			return 0, err
		}
	}

	count = postLikeCount + commentLikeCount
	return count, nil
}
