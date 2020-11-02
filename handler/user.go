package handler

import (
	"net/http"

	"github.com/fishjar/gin-boilerplate/db"
	"github.com/fishjar/gin-boilerplate/model"
	"github.com/fishjar/gin-boilerplate/service"

	"github.com/gin-gonic/gin"
)

// UserFindAndCountAll 查询多条信息
// @Summary				查询多条信息
// @Description			查询多条信息...
// @Tags				user
// @Accept				json
// @Produce				json
// @Param				q query model.PaginReq false "参数"
// @Success				200 {object} model.UserListRes
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/users [get]
// @Security			ApiKeyAuth
func UserFindAndCountAll(c *gin.Context) {

	// 参数绑定
	var q *model.PaginReq
	if err := c.ShouldBindQuery(&q); err != nil {
		// c.JSON(http.StatusBadRequest, model.HTTPError{
		// 	Code:    http.StatusBadRequest,
		// 	Message: "参数有误",
		// 	Errors:  []error{err},
		// })
		service.HTTPError(c, "参数有误", http.StatusBadRequest, err)
		return
	}

	// 条件参数
	params := c.QueryMap("params")
	// map 参数，map[string]T 必须转为 map[string]interface{}
	where := make(map[string]interface{}, len(params))
	for k, v := range params {
		// 有模型不存在的key时，后面的查询会报错，需要过滤掉
		if k == "name" || k == "gender" {
			where[k] = v
		}
	}
	// struct 参数，后面使用指针
	// var where model.User
	// if err := mapstructure.Decode(params, &where); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"err": err.Error(),
	// 		"message": "查询参数有误",
	// 	})
	// 	return
	// }

	// 分页参数
	offset := (q.Page - 1) * q.Size
	var total uint
	var rows []model.User

	// 查询数据
	if err := db.DB.Model(&rows).Where(where).Count(&total).Limit(q.Size).Offset(offset).Order(q.Sort).Preload("Auths").Preload("Roles").Preload("Groups").Preload("Friends").Find(&rows).Error; err != nil {
		service.HTTPError(c, "查询多条信息失败", http.StatusInternalServerError, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.UserListRes2{
		Data: model.UserPaginData{
			Pagin: q.Pagin(total),
			Rows:  rows,
		},
	})
}

// UserFindByPk 根据主键查询单条信息
// @Summary				查询单条信息
// @Description			查询单条信息...
// @Tags				user
// @Accept				json
// @Produce				json
// @Param				id path string true "ID"
// @Success				200 {object} model.UserRes
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/users/{id} [get]
// @Security			ApiKeyAuth
func UserFindByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.User
	if err := db.DB.Preload("Auths").Preload("Roles").Preload("Groups").Preload("Friends").First(&data, "id = ?", id).Error; err != nil {
		service.HTTPError(c, "查询失败", http.StatusNotFound, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.UserRes{
		Data: data,
	})
}

// UserSingleCreate 创建单条信息
// @Summary				创建单条信息
// @Description			创建单条信息...
// @Tags				user
// @Accept				json
// @Produce				json
// @Param				user body model.User true "参数"
// @Success				200 {object} model.UserRes
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/users [post]
// @Security			ApiKeyAuth
func UserSingleCreate(c *gin.Context) {

	// 绑定数据
	var data model.User
	if err := c.ShouldBind(&data); err != nil {
		service.HTTPError(c, "数据绑定失败", http.StatusBadRequest, err)
		return
	}

	// 插入数据
	if err := db.DB.Create(&data).Error; err != nil {
		service.HTTPError(c, "插入数据失败", http.StatusInternalServerError, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.UserRes{
		Data: data,
	})
}

// UserUpdateByPk 更新单条信息
// @Summary				更新单条信息
// @Description			更新单条信息...
// @Tags				user
// @Accept				json
// @Produce				json
// @Param				id path string true "ID"
// @Param				user body model.User true "更新单条信息"
// @Success				200 {object} model.UserRes
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/users/{id} [patch]
// @Security			ApiKeyAuth
func UserUpdateByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.User
	if err := db.DB.First(&data, "id = ?", id).Error; err != nil {
		service.HTTPError(c, "查询失败", http.StatusNotFound, err)
		return
	}

	// 绑定新数据
	var obj map[string]interface{}
	if err := c.ShouldBind(&obj); err != nil {
		service.HTTPError(c, "数据绑定失败", http.StatusBadRequest, err)
		return
	}

	// 更新数据
	if err := db.DB.Model(&data).Updates(obj).Error; err != nil {
		service.HTTPError(c, "更新失败", http.StatusInternalServerError, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.UserRes{
		Data: data,
	})
}

// UserDestroyByPk 删除单条信息
// @Summary				删除单条信息
// @Description			删除单条信息...
// @Tags				user
// @Accept				json
// @Produce				json
// @Param				id path string true "ID"
// @Success				200 {object} model.HTTPDeleteSuccess
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/users/{id} [delete]
// @Security			ApiKeyAuth
func UserDestroyByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.User
	if err := db.DB.Where("id = ?", id).First(&data).Error; err != nil {
		service.HTTPError(c, "查询失败", http.StatusNotFound, err)
		return
	}

	// 删除
	if err := db.DB.Delete(&data).Error; err != nil {
		service.HTTPError(c, "删除失败", http.StatusInternalServerError, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.HTTPDeleteSuccess{
		Data: data.ID,
	})
}

// UserFindOrCreate 查询或创建单条信息
// @Summary				查询或创建单条信息
// @Description			查询或创建单条信息...
// @Tags				user
// @Accept				json
// @Produce				json
// @Param				user body model.User true "查询或创建单条信息"
// @Success				200 {object} model.UserRes
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/user [post]
// @Security			ApiKeyAuth
func UserFindOrCreate(c *gin.Context) {

	// 绑定数据
	var data model.User
	if err := c.ShouldBind(&data); err != nil {
		service.HTTPError(c, "数据绑定失败", http.StatusBadRequest, err)
		return
	}

	// 插入数据
	if err := db.DB.FirstOrCreate(&data, data).Error; err != nil {
		service.HTTPError(c, "查询或创建数据失败", http.StatusInternalServerError, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.UserRes{
		Data: data,
	})
}

// UserUpdateBulk 批量更新
// @Summary				批量更新
// @Description			批量更新...
// @Tags				user
// @Accept				json
// @Produce				json
// @Param				user body model.BulkUpdate true "批量更新"
// @Success				200 {object} model.HTTPBulkSuccess
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/users [patch]
// @Security			ApiKeyAuth
func UserUpdateBulk(c *gin.Context) {

	var data model.BulkUpdate

	// 绑定数据
	if err := c.ShouldBind(&data); err != nil {
		service.HTTPError(c, "数据绑定失败", http.StatusBadRequest, err)
		return
	}

	// 判断ID列表是否为空
	// if len(data.IDs) == 0 {
	// 	service.HTTPError(c, "ids列表不能空", http.StatusBadRequest, nil)
	// 	return
	// }

	// 更新数据
	if err := db.DB.Model(model.User{}).Where("id IN (?)", data.IDs).Updates(data.Obj).Error; err != nil {
		service.HTTPError(c, "更新失败", http.StatusInternalServerError, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.HTTPBulkSuccess{
		Data: data.IDs,
	})
}

// UserDestroyBulk 批量删除
// @Summary				批量删除
// @Description			批量删除...
// @Tags				user
// @Accept				json
// @Produce				json
// @Param				user body model.BulkDelete true "批量删除"
// @Success				200 {object} model.HTTPBulkSuccess
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/users [delete]
// @Security			ApiKeyAuth
func UserDestroyBulk(c *gin.Context) {

	var data model.BulkDelete

	// 绑定数据
	if err := c.ShouldBind(&data); err != nil {
		service.HTTPError(c, "数据绑定失败", http.StatusBadRequest, err)
		return
	}

	// 删除数据
	if err := db.DB.Delete(model.User{}, "id IN (?)", data.IDs).Error; err != nil {
		service.HTTPError(c, "删除失败", http.StatusInternalServerError, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.HTTPBulkSuccess{
		Data: data.IDs,
	})
}

// UserFindMyRoles 查找本人角色
// @Summary				查找本人角色
// @Description			查找本人角色...
// @Tags				user
// @Accept				json
// @Produce				json
// @Success				200 {object} model.RoleListRes
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/user/roles [get]
// @Security			ApiKeyAuth
func UserFindMyRoles(c *gin.Context) {
	// 获取当前用户角色列表
	roles, err := service.GetCurrentUserRoles(c)
	if err != nil {
		service.HTTPError(c, "查询角色失败", http.StatusInternalServerError, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.RoleListRes{
		Data: roles,
	})
}

// UserFindMyGroups 查找本人组
// @Summary				查找本人组
// @Description			查找本人组...
// @Tags				user
// @Accept				json
// @Produce				json
// @Success				200 {object} model.UserGroupListRes
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/user/roles [get]
// @Security			ApiKeyAuth
func UserFindMyGroups(c *gin.Context) {
	// 获取当前用户组列表
	groups, err := service.GetCurrentUserGroups(c)
	if err != nil {
		service.HTTPError(c, "查询角色失败", http.StatusInternalServerError, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.UserGroupListRes{
		Data: groups,
	})
}

// UserFindMyMenus 查找本人菜单
// @Summary				查找本人菜单
// @Description			查找本人菜单...
// @Tags				user
// @Accept				json
// @Produce				json
// @Success				200 {object} model.MenuListRes
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/user/menus [get]
// @Security			ApiKeyAuth
func UserFindMyMenus(c *gin.Context) {
	// 获取当前用户菜单列表
	menus, err := service.GetCurrentUserMenus(c)
	if err != nil {
		service.HTTPError(c, "查询菜单失败", http.StatusInternalServerError, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.MenuListRes{
		Data: menus,
	})
}
