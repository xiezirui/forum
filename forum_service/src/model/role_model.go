package model

import "time"

// Role 角色表模型
type Role struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"column:name;type:varchar(50);not null;uniqueIndex" json:"name"`
	Description string    `gorm:"column:description;type:varchar(255)" json:"description"`
	Status      int       `gorm:"column:status;type:tinyint;default:1" json:"status"` // 1:正常 0:禁用
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (Role) TableName() string {
	return "role"
}

// GetID 获取角色ID
func (r *Role) GetID() int64 {
	return r.ID
}

// SetID 设置角色ID
func (r *Role) SetID(id int64) {
	r.ID = id
}

// GetName 获取角色名称
func (r *Role) GetName() string {
	return r.Name
}

// SetName 设置角色名称
func (r *Role) SetName(name string) {
	r.Name = name
}

// GetDescription 获取描述
func (r *Role) GetDescription() string {
	return r.Description
}

// SetDescription 设置描述
func (r *Role) SetDescription(description string) {
	r.Description = description
}

// GetStatus 获取状态
func (r *Role) GetStatus() int {
	return r.Status
}

// SetStatus 设置状态
func (r *Role) SetStatus(status int) {
	r.Status = status
}

// GetCreatedAt 获取创建时间
func (r *Role) GetCreatedAt() time.Time {
	return r.CreatedAt
}

// SetCreatedAt 设置创建时间
func (r *Role) SetCreatedAt() {
	r.CreatedAt = time.Now()
}

// GetUpdatedAt 获取更新时间
func (r *Role) GetUpdatedAt() time.Time {
	return r.UpdatedAt
}

// SetUpdatedAt 设置更新时间
func (r *Role) SetUpdatedAt() {
	r.UpdatedAt = time.Now()
}
