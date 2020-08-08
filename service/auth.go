package service

import (
	"time"

	"github.com/fishjar/gin-boilerplate/db"
	"github.com/fishjar/gin-boilerplate/model"
	"github.com/fishjar/gin-boilerplate/utils"
)

// GetAuth 获取指定ID认证帐号
func GetAuth(id string) (model.Auth, error) {
	var auth model.Auth
	if err := db.DB.First(&auth, "id = ?", id).Error; err != nil {
		return auth, err
	}
	return auth, nil
}

// GetAuthWithRoles 获取指定ID认证帐号
func GetAuthWithRoles(id string) (model.Auth, error) {
	var auth model.Auth
	if err := db.DB.Preload("User").Preload("User.Roles").First(&auth, "id = ?", id).Error; err != nil {
		return auth, err
	}
	return auth, nil
}

// GetAuthWithUser 获取指定ID认证帐号及用户资料
func GetAuthWithUser(id string) (model.Auth, error) {
	var auth model.Auth
	if err := db.DB.Preload("User").First(&auth, "id = ?", id).Error; err != nil {
		return auth, err
	}
	return auth, nil
}

// // AuthAndUserCheck 从数据库检查authID和userID有效性
// func AuthAndUserCheck(authID string, userID string) bool {
// 	if auth, err := GetAuth(authID); err != nil {
// 		return false // Auth不存在
// 	} else if err := auth.CheckEnabled(); err != nil {
// 		return false // 禁用或过期
// 	}
// 	if _, err := GetUser(userID); err != nil {
// 		return false // User不存在
// 	}
// 	return true
// }

// CreateAuthAccount 创建帐号
func CreateAuthAccount(data *model.AuthAccountCreateReq) error {
	// 开始事务
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 回滚
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 创建用户
	user := model.User{
		Nickname: data.Username,
	}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 创建帐号
	password := utils.MD5Pwd(data.Username, data.Password)
	now := time.Now()
	auth := model.Auth{
		User:       &user,
		AuthType:   "account",
		AuthName:   data.Username,
		AuthCode:   &password,
		VerifyTime: &now,
	}
	if err := tx.Create(&auth).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
