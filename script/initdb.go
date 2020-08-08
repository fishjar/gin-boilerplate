package script

import (
	"fmt"
	"time"

	"github.com/fishjar/gin-boilerplate/db"
	"github.com/fishjar/gin-boilerplate/model"
	"github.com/fishjar/gin-boilerplate/utils"
)

// InitDB 数据库初始化
func InitDB() {
	// 创建默认用户
	name := "gabe"
	gender := 1
	user := model.User{
		Nickname: name,
		Gender:   &gender,
	}
	if err := db.DB.Where(&user).First(&user).Error; err == nil {
		fmt.Println("-------数据已存在-----")
		return
	}

	// 创建默认用户
	db.DB.Create(&user)
	fmt.Println("默认用户已创建")
	fmt.Println("user：", user)

	// 创建认证帐号
	authType := "account"
	authCode := utils.MD5Pwd("gabe", "123456")
	now := time.Now()
	auth := model.Auth{
		UserID:     user.ID,
		AuthType:   authType,
		AuthName:   name,
		AuthCode:   &authCode,
		VerifyTime: &now,
	}
	db.DB.Create(&auth)

	// 创建测试用户
	jack := model.User{Nickname: "jack"}
	rose := model.User{Nickname: "rose"}
	db.DB.Create(&jack)
	db.DB.Create(&rose)

	// 创建角色
	adminRole := model.Role{Name: "admin"}
	userRole := model.Role{Name: "user"}
	guestRole := model.Role{Name: "guest"}
	db.DB.Create(&adminRole)
	db.DB.Create(&userRole)
	db.DB.Create(&guestRole)

	// 创建组
	titanicGroup := model.Group{
		Name:     "titanic",
		LeaderID: jack.ID,
	}
	rayjarGroup := model.Group{
		Name:     "rayjar",
		LeaderID: user.ID,
	}
	db.DB.Create(&titanicGroup)
	db.DB.Create(&rayjarGroup)

	// 创建菜单
	icon := "smile"
	sort := 0
	welcomeMenu := model.Menu{
		Name: "welcome",
		Path: "/welcome",
		Icon: &icon,
		Sort: &sort,
	}
	db.DB.Create(&welcomeMenu)

	icon = "dashboard"
	sort = 1
	dashboardMenu := model.Menu{
		Name: "dashboard",
		Path: "/dashboard",
		Icon: &icon,
		Sort: &sort,
	}
	db.DB.Create(&dashboardMenu)

	sort = 0
	usersMenu := model.Menu{
		ParentID: dashboardMenu.ID,
		Name:     "welcome",
		Path:     "/dashboard/users",
		Sort:     &sort,
	}
	db.DB.Create(&usersMenu)

	sort = 1
	authsMenu := model.Menu{
		ParentID: dashboardMenu.ID,
		Name:     "auths",
		Path:     "/dashboard/auths",
		Sort:     &sort,
	}
	db.DB.Create(&authsMenu)

	// 关联角色菜单
	db.DB.Model(&adminRole).Association("Menus").Append([]model.Menu{welcomeMenu, dashboardMenu, usersMenu, authsMenu})
	db.DB.Model(&userRole).Association("Menus").Append([]model.Menu{welcomeMenu, dashboardMenu, usersMenu})

	// 关联用户角色
	db.DB.Model(&user).Association("Roles").Append([]model.Role{adminRole, userRole, guestRole})
	db.DB.Model(&jack).Association("Roles").Append([]model.Role{userRole, guestRole})
	db.DB.Model(&rose).Association("Roles").Append([]model.Role{guestRole})

	// 关联用户团队
	// db.DB.Model(&user).Association("Groups").Append([]model.Group{titanicGroup})
	// db.DB.Model(&jack).Association("Groups").Append([]model.Group{titanicGroup, rayjarGroup})
	// db.DB.Model(&rose).Association("Groups").Append([]model.Group{titanicGroup, rayjarGroup})
	level1 := 1
	db.DB.Create(&model.UserGroup{
		User:  &user,
		Group: &rayjarGroup,
		Level: &level1,
	})
	db.DB.Create(&model.UserGroup{
		User:  &user,
		Group: &titanicGroup,
		Level: &level1,
	})
	db.DB.Create(&model.UserGroup{
		User:  &jack,
		Group: &rayjarGroup,
		Level: &level1,
	})
	db.DB.Create(&model.UserGroup{
		User:  &jack,
		Group: &titanicGroup,
		Level: &level1,
	})
	db.DB.Create(&model.UserGroup{
		User:  &rose,
		Group: &titanicGroup,
		Level: &level1,
	})

	// 关联用户友谊
	db.DB.Model(&user).Association("Friends").Append(jack)
	db.DB.Model(&user).Association("Friends").Append(rose)
	db.DB.Model(&jack).Association("Friends").Append(rose)
}
