package db

import (
	"forum_service/src/model"
	"gorm.io/gorm"
)

// AddUser 添加用户
func AddUser(db *gorm.DB, user *model.User) error {
	return db.Create(user).Error
}
