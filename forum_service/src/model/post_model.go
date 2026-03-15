package model

import "time"

// Post 帖子表模型
type Post struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title      string    `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Content    string    `gorm:"column:content;type:text;not null" json:"content"`
	UserID     int64     `gorm:"column:user_id;not null;index" json:"user_id"`     // 发布者ID
	CategoryID int64     `gorm:"column:category_id;index" json:"category_id"`       // 分类ID
	Views      int64     `gorm:"column:views;default:0" json:"views"`               // 浏览量
	Status     int       `gorm:"column:status;type:tinyint;default:1" json:"status"` // 1:正常 0:禁用
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (Post) TableName() string {
	return "post"
}

// GetID 获取帖子ID
func (p *Post) GetID() int64 {
	return p.ID
}

// SetID 设置帖子ID
func (p *Post) SetID(id int64) {
	p.ID = id
}

// GetTitle 获取标题
func (p *Post) GetTitle() string {
	return p.Title
}

// SetTitle 设置标题
func (p *Post) SetTitle(title string) {
	p.Title = title
}

// GetContent 获取内容
func (p *Post) GetContent() string {
	return p.Content
}

// SetContent 设置内容
func (p *Post) SetContent(content string) {
	p.Content = content
}

// GetUserID 获取用户ID
func (p *Post) GetUserID() int64 {
	return p.UserID
}

// SetUserID 设置用户ID
func (p *Post) SetUserID(userID int64) {
	p.UserID = userID
}

// GetCategoryID 获取分类ID
func (p *Post) GetCategoryID() int64 {
	return p.CategoryID
}

// SetCategoryID 设置分类ID
func (p *Post) SetCategoryID(categoryID int64) {
	p.CategoryID = categoryID
}

// GetViews 获取浏览量
func (p *Post) GetViews() int64 {
	return p.Views
}

// SetViews 设置浏览量
func (p *Post) SetViews(views int64) {
	p.Views = views
}

// GetStatus 获取状态
func (p *Post) GetStatus() int {
	return p.Status
}

// SetStatus 设置状态
func (p *Post) SetStatus(status int) {
	p.Status = status
}

// GetCreatedAt 获取创建时间
func (p *Post) GetCreatedAt() time.Time {
	return p.CreatedAt
}

// SetCreatedAt 设置创建时间
func (p *Post) SetCreatedAt() {
	p.CreatedAt = time.Now()
}

// GetUpdatedAt 获取更新时间
func (p *Post) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

// SetUpdatedAt 设置更新时间
func (p *Post) SetUpdatedAt() {
	p.UpdatedAt = time.Now()
}
