package service

import (
	"forum_service/src/db"
	"forum_service/src/model"
	"errors"
	"gorm.io/gorm"
)

// ProfileService 用户资料服务
type ProfileService struct {
	db *gorm.DB
}

// NewProfileService 创建用户资料服务实例
func NewProfileService(database *gorm.DB) *ProfileService {
	return &ProfileService{db: database}
}

// GetProfile 获取用户资料
func (s *ProfileService) GetProfile(userID, currentUserID int64) (*ProfileData, error) {
	// 获取用户信息
	var user model.User
	err := s.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	// 获取关注数
	followeeCount, err := db.GetFolloweeCount(s.db, userID)
	if err != nil {
		return nil, err
	}

	// 获取粉丝数
	followerCount, err := db.GetFollowerCount(s.db, userID)
	if err != nil {
		return nil, err
	}

	// 获取点赞数
	likeCount, err := db.GetUserLikeCount(s.db, userID)
	if err != nil {
		return nil, err
	}

	// 判断是否为自己
	isMine := userID == currentUserID

	// 判断是否已关注
	var hasFollowed bool
	if !isMine && currentUserID > 0 {
		_, err = db.GetFollow(s.db, currentUserID, userID)
		if err == nil {
			hasFollowed = true
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// 构造返回数据
	profileData := &ProfileData{
		User:          &user,
		IsMine:        isMine,
		FolloweeCount: followeeCount,
		FollowerCount: followerCount,
		LikeCount:     likeCount,
		HasFollowed:   hasFollowed,
		CreateTimeStr: user.CreateTime.Format("2006-01-02"),
	}

	return profileData, nil
}

// Follow 关注用户
func (s *ProfileService) Follow(userID, followeeID int64) error {
	// 检查是否已关注
	_, err := db.GetFollow(s.db, userID, followeeID)
	if err == nil {
		// 已关注，返回错误
		return errors.New("已关注该用户")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 其他错误
		return err
	}

	// 创建关注记录
	follow := &model.Follow{
		UserID:     userID,
		FolloweeID: followeeID,
		Status:     1,
	}

	return db.CreateFollow(s.db, follow)
}

// Unfollow 取消关注用户
func (s *ProfileService) Unfollow(userID, followeeID int64) error {
	// 检查是否已关注
	_, err := db.GetFollow(s.db, userID, followeeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("未关注该用户")
		}
		return err
	}

	// 删除关注记录
	return db.DeleteFollow(s.db, userID, followeeID)
}

// GetFollowerCount 获取粉丝数
func (s *ProfileService) GetFollowerCount(userID int64) (int64, error) {
	return db.GetFollowerCount(s.db, userID)
}

// ProfileData 用户资料数据结构
type ProfileData struct {
	User          *model.User `json:"user"`
	IsMine        bool        `json:"isMine"`
	FolloweeCount int64       `json:"followeeCount"`
	FollowerCount int64       `json:"followerCount"`
	LikeCount     int64       `json:"likeCount"`
	HasFollowed   bool        `json:"hasFollowed"`
	CreateTimeStr string      `json:"createTimeStr"` // 格式化的创建时间
}
