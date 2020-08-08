/*
	路由配置
*/

package router

import (
	"github.com/fishjar/gin-boilerplate/config"
	"github.com/fishjar/gin-boilerplate/handler"
	"github.com/fishjar/gin-boilerplate/middleware"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter 注入路由，及返回一个gin对象
func InitRouter() *gin.Engine {

	// r := gin.New()
	r := gin.Default()               // Default 使用 Logger 和 Recovery 中间件
	r.Use(middleware.LoggerToFile()) // 日志中间件
	r.Use(cors.Default())            // 跨域中间件，Default() allows all origins
	r.GET("/ping", handler.Pong)

	r.GET("/swagger/*any", gin.BasicAuth(gin.Accounts{
		config.Config.SwaggerName: config.Config.SwaggerPassword,
	}), ginSwagger.WrapHandler(swaggerFiles.Handler)) // swagger
	r.POST("/admin/account/login", handler.LoginAccount) //登录

	admin := r.Group("/admin")      // JWT验证路由组
	admin.Use(middleware.JWTAuth()) // JWT验证中间件
	rc := middleware.RoleCheck      // 角色检查中间件，TODO:角色权限可以做到数据库里面管理
	{
		admin.POST("/account/create", handler.AuthAccountCreate)   // 创建帐号
		admin.POST("/token/refresh", handler.TokenRefresh)         // 刷新token
		admin.POST("/account/logout", handler.LogoutAccount)       // 退出登录
		admin.POST("/files/upload", handler.UploudFile)            // 上传单文件
		admin.POST("/files/upload/multi", handler.UploudMultiFile) // 上传多文件
	}
	{

		admin.GET("/auths", handler.AuthFindAndCountAll)    // 获取多条
		admin.GET("/auths/:id", handler.AuthFindByPk)       // 按ID查找
		admin.POST("/auths", handler.AuthSingleCreate)      // 创建单条
		admin.PATCH("/auths/:id", handler.AuthUpdateByPk)   // 按ID更新
		admin.DELETE("/auths/:id", handler.AuthDestroyByPk) // 按ID删除
		admin.POST("/auth", handler.AuthFindOrCreate)       // 查询或创建
		admin.PATCH("/auths", handler.AuthUpdateBulk)       // 批量更新
		admin.DELETE("/auths", handler.AuthDestroyBulk)     // 批量删除
	}
	{
		admin.GET("/groups", handler.GroupFindAndCountAll)    // 获取多条
		admin.GET("/groups/:id", handler.GroupFindByPk)       // 按ID查找
		admin.POST("/groups", handler.GroupSingleCreate)      // 创建单条
		admin.PATCH("/groups/:id", handler.GroupUpdateByPk)   // 按ID更新
		admin.DELETE("/groups/:id", handler.GroupDestroyByPk) // 按ID删除
		admin.POST("/group", handler.GroupFindOrCreate)       // 查询或创建
		admin.PATCH("/groups", handler.GroupUpdateBulk)       // 批量更新
		admin.DELETE("/groups", handler.GroupDestroyBulk)     // 批量删除
	}
	{
		admin.GET("/menus", handler.MenuFindAndCountAll)    // 获取多条
		admin.GET("/menus/:id", handler.MenuFindByPk)       // 按ID查找
		admin.POST("/menus", handler.MenuSingleCreate)      // 创建单条
		admin.PATCH("/menus/:id", handler.MenuUpdateByPk)   // 按ID更新
		admin.DELETE("/menus/:id", handler.MenuDestroyByPk) // 按ID删除
		admin.POST("/menu", handler.MenuFindOrCreate)       // 查询或创建
		admin.PATCH("/menus", handler.MenuUpdateBulk)       // 批量更新
		admin.DELETE("/menus", handler.MenuDestroyBulk)     // 批量删除
	}
	{
		admin.GET("/roles", handler.RoleFindAndCountAll)    // 获取多条
		admin.GET("/roles/:id", handler.RoleFindByPk)       // 按ID查找
		admin.POST("/roles", handler.RoleSingleCreate)      // 创建单条
		admin.PATCH("/roles/:id", handler.RoleUpdateByPk)   // 按ID更新
		admin.DELETE("/roles/:id", handler.RoleDestroyByPk) // 按ID删除
		admin.POST("/role", handler.RoleFindOrCreate)       // 查询或创建
		admin.PATCH("/roles", handler.RoleUpdateBulk)       // 批量更新
		admin.DELETE("/roles", handler.RoleDestroyBulk)     // 批量删除
	}
	{
		admin.GET("/users", rc([]string{"admin"}), handler.UserFindAndCountAll) // 获取多条
		admin.GET("/users/:id", handler.UserFindByPk)                           // 按ID查找
		admin.POST("/users", handler.UserSingleCreate)                          // 创建单条
		admin.PATCH("/users/:id", handler.UserUpdateByPk)                       // 按ID更新
		admin.DELETE("/users/:id", handler.UserDestroyByPk)                     // 按ID删除
		admin.POST("/user", handler.UserFindOrCreate)                           // 查询或创建
		admin.PATCH("/users", handler.UserUpdateBulk)                           // 批量更新
		admin.DELETE("/users", handler.UserDestroyBulk)                         // 批量删除
		admin.GET("/user/roles", handler.UserFindMyRoles)                       // 获取角色列表
		admin.GET("/user/menus", handler.UserFindMyMenus)                       // 获取菜单列表
		admin.GET("/user/groups", handler.UserFindMyGroups)                     // 获取组列表
	}
	{
		admin.GET("/usergroups", handler.UserGroupFindAndCountAll)    // 获取多条
		admin.GET("/usergroups/:id", handler.UserGroupFindByPk)       // 按ID查找
		admin.POST("/usergroups", handler.UserGroupSingleCreate)      // 创建单条
		admin.PATCH("/usergroups/:id", handler.UserGroupUpdateByPk)   // 按ID更新
		admin.DELETE("/usergroups/:id", handler.UserGroupDestroyByPk) // 按ID删除
		admin.POST("/usergroup", handler.UserGroupFindOrCreate)       // 查询或创建
		admin.PATCH("/usergroups", handler.UserGroupUpdateBulk)       // 批量更新
		admin.DELETE("/usergroups", handler.UserGroupDestroyBulk)     // 批量删除
	}

	public := r.Group("/public") // 大众用户路由组
	{
		public.GET("/ping", func(c *gin.Context) { // pingpong
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	return r
}
