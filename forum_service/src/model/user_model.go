package model

import "time"

// User 用户表模型
type User struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username   string    `gorm:"column:username;type:varchar(20);not null" json:"username"`
	Password   string    `gorm:"column:password;type:varchar(60);not null" json:"-"` // 密码不返回给前端
	Email      string    `gorm:"column:email;type:varchar(60);not null" json:"email"`
	Gender     int       `gorm:"column:gender;type:int(11);not null" json:"gender"`
	CreateTime time.Time `gorm:"column:create_time;type:date;not null" json:"createTime"`
	Avatar     string    `gorm:"column:avatar;type:varchar(100)" json:"avatar"`
	Code       string    `gorm:"column:code;type:varchar(50)" json:"code"`
	State      int       `gorm:"column:state;type:int(1)" json:"state"` // 状态
	RID        int       `gorm:"column:rid;type:int(1);not null;default:1" json:"rid"` // 角色ID
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}

// GetID 获取用户ID
func (u *User) GetID() int64 {
	return u.ID
}

// SetID 设置用户ID
func (u *User) SetID(id int64) {
	u.ID = id
}

// GetUsername 获取用户名
func (u *User) GetUsername() string {
	return u.Username
}

// SetUsername 设置用户名
func (u *User) SetUsername(username string) {
	u.Username = username
}

// GetPassword 获取密码
func (u *User) GetPassword() string {
	return u.Password
}

// SetPassword 设置密码
func (u *User) SetPassword(password string) {
	u.Password = password
}

// GetEmail 获取邮箱
func (u *User) GetEmail() string {
	return u.Email
}

// SetEmail 设置邮箱
func (u *User) SetEmail(email string) {
	u.Email = email
}

// GetGender 获取性别
func (u *User) GetGender() int {
	return u.Gender
}

// SetGender 设置性别
func (u *User) SetGender(gender int) {
	u.Gender = gender
}

// GetCreateTime 获取创建时间
func (u *User) GetCreateTime() time.Time {
	return u.CreateTime
}

// SetCreateTime 设置创建时间
func (u *User) SetCreateTime(createTime time.Time) {
	u.CreateTime = createTime
}

// GetAvatar 获取头像
func (u *User) GetAvatar() string {
	return u.Avatar
}

// SetAvatar 设置头像
func (u *User) SetAvatar(avatar string) {
	u.Avatar = avatar
}

// GetCode 获取验证码
func (u *User) GetCode() string {
	return u.Code
}

// SetCode 设置验证码
func (u *User) SetCode(code string) {
	u.Code = code
}

// GetState 获取状态
func (u *User) GetState() int {
	return u.State
}

// SetState 设置状态
func (u *User) SetState(state int) {
	u.State = state
}

// GetRID 获取角色ID
func (u *User) GetRID() int {
	return u.RID
}

// SetRID 设置角色ID
func (u *User) SetRID(rid int) {
	u.RID = rid
}
