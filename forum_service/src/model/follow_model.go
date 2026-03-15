package model

import "time"

// Follow 关注表模型
type Follow struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID     int64     `gorm:"column:user_id;not null;index" json:"user_id"`         // 关注者ID
	FolloweeID int64     `gorm:"column:followee_id;not null;index" json:"followee_id"` // 被关注者ID
	Status     int       `gorm:"column:status;type:tinyint;default:1" json:"status"`   // 1:正常 0:取消关注
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (Follow) TableName() string {
	return "follow"
}

// GetID 获取关注ID
func (f *Follow) GetID() int64 {
	return f.ID
}

// SetID 设置关注ID
func (f *Follow) SetID(id int64) {
	f.ID = id
}

// GetUserID 获取关注者ID
func (f *Follow) GetUserID() int64 {
	return f.UserID
}

// SetUserID 设置关注者ID
func (f *Follow) SetUserID(userID int64) {
	f.UserID = userID
}

// GetFolloweeID 获取被关注者ID
func (f *Follow) GetFolloweeID() int64 {
	return f.FolloweeID
}

// SetFolloweeID 设置被关注者ID
func (f *Follow) SetFolloweeID(followeeID int64) {
	f.FolloweeID = followeeID
}

// GetStatus 获取状态
func (f *Follow) GetStatus() int {
	return f.Status
}

// SetStatus 设置状态
func (f *Follow) SetStatus(status int) {
	f.Status = status
}

// GetCreatedAt 获取创建时间
func (f *Follow) GetCreatedAt() time.Time {
	return f.CreatedAt
}

// SetCreatedAt 设置创建时间
func (f *Follow) SetCreatedAt() {
	f.CreatedAt = time.Now()
}

// GetUpdatedAt 获取更新时间
func (f *Follow) GetUpdatedAt() time.Time {
	return f.UpdatedAt
}

// SetUpdatedAt 设置更新时间
func (f *Follow) SetUpdatedAt() {
	f.UpdatedAt = time.Now()
}
