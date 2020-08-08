package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// UserGroup 定义模型
type UserGroup struct {
	Base
	UserID   uuid.UUID  `json:"userId" gorm:"column:user_id;not null"`                                // 用户ID
	User     *User      `json:"user" gorm:"foreignkey:UserID"`                                        // 用户
	GroupID  uuid.UUID  `json:"groupId" gorm:"column:group_id;not null"`                              // 组ID
	Group    *Group     `json:"group" gorm:"foreignkey:GroupID"`                                      // 组
	Level    *int       `json:"level" gorm:"column:level;type:TINYINT;default:0" binding:"omitempty"` // 级别
	JoinTime *time.Time `json:"joinTime" gorm:"column:join_time;type:DATETIME" binding:"omitempty"`   // 加入时间
}

// UserGroupRes 返回单个
type UserGroupRes struct {
	HTTPSuccess
	Data UserGroup `json:"data" binding:"required"`
}

// UserGroupListRes 返回列表
type UserGroupListRes struct {
	HTTPSuccess
	Pagin PaginRes    `json:"pagin" binding:"required"`
	Data  []UserGroup `json:"data" binding:"required"`
}

// TableName 自定义表名
func (UserGroup) TableName() string {
	return "usergroup"
}
