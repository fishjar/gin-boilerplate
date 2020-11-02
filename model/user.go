package model

import (
	"time"

	"github.com/fishjar/gin-boilerplate/db"
	"github.com/fishjar/gin-boilerplate/utils"
)

// User 用户模型
type User struct {
	Base
	Name         *string      `json:"name" gorm:"column:name;type:VARCHAR(32)" binding:"min=3,max=20"`                                              // 姓名
	Nickname     string       `json:"nickname" gorm:"column:nickname;not null"`                                                                     // 昵称
	Gender       *int         `json:"gender" gorm:"column:gender;type:TINYINT;default:0" binding:"omitempty,eq=0|eq=1|eq=2"`                        // 性别
	Avatar       *string      `json:"avatar" gorm:"column:avatar" binding:"omitempty"`                                                              // 昵称
	Mobile       *string      `json:"mobile" gorm:"column:mobile;type:VARCHAR(16)" binding:"omitempty"`                                             // 手机
	Email        *string      `json:"email" gorm:"column:email;unique" binding:"omitempty,email"`                                                   // 邮箱
	Homepage     *string      `json:"homepage" gorm:"column:homepage" binding:"omitempty,url"`                                                      // 个人主页
	Birthday     *time.Time   `json:"birthday" gorm:"column:birthday;type:DATE" binding:"omitempty"`                                                // 生日
	Height       *float32     `json:"height" gorm:"column:height;type:FLOAT" binding:"omitempty,min=0.01,max=300"`                                  // 身高(cm)
	BloodType    *string      `json:"bloodType" gorm:"column:blood_type;type:VARCHAR(8)" binding:"omitempty,eq=A|eq=B|eq=AB|eq=O|eq=NULL"`          // 血型(ABO)
	Notice       *string      `json:"notice" gorm:"column:notice;type:TEXT" binding:"omitempty"`                                                    // 备注
	Intro        *string      `json:"intro" gorm:"column:intro;type:TEXT" binding:"omitempty"`                                                      // 简介
	Address      *string      `json:"address" gorm:"column:address;type:JSON" binding:"omitempty"`                                                  // 地址
	Lives        *string      `json:"lives" gorm:"column:lives;type:JSON" binding:"omitempty"`                                                      // 生活轨迹
	Tags         *string      `json:"tags" gorm:"column:tags;type:JSON" binding:"omitempty"`                                                        // 标签
	LuckyNumbers *string      `json:"luckyNumbers" gorm:"column:lucky_numbers;type:JSON" binding:"omitempty"`                                       // 幸运数字
	Score        *int         `json:"score" gorm:"column:score;default:0" binding:"omitempty"`                                                      // 积分
	UserNo       int          `json:"userNo" gorm:"column:user_no;AUTO_INCREMENT"`                                                                  // 编号
	Auths        []*Auth      `json:"auths" gorm:"foreignkey:UserID"`                                                                               // 帐号
	Groups       []*UserGroup `json:"groups" gorm:"foreignkey:UserID"`                                                                              // 用户组
	Roles        []*Role      `json:"roles" gorm:"many2many:userrole;"`                                                                             // 角色
	Friends      []*User      `json:"friends" gorm:"many2many:userfriend;association_jointable_foreignkey:user_id;jointable_foreignkey:friend_id;"` // 友
	// Groups       []*Group   `json:"groups" gorm:"many2many:usergroup;"`                                                                           // 组
}

// // UserUpdate 更新用户信息
// type UserUpdate struct {
// 	*User
// 	Name *string `json:"name" form:"-"`
// }

// // UserPublic 公开用户信息
// type UserPublic struct {
// 	*User
// 	// Nickname *string `json:"nickname" binding:"-"`
// 	// Gender   *int    `json:"gender,omitempty"`
// 	// Nickname *string `json:"nickname,omitempty"`
// 	Name string `json:"name" binding:"-"`
// }

// UserJWT JWT用户数据
type UserJWT struct {
	AuthID string `json:"aid" binding:"required"`
	UserID string `json:"uid" binding:"required"`
}

// UserCurrent JWT用户数据
type UserCurrent struct {
	IssuedAt int64    `json:"iss" binding:"required"`
	AuthID   string   `json:"aid" binding:"required"`
	UserID   string   `json:"uid" binding:"required"`
	Roles    []string `json:"roles" binding:"required"`
}

// UserRes 返回单个
type UserRes struct {
	HTTPSuccess
	Data User `json:"data" binding:"required"`
}

type UserPaginData struct {
	Pagin PaginRes `json:"pagin" binding:"required"`
	Rows  []User   `json:"rows" binding:"required"`
}

type UserListRes2 struct {
	HTTPSuccess
	Data UserPaginData `json:"data" binding:"required"`
}

// UserListRes 返回列表
type UserListRes struct {
	HTTPSuccess
	Pagin PaginRes `json:"pagin" binding:"required"`
	Data  []User   `json:"data" binding:"required"`
}

// GetRoles 获取用户角色列表
func (user User) GetRoles() ([]Role, error) {
	var roles []Role
	if err := db.DB.Model(&user).Preload("Menus").Related(&roles, "Roles").Error; err != nil {
		return roles, err
	}
	return roles, nil
}

// GetGroups 获取用户组列表
func (user User) GetGroups() ([]UserGroup, error) {
	var userGroups []UserGroup
	if err := db.DB.Model(&user).Preload("Group").Related(&userGroups, "UserGroups").Error; err != nil {
		return userGroups, err
	}
	return userGroups, nil
}

// GetMenus 获取用户菜单列表
func (user User) GetMenus() ([]Menu, error) {
	var menus []Menu
	var tmpMenus []utils.IFUniqueItem

	roles, err := user.GetRoles()
	if err != nil {
		return menus, err
	}

	for _, role := range roles {
		for _, menu := range role.Menus {
			tmpMenus = append(tmpMenus, *menu)
		}
	}
	tmpMenus = utils.RemoveDuplicateElemt(tmpMenus) // 去重
	for _, v := range tmpMenus {
		menus = append(menus, v.(Menu))
	}

	return menus, nil
}

// TableName 自定义用户表名
func (User) TableName() string {
	return "user"
}
