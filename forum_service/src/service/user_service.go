package service

import (
	"errors"
	"forum_service/src/db"
	"forum_service/src/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService 用户服务
type UserService struct {
	db *gorm.DB
}

// JWT密钥
var jwtSecret = []byte("forum_secret_key")

// Claims 自定义声明结构体
type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// NewUserService 创建用户服务实例
func NewUserService(database *gorm.DB) *UserService {
	return &UserService{db: database}
}

// Register 用户注册
func (s *UserService) Register(username, password, email string, gender int) (*model.User, error) {
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建用户对象
	user := &model.User{
		Username:   username,
		Password:   string(hashedPassword),
		Email:      email,
		Gender:     gender,
		CreateTime: time.Now(),
		State:      1, // 默认状态为正常
	}

	// 保存用户到数据库
	if err := db.AddUser(s.db, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login 用户登录
func (s *UserService) Login(username, password string) (*model.User, error) {
	// 根据用户名查询用户
	user, err := db.GetUserByUsername(s.db, username)
	if err != nil {
		return nil, err
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	// 返回用户信息（不包含密码）
	return user, nil
}

// GenerateToken 生成JWT token
func (s *UserService) GenerateToken(user *model.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "forum",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 解析JWT token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// GetInfo 根据token获取用户信息
func (s *UserService) GetInfo(token string) (*model.User, error) {
	// 解析token
	claims, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	// 根据用户ID查询用户信息
	var user model.User
	err = s.db.Where("id = ?", claims.UserID).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
