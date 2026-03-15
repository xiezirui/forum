package model

import "time"

// Like 点赞表模型
type Like struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID     int64     `gorm:"column:user_id;not null;index" json:"user_id"`         // 点赞者ID
	EntityID   int64     `gorm:"column:entity_id;not null;index" json:"entity_id"`       // 被点赞实体ID
	EntityType int       `gorm:"column:entity_type;not null;index" json:"entity_type"` // 被点赞实体类型(1:帖子, 2:评论)
	Status     int       `gorm:"column:status;type:tinyint;default:1" json:"status"`    // 1:点赞 0:取消点赞
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (Like) TableName() string {
	return "like"
}

// GetID 获取点赞ID
func (l *Like) GetID() int64 {
	return l.ID
}

// SetID 设置点赞ID
func (l *Like) SetID(id int64) {
	l.ID = id
}

// GetUserID 获取点赞者ID
func (l *Like) GetUserID() int64 {
	return l.UserID
}

// SetUserID 设置点赞者ID
func (l *Like) SetUserID(userID int64) {
	l.UserID = userID
}

// GetEntityID 获取被点赞实体ID
func (l *Like) GetEntityID() int64 {
	return l.EntityID
}

// SetEntityID 设置被点赞实体ID
func (l *Like) SetEntityID(entityID int64) {
	l.EntityID = entityID
}

// GetEntityType 获取被点赞实体类型
func (l *Like) GetEntityType() int {
	return l.EntityType
}

// SetEntityType 设置被点赞实体类型
func (l *Like) SetEntityType(entityType int) {
	l.EntityType = entityType
}

// GetStatus 获取状态
func (l *Like) GetStatus() int {
	return l.Status
}

// SetStatus 设置状态
func (l *Like) SetStatus(status int) {
	l.Status = status
}

// GetCreatedAt 获取创建时间
func (l *Like) GetCreatedAt() time.Time {
	return l.CreatedAt
}

// SetCreatedAt 设置创建时间
func (l *Like) SetCreatedAt() {
	l.CreatedAt = time.Now()
}

// GetUpdatedAt 获取更新时间
func (l *Like) GetUpdatedAt() time.Time {
	return l.UpdatedAt
}

// SetUpdatedAt 设置更新时间
func (l *Like) SetUpdatedAt() {
	l.UpdatedAt = time.Now()
}
