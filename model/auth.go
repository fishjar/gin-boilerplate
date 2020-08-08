package model

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Auth 定义模型
type Auth struct {
	Base
	UserID     uuid.UUID  `json:"userId" gorm:"column:user_id;not null"`                                  // 用户ID
	User       *User      `json:"user" gorm:"foreignkey:UserID"`                                          // 用户
	AuthType   string     `json:"authType" gorm:"column:auth_type;type:VARCHAR(16);not null"`             // 鉴权类型
	AuthName   string     `json:"authName" gorm:"column:auth_name;type:VARCHAR(128);not null"`            // 鉴权名称
	AuthCode   *string    `json:"authCode" gorm:"column:auth_code" binding:"omitempty"`                   // 鉴权识别码
	VerifyTime *time.Time `json:"verifyTime" gorm:"column:verify_time;type:DATETIME" binding:"omitempty"` // 认证时间
	ExpireTime *time.Time `json:"expireTime" gorm:"column:expire_time;type:DATETIME" binding:"omitempty"` // 过期时间
	IsEnabled  bool       `json:"isEnabled" gorm:"column:is_enabled;default:true" binding:"omitempty"`    // 是否启用
}

// AuthRes 返回单个
type AuthRes struct {
	HTTPSuccess
	Data Auth `json:"data" binding:"required"`
}

// AuthListRes 返回列表
type AuthListRes struct {
	HTTPSuccess
	Pagin PaginRes `json:"pagin" binding:"required"`
	Data  []Auth   `json:"data" binding:"required"`
}

// // AuthPublic 公开模型
// type AuthPublic struct {
// 	*Auth
// 	AuthName string  `json:"-" binding:"-"`
// 	AuthCode *string `json:"-" binding:"-"`
// }

// AuthAccountLoginReq 帐号登录表单
type AuthAccountLoginReq struct {
	Username string `json:"username" form:"username" binding:"required" example:"gabe"`   // 用户名
	Password string `json:"password" form:"password" binding:"required" example:"123456"` // 密码
}

// AuthAccountLoginRes 登录成功返回数据
type AuthAccountLoginRes struct {
	TokenType   string `json:"tokenType" binding:"required"`   // token类型
	AccessToken string `json:"accessToken" binding:"required"` // token
	ExpiresIn   int64  `json:"expiresIn" binding:"required"`   // 过期时间（毫秒）
}

// AuthAccountLoginSuccess 登录成功
type AuthAccountLoginSuccess struct {
	HTTPSuccess
	Data AuthAccountLoginRes `json:"data" binding:"required"`
}

// AuthAccountCreateReq 创建帐号
type AuthAccountCreateReq struct {
	Username string `form:"username" binding:"required"` // 帐号
	Password string `form:"password" binding:"required"` // 密码
	// Nickname string `form:"nickname" binding:"required"` // 昵称
	// Mobile   string `form:"mobile" binding:"required"`   // 手机
}

// CheckEnabled 检查帐号有效性
func (auth Auth) CheckEnabled() error {
	if !auth.IsEnabled {
		return errors.New("帐号已禁用") // 禁用
	}
	if auth.VerifyTime == nil {
		return errors.New("帐号未认证") // 未认证
	}
	expireTime := auth.ExpireTime
	if expireTime != nil && expireTime.Before(time.Now()) {
		return errors.New("帐号已过期") // 过期
	}
	return nil
}

// TableName 自定义表名
func (Auth) TableName() string {
	return "auth"
}

// 自定义验证器
// var bookableDate validator.Func = func(fl validator.FieldLevel) bool
// v.RegisterValidation("bookabledate", bookableDate)
// v.RegisterStructValidation(UserStructLevelValidation, User{})
