package db

import (
	"forum_service/src/model"
	"gorm.io/gorm"
)

// AddUser 添加用户
func AddUser(db *gorm.DB, user *model.User) error {
	return db.Create(user).Error
}

// GetUserByUsername 根据用户名查询用户
func GetUserByUsername(db *gorm.DB, username string) (*model.User, error) {
	var user model.User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID 根据ID获取用户
func GetUserByID(db *gorm.DB, userID int64) (*model.User, error) {
	var user model.User
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
