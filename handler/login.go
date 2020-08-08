/*
	对应路由的处理函数
*/

package handler

import (
	"fmt"
	"net/http"

	"github.com/fishjar/gin-boilerplate/config"
	"github.com/fishjar/gin-boilerplate/db"
	"github.com/fishjar/gin-boilerplate/model"
	"github.com/fishjar/gin-boilerplate/service"
	"github.com/fishjar/gin-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

// LoginAccount 登录处理函数
// TODO：生产环境，错误信息不需要详细情况
// @Summary				帐号登录
// @Description			帐号登录...
// @Tags				admin
// @Accept				json
// @Produce				json
// @Param				q body model.AuthAccountLoginReq true "登录"
// @Success				200 {object} model.AuthAccountLoginSuccess
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/account/login [post]
func LoginAccount(c *gin.Context) {

	var loginForm model.AuthAccountLoginReq

	// 绑定数据
	if err := c.ShouldBind(&loginForm); err != nil {
		service.HTTPError(c, "登录失败，参数有误", http.StatusUnauthorized, err)
		return
	}

	// 查询帐号名是否存在
	authType := "account"
	password := utils.MD5Pwd(loginForm.Username, loginForm.Password)
	var auth model.Auth
	if err := db.DB.Where(&model.Auth{
		AuthName: loginForm.Username,
		AuthType: authType,
	}).First(&auth).Error; err != nil {
		service.HTTPError(c, "登录失败，用户名不存在", http.StatusUnauthorized, err)
		return
	}

	// 检查禁用或过期
	if err := auth.CheckEnabled(); err != nil {
		service.HTTPError(c, "登录失败，帐号禁用或过期", http.StatusUnauthorized, err)
		return
	}

	// 验证密码
	if password != *auth.AuthCode {
		service.HTTPError(c, "登录失败，密码错误", http.StatusUnauthorized, fmt.Errorf("username:%s, password:%s", loginForm.Username, loginForm.Password))
		return
	}

	// 查询用户是否存在
	user, err := service.GetUser(auth.UserID.String())
	if err != nil {
		service.HTTPError(c, "登录失败，用户数据有误", http.StatusUnauthorized, err)
		return
	}

	// 查询角色列表
	roles, err := user.GetRoles()
	if err != nil {
		service.HTTPError(c, "登录失败，获取角色列表失败", http.StatusUnauthorized, err)
		return
	}
	roleNames := service.RolesToNames(roles)
	aid := auth.ID.String()
	uid := auth.UserID.String()

	// 生成token
	issuedAt, accessToken, err := service.MakeToken(&model.UserJWT{
		AuthID: aid,
		UserID: uid,
	})
	if err != nil {
		service.HTTPError(c, "登录失败，获取token失败", http.StatusUnauthorized, err)
		return
	}

	// 保存登录信息到redis
	err = service.SetUserToRedis(&model.UserCurrent{
		IssuedAt: issuedAt,
		UserID:   uid,
		AuthID:   aid,
		Roles:    roleNames,
	})
	if err != nil {
		service.HTTPError(c, "保存登录信息到redis失败", http.StatusInternalServerError, err)
		return
	}

	// 登录成功
	c.JSON(http.StatusOK, model.AuthAccountLoginSuccess{
		HTTPSuccess: model.HTTPSuccess{
			Message: "登录成功",
		},
		Data: model.AuthAccountLoginRes{
			TokenType:   "Bearer",
			AccessToken: accessToken,
			ExpiresIn:   config.Config.JWTExpiresIn.Milliseconds(),
		},
	})

}

// TokenRefresh 刷新token
// @Summary				刷新token
// @Description			刷新token...
// @Tags				admin
// @Accept				json
// @Produce				json
// @Success				200 {object} model.AuthAccountLoginSuccess
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/token/refresh [post]
// @Security			ApiKeyAuth
func TokenRefresh(c *gin.Context) {
	// 获取当前用户信息
	user := c.MustGet("UserInfo").(model.UserCurrent)
	aid := user.AuthID
	uid := user.UserID
	roles := user.Roles

	// 生成token
	// TODO: 返回JWT过期时间，并与redis 过期时间保持一致
	issuedAt, newToken, err := service.MakeToken(&model.UserJWT{
		AuthID: aid,
		UserID: uid,
	})
	if err != nil {
		service.HTTPError(c, "刷新token失败", http.StatusInternalServerError, err)
		return
	}

	// 保存登录信息到redis
	if err := service.SetUserToRedis(&model.UserCurrent{
		IssuedAt: issuedAt,
		UserID:   uid,
		AuthID:   aid,
		Roles:    roles,
	}); err != nil {
		service.HTTPError(c, "保存登录信息到redis失败", http.StatusInternalServerError, err)
		return
	}

	// 更新成功
	c.JSON(http.StatusOK, model.AuthAccountLoginSuccess{
		HTTPSuccess: model.HTTPSuccess{
			Message: "刷新成功",
		},
		Data: model.AuthAccountLoginRes{
			TokenType:   "Bearer",
			AccessToken: newToken,
			ExpiresIn:   config.Config.JWTExpiresIn.Milliseconds(),
		},
	})
}

// LogoutAccount 登出
// @Summary				退出登录
// @Description			退出登录...
// @Tags				admin
// @Accept				json
// @Produce				json
// @Success				200 {object} model.HTTPSuccess
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/account/logout [post]
// @Security			ApiKeyAuth
func LogoutAccount(c *gin.Context) {
	user := c.MustGet("UserInfo").(model.UserCurrent)
	if err := service.ClearUserFromRedis(user.UserID); err != nil {
		service.HTTPError(c, "退出登录失败", http.StatusInternalServerError, err)
	}
	service.HTTPSuccess(c, "退出成功")
}
