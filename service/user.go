package service

import (
	"github.com/fishjar/gin-boilerplate/db"
	"github.com/fishjar/gin-boilerplate/model"
	"github.com/gin-gonic/gin"
)

// GetUser 获取指定ID用户
func GetUser(id string) (model.User, error) {
	var user model.User
	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetCurrentUser 获取当前用户
func GetCurrentUser(c *gin.Context) (model.User, error) {
	userInfo := c.MustGet("UserInfo").(model.UserCurrent)
	uid := userInfo.UserID

	user, err := GetUser(uid)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetUserMenus 获取用户菜单（作废）
func GetUserMenus(user model.User) ([]model.Menu, error) {
	var menus []model.Menu

	roles, err := user.GetRoles()
	if err != nil {
		return menus, err
	}

	for _, role := range roles {
		for _, menu := range role.Menus {
			menus = append(menus, *menu)
		}
	}
	menus = RemoveDuplicateMenu(menus) // 去重

	return menus, nil
}

// GetCurrentUserRoles 获取当前用户角色列表
func GetCurrentUserRoles(c *gin.Context) ([]model.Role, error) {
	var roles []model.Role

	user, err := GetCurrentUser(c)
	if err != nil {
		return roles, err
	}

	roles, err = user.GetRoles()
	if err != nil {
		return roles, err
	}

	return roles, nil
}

// GetCurrentUserGroups 获取当前用户组列表
func GetCurrentUserGroups(c *gin.Context) ([]model.UserGroup, error) {
	var userGroup []model.UserGroup

	user, err := GetCurrentUser(c)
	if err != nil {
		return userGroup, err
	}

	userGroup, err = user.GetGroups()
	if err != nil {
		return userGroup, err
	}

	return userGroup, nil
}

// GetCurrentUserMenus 获取当前用户菜单列表
func GetCurrentUserMenus(c *gin.Context) ([]model.Menu, error) {
	var menus []model.Menu

	user, err := GetCurrentUser(c)
	if err != nil {
		return menus, err
	}

	menus, err = user.GetMenus()
	if err != nil {
		return menus, err
	}

	return menus, nil
}

// RemoveDuplicateMenu Menu去重（作废，已用IFUniqueItem接口实现）
func RemoveDuplicateMenu(menus []model.Menu) []model.Menu {
	result := make([]model.Menu, 0, len(menus))
	temp := map[string]struct{}{}
	for _, item := range menus {
		id := item.ID.String()
		if _, ok := temp[id]; !ok {
			temp[id] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// RolesToNames 获取角色名列表
func RolesToNames(roles []model.Role) []string {
	roleNames := make([]string, len(roles))
	for i, role := range roles {
		roleNames[i] = role.Name
	}
	return roleNames
}
