package model

import "time"

// Event 事件表模型
type Event struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Description string    `gorm:"column:description;type:text" json:"description"`
	Location    string    `gorm:"column:location;type:varchar(255)" json:"location"`
	StartTime   time.Time `gorm:"column:start_time;not null" json:"start_time"`
	EndTime     time.Time `gorm:"column:end_time" json:"end_time"`
	UserID      int64     `gorm:"column:user_id;not null;index" json:"user_id"` // 创建者ID
	Status      int       `gorm:"column:status;type:tinyint;default:1" json:"status"` // 1:正常 0:取消
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (Event) TableName() string {
	return "event"
}

// GetID 获取事件ID
func (e *Event) GetID() int64 {
	return e.ID
}

// SetID 设置事件ID
func (e *Event) SetID(id int64) {
	e.ID = id
}

// GetTitle 获取标题
func (e *Event) GetTitle() string {
	return e.Title
}

// SetTitle 设置标题
func (e *Event) SetTitle(title string) {
	e.Title = title
}

// GetDescription 获取描述
func (e *Event) GetDescription() string {
	return e.Description
}

// SetDescription 设置描述
func (e *Event) SetDescription(description string) {
	e.Description = description
}

// GetLocation 获取地点
func (e *Event) GetLocation() string {
	return e.Location
}

// SetLocation 设置地点
func (e *Event) SetLocation(location string) {
	e.Location = location
}

// GetStartTime 获取开始时间
func (e *Event) GetStartTime() time.Time {
	return e.StartTime
}

// SetStartTime 设置开始时间
func (e *Event) SetStartTime(startTime time.Time) {
	e.StartTime = startTime
}

// GetEndTime 获取结束时间
func (e *Event) GetEndTime() time.Time {
	return e.EndTime
}

// SetEndTime 设置结束时间
func (e *Event) SetEndTime(endTime time.Time) {
	e.EndTime = endTime
}

// GetUserID 获取用户ID
func (e *Event) GetUserID() int64 {
	return e.UserID
}

// SetUserID 设置用户ID
func (e *Event) SetUserID(userID int64) {
	e.UserID = userID
}

// GetStatus 获取状态
func (e *Event) GetStatus() int {
	return e.Status
}

// SetStatus 设置状态
func (e *Event) SetStatus(status int) {
	e.Status = status
}

// GetCreatedAt 获取创建时间
func (e *Event) GetCreatedAt() time.Time {
	return e.CreatedAt
}

// SetCreatedAt 设置创建时间
func (e *Event) SetCreatedAt() {
	e.CreatedAt = time.Now()
}

// GetUpdatedAt 获取更新时间
func (e *Event) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}

// SetUpdatedAt 设置更新时间
func (e *Event) SetUpdatedAt() {
	e.UpdatedAt = time.Now()
}
