package model

import (
	uuid "github.com/satori/go.uuid"
)

// Group 定义模型
type Group struct {
	Base
	Name     string       `json:"name" gorm:"column:name;type:VARCHAR(32);not null" binding:"min=3,max=20"` // 组名称
	LeaderID uuid.UUID    `json:"leaderId" gorm:"column:leader_id;not null"`                                // 队长ID
	Leader   *User        `json:"leader" gorm:"foreignkey:LeaderID"`                                        // 队长
	Members  []*UserGroup `json:"members" gorm:"foreignkey:GroupID"`                                        // 成员
	// Users    []*User   `json:"users" gorm:"many2many:usergroup;"`                                        // 队员
}

// GroupRes 返回单个
type GroupRes struct {
	HTTPSuccess
	Data Group `json:"data" binding:"required"`
}

// GroupListRes 返回列表
type GroupListRes struct {
	HTTPSuccess
	Pagin PaginRes `json:"pagin" binding:"required"`
	Data  []Group  `json:"data" binding:"required"`
}

// TableName 自定义表名
func (Group) TableName() string {
	return "group"
}
