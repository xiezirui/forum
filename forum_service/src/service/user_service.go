package service

import (
	"forum_service/src/db"
	"forum_service/src/model"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService 用户服务
type UserService struct {
	db *gorm.DB
}

// NewUserService 创建用户服务实例
func NewUserService(database *gorm.DB) *UserService {
	return &UserService{db: database}
}

// Register 用户注册
func (s *UserService) Register(username, password, email string, gender int) (*model.User, error) {
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建用户对象
	user := &model.User{
		Username:   username,
		Password:   string(hashedPassword),
		Email:      email,
		Gender:     gender,
		CreateTime: time.Now(),
		State:      1, // 默认状态为正常
	}

	// 保存用户到数据库
	if err := db.AddUser(s.db, user); err != nil {
		return nil, err
	}

	return user, nil
}
