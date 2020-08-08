package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/fishjar/gin-boilerplate/config"
	"github.com/fishjar/gin-boilerplate/db"
	"github.com/fishjar/gin-boilerplate/logger"
	"github.com/fishjar/gin-boilerplate/model"
)

// SetUserToRedis 保存用户登录信息到redis
func SetUserToRedis(user *model.UserCurrent) error {
	// 保存登录信息到redis
	userKey := "token:user:" + user.UserID
	err := db.Redis.HSet(userKey, map[string]interface{}{
		"aid":   user.AuthID,
		"iss":   strconv.FormatInt(user.IssuedAt, 10),
		"roles": strings.Join(user.Roles, ","),
	}).Err()
	if err != nil {
		return err
	}

	// 设置redis过期时间
	// TODO: 还应判断帐号的过期时间是否小于 config.Config.JWTExpiresIn
	db.Redis.Expire(userKey, config.Config.JWTExpiresIn)

	return nil
}

// ClearUserFromRedis 清除登录信息
func ClearUserFromRedis(uid string) error {
	userKey := "token:user:" + uid
	if err := db.Redis.Del(userKey).Err(); err != nil {
		go logger.LogReq.Warn(err)
		return err
	}
	return nil
}

// GetUserFromRedis 从缓存中获取用户信息
func GetUserFromRedis(uid string) (model.UserCurrent, error) {
	var user model.UserCurrent
	userKey := "token:user:" + uid
	userMap, err := db.Redis.HGetAll(userKey).Result()
	if err != nil {
		return user, err
	}

	iss, ok := userMap["iss"]
	if !ok {
		return user, errors.New("缓存信息不存在：iss")
	}
	aid, ok := userMap["aid"]
	if !ok {
		return user, errors.New("缓存信息不存在：aid")
	}
	roles, ok := userMap["roles"]
	if !ok {
		return user, errors.New("缓存信息不存在：roles")
	}
	issuedAt, _ := strconv.ParseInt(iss, 10, 64)
	user = model.UserCurrent{
		IssuedAt: issuedAt,
		UserID:   uid,
		AuthID:   aid,
		Roles:    strings.Split(roles, ","),
	}
	return user, nil
}
