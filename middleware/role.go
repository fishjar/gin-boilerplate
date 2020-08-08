package middleware

import (
	"fmt"
	"net/http"

	"github.com/fishjar/gin-boilerplate/model"
	"github.com/fishjar/gin-boilerplate/utils"

	"github.com/gin-gonic/gin"
)

// RoleCheck 角色检查中间件
// TODO:角色权限可以做到数据库里面管理
func RoleCheck(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo := c.MustGet("UserInfo").(model.UserCurrent)
		userRoles := userInfo.Roles
		fmt.Println("allowedRoles", allowedRoles)
		fmt.Println("userRoles", userRoles)

		if len(allowedRoles) == 0 { // 不允许任何角色
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "没有权限：不允许任何角色",
			})
			return
		}
		if len(userRoles) == 0 { // 用户未授予角色
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "没有权限：用户未授予角色",
			})
			return
		}
		if len(utils.Intersect(allowedRoles, userRoles)) == 0 { // 用户角色未授权
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "没有权限：用户角色未授权",
			})
			return
		}

		c.Next()
	}
}
