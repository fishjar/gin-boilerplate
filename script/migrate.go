package script

import (
	"fmt"

	"github.com/fishjar/gin-boilerplate/db"
	"github.com/fishjar/gin-boilerplate/model"
)

// Migrate 同步数据库
func Migrate() {
	// 判断是否存在数据
	if db.DB.HasTable(&model.User{}) {
		fmt.Println("--------数据表已存在----------")
		return
	}
	db.DB.AutoMigrate(&model.Auth{}, &model.Group{}, &model.Menu{})
	db.DB.AutoMigrate(&model.Role{}, &model.User{}, &model.UserGroup{})
}
