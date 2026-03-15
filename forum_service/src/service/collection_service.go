package service

import (
	"forum_service/src/db"
	"forum_service/src/model"
	"gorm.io/gorm"
)

// CollectionService 收藏服务
type CollectionService struct {
	db *gorm.DB
}

// NewCollectionService 创建收藏服务实例
func NewCollectionService(database *gorm.DB) *CollectionService {
	return &CollectionService{db: database}
}

// GetCollections 获取用户收藏列表
func (s *CollectionService) GetCollections(userID int64, currentPage, pageSize int) (*CollectionData, error) {
	// 获取用户信息
	var user model.User
	err := s.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	// 获取收藏列表
	collections, total, err := db.GetCollectionsByUserID(s.db, userID, currentPage, pageSize)
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	collectionData := &CollectionData{
		User:       &user,
		Posts:      collections,
		Total:      total,
		CurrentPage: currentPage,
		PageSize:   pageSize,
	}

	return collectionData, nil
}

// CollectionData 收藏数据结构
type CollectionData struct {
	User       *model.User       `json:"user"`
	Posts      []*model.Collection `json:"posts"`
	Total      int64            `json:"total"`
	CurrentPage int              `json:"currentPage"`
	PageSize   int              `json:"pageSize"`
}
