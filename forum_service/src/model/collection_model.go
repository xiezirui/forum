package model

import "time"

// Collection 收藏表模型
type Collection struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID    int64     `gorm:"column:user_id;not null;index" json:"user_id"`    // 收藏者ID
	PostID    int64     `gorm:"column:post_id;not null;index" json:"post_id"`    // 被收藏帖子ID
	Status    int       `gorm:"column:status;type:tinyint;default:1" json:"status"` // 1:收藏 0:取消收藏
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	Post      *Post     `gorm:"-" json:"post"`             // 关联的帖子
}

// TableName 指定表名
func (Collection) TableName() string {
	return "collection"
}

// GetID 获取收藏ID
func (c *Collection) GetID() int64 {
	return c.ID
}

// SetID 设置收藏ID
func (c *Collection) SetID(id int64) {
	c.ID = id
}

// GetUserID 获取收藏者ID
func (c *Collection) GetUserID() int64 {
	return c.UserID
}

// SetUserID 设置收藏者ID
func (c *Collection) SetUserID(userID int64) {
	c.UserID = userID
}

// GetPostID 获取被收藏帖子ID
func (c *Collection) GetPostID() int64 {
	return c.PostID
}

// SetPostID 设置被收藏帖子ID
func (c *Collection) SetPostID(postID int64) {
	c.PostID = postID
}

// GetStatus 获取状态
func (c *Collection) GetStatus() int {
	return c.Status
}

// SetStatus 设置状态
func (c *Collection) SetStatus(status int) {
	c.Status = status
}

// GetCreatedAt 获取创建时间
func (c *Collection) GetCreatedAt() time.Time {
	return c.CreatedAt
}

// SetCreatedAt 设置创建时间
func (c *Collection) SetCreatedAt() {
	c.CreatedAt = time.Now()
}

// GetUpdatedAt 获取更新时间
func (c *Collection) GetUpdatedAt() time.Time {
	return c.UpdatedAt
}

// SetUpdatedAt 设置更新时间
func (c *Collection) SetUpdatedAt() {
	c.UpdatedAt = time.Now()
}
