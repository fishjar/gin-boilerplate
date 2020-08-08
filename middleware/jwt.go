/*
	中间件
*/

package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fishjar/gin-boilerplate/service"

	"github.com/gin-gonic/gin"
)

// JWTAuth 验证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取token
		// authorization := c.Request.Header.Get("Authorization")
		authorization := c.GetHeader("Authorization")
		accessToken := strings.Replace(authorization, "Bearer ", "", 1)

		// token 为空
		if len(accessToken) == 0 {
			// 验证失败
			service.HTTPAbortError(c, "没有权限：token不能为空", http.StatusUnauthorized, nil)
			return
		}

		// 解析token
		claims, err := service.ParseToken(accessToken)
		if claims == nil || err != nil {
			// 验证失败
			service.HTTPAbortError(c, "没有权限：JWT验证失败", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
			return
		}

		// // 从数据库验证token有效性
		// aid := claims.Subject
		// uid := claims.UserID
		// iss := claims.IssuedAt
		// auth, err := service.GetAuthWithUser(aid)
		// if err != nil { // 帐号不存在
		// 	service.HTTPAbortError(c, "没有权限：帐号不存在", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
		// 	return
		// }
		// if err := auth.CheckEnabled(); err != nil { // 禁用或过期
		// 	service.HTTPAbortError(c, "没有权限：禁用或过期", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
		// 	return
		// }
		// if auth.User.ID.String() != uid { // 用户数据有误
		// 	service.HTTPAbortError(c, "没有权限：用户数据有误", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
		// 	return
		// }
		// roles, err := auth.User.GetRoles() // 获取用户角色列表
		// if err != nil {
		// 	service.HTTPAbortError(c, "没有权限：帐号角色信息有误", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
		// 	return
		// }
		// roleNames := service.RolesToNames(roles)

		// 从redis认证token有效性
		user, err := service.GetUserFromRedis(claims.UserID)
		if err != nil {
			service.HTTPAbortError(c, "没有权限：获取用户缓存信息失败", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
			return
		} else if user.AuthID != claims.Subject {
			service.HTTPAbortError(c, "没有权限：帐号ID不一致", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
			return
		} else if user.IssuedAt != claims.IssuedAt {
			service.HTTPAbortError(c, "没有权限：签发时间不一致", http.StatusUnauthorized, fmt.Errorf("token:%s", accessToken))
			return
		}

		// 验证成功
		// 当前用户信息挂载到内存
		c.Set("UserInfo", user)

		// TODO: 返回一个新token给客户端(每次自动刷新token)

		c.Next()

	}
}
