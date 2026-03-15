package model

import "time"

// Comment 评论表模型
type Comment struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PostID    int64     `gorm:"column:post_id;not null;index" json:"post_id"`
	UserID    int64     `gorm:"column:user_id;not null;index" json:"user_id"`
	Content   string    `gorm:"column:content;type:text;not null" json:"content"`
	ParentID  int64     `gorm:"column:parent_id;default:0" json:"parent_id"` // 父评论ID，0表示一级评论
	Status    int       `gorm:"column:status;type:tinyint;default:1" json:"status"` // 1:正常 0:禁用
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (Comment) TableName() string {
	return "comment"
}

// GetID 获取评论ID
func (c *Comment) GetID() int64 {
	return c.ID
}

// SetID 设置评论ID
func (c *Comment) SetID(id int64) {
	c.ID = id
}

// GetPostID 获取帖子ID
func (c *Comment) GetPostID() int64 {
	return c.PostID
}

// SetPostID 设置帖子ID
func (c *Comment) SetPostID(postID int64) {
	c.PostID = postID
}

// GetUserID 获取用户ID
func (c *Comment) GetUserID() int64 {
	return c.UserID
}

// SetUserID 设置用户ID
func (c *Comment) SetUserID(userID int64) {
	c.UserID = userID
}

// GetContent 获取内容
func (c *Comment) GetContent() string {
	return c.Content
}

// SetContent 设置内容
func (c *Comment) SetContent(content string) {
	c.Content = content
}

// GetParentID 获取父评论ID
func (c *Comment) GetParentID() int64 {
	return c.ParentID
}

// SetParentID 设置父评论ID
func (c *Comment) SetParentID(parentID int64) {
	c.ParentID = parentID
}

// GetStatus 获取状态
func (c *Comment) GetStatus() int {
	return c.Status
}

// SetStatus 设置状态
func (c *Comment) SetStatus(status int) {
	c.Status = status
}

// GetCreatedAt 获取创建时间
func (c *Comment) GetCreatedAt() time.Time {
	return c.CreatedAt
}

// SetCreatedAt 设置创建时间
func (c *Comment) SetCreatedAt() {
	c.CreatedAt = time.Now()
}

// GetUpdatedAt 获取更新时间
func (c *Comment) GetUpdatedAt() time.Time {
	return c.UpdatedAt
}

// SetUpdatedAt 设置更新时间
func (c *Comment) SetUpdatedAt() {
	c.UpdatedAt = time.Now()
}
