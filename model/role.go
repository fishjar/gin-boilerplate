package model

// Role 定义模型
type Role struct {
	Base
	Name  string  `json:"name" gorm:"column:name;type:VARCHAR(32);unique;not null" binding:"min=3,max=20"` // 角色名称
	Users []*User `json:"users" gorm:"many2many:userrole;"`                                                // 用户
	Menus []*Menu `json:"menus" gorm:"many2many:rolemenu;"`                                                // 菜单
}

// RoleRes 返回单个
type RoleRes struct {
	HTTPSuccess
	Data Role `json:"data" binding:"required"`
}

// RoleListRes 角色列表
type RoleListRes struct {
	HTTPSuccess
	Pagin PaginRes `json:"pagin" binding:"required"`
	Data  []Role   `json:"data" binding:"required"`
}

// TableName 自定义表名
func (Role) TableName() string {
	return "role"
}
