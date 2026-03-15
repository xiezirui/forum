package model

import "time"

// Message 消息表模型
type Message struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	FromID    int64     `gorm:"column:from_id;not null;index" json:"from_id"`   // 发送者ID
	ToID      int64     `gorm:"column:to_id;not null;index" json:"to_id"`       // 接收者ID
	Content   string    `gorm:"column:content;type:text;not null" json:"content"`
	Type      int       `gorm:"column:type;type:tinyint;default:1" json:"type"` // 1:私信 2:系统通知
	IsRead    bool      `gorm:"column:is_read;default:false" json:"is_read"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (Message) TableName() string {
	return "message"
}

// GetID 获取消息ID
func (m *Message) GetID() int64 {
	return m.ID
}

// SetID 设置消息ID
func (m *Message) SetID(id int64) {
	m.ID = id
}

// GetFromID 获取发送者ID
func (m *Message) GetFromID() int64 {
	return m.FromID
}

// SetFromID 设置发送者ID
func (m *Message) SetFromID(fromID int64) {
	m.FromID = fromID
}

// GetToID 获取接收者ID
func (m *Message) GetToID() int64 {
	return m.ToID
}

// SetToID 设置接收者ID
func (m *Message) SetToID(toID int64) {
	m.ToID = toID
}

// GetContent 获取内容
func (m *Message) GetContent() string {
	return m.Content
}

// SetContent 设置内容
func (m *Message) SetContent(content string) {
	m.Content = content
}

// GetType 获取消息类型
func (m *Message) GetType() int {
	return m.Type
}

// SetType 设置消息类型
func (m *Message) SetType(msgType int) {
	m.Type = msgType
}

// GetIsRead 获取是否已读
func (m *Message) GetIsRead() bool {
	return m.IsRead
}

// SetIsRead 设置是否已读
func (m *Message) SetIsRead(isRead bool) {
	m.IsRead = isRead
}

// GetCreatedAt 获取创建时间
func (m *Message) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt 设置创建时间
func (m *Message) SetCreatedAt() {
	m.CreatedAt = time.Now()
}

// GetUpdatedAt 获取更新时间
func (m *Message) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt 设置更新时间
func (m *Message) SetUpdatedAt() {
	m.UpdatedAt = time.Now()
}
