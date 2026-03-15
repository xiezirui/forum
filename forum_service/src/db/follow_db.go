package db

import (
	"forum_service/src/model"
	"gorm.io/gorm"
)

// CreateFollow 创建关注记录
func CreateFollow(db *gorm.DB, follow *model.Follow) error {
	return db.Create(follow).Error
}

// DeleteFollow 删除关注记录
func DeleteFollow(db *gorm.DB, userID, followeeID int64) error {
	return db.Where("user_id = ? AND followee_id = ?", userID, followeeID).Delete(&model.Follow{}).Error
}

// GetFollow 获取关注记录
func GetFollow(db *gorm.DB, userID, followeeID int64) (*model.Follow, error) {
	var follow model.Follow
	err := db.Where("user_id = ? AND followee_id = ? AND status = 1", userID, followeeID).First(&follow).Error
	if err != nil {
		return nil, err
	}
	return &follow, nil
}

// GetFolloweeCount 获取关注数
func GetFolloweeCount(db *gorm.DB, userID int64) (int64, error) {
	var count int64
	err := db.Model(&model.Follow{}).Where("user_id = ? AND status = 1", userID).Count(&count).Error
	return count, err
}

// GetFollowerCount 获取粉丝数
func GetFollowerCount(db *gorm.DB, userID int64) (int64, error) {
	var count int64
	err := db.Model(&model.Follow{}).Where("followee_id = ? AND status = 1", userID).Count(&count).Error
	return count, err
}
