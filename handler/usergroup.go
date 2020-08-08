package handler

import (
	"net/http"

	"github.com/fishjar/gin-boilerplate/db"
	"github.com/fishjar/gin-boilerplate/model"
	"github.com/fishjar/gin-boilerplate/service"

	"github.com/gin-gonic/gin"
)

// UserGroupFindAndCountAll 查询多条信息
// @Summary				查询多条信息
// @Description			查询多条信息...
// @Tags				usergroup
// @Accept				json
// @Produce				json
// @Param				q query model.PaginReq false "参数"
// @Success				200 {object} model.UserGroupListRes
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/usergroups [get]
// @Security			ApiKeyAuth
func UserGroupFindAndCountAll(c *gin.Context) {

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
	// var where model.UserGroup
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
	var rows []model.UserGroup

	// 查询数据
	if err := db.DB.Model(&rows).Where(where).Count(&total).Limit(q.Size).Offset(offset).Order(q.Sort).Preload("User").Preload("Group").Find(&rows).Error; err != nil {
		service.HTTPError(c, "查询多条信息失败", http.StatusInternalServerError, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.UserGroupListRes{
		Pagin: q.Pagin(total),
		Data:  rows,
	})
}

// UserGroupFindByPk 根据主键查询单条信息
// @Summary				查询单条信息
// @Description			查询单条信息...
// @Tags				usergroup
// @Accept				json
// @Produce				json
// @Param				id path string true "ID"
// @Success				200 {object} model.UserGroupRes
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/usergroups/{id} [get]
// @Security			ApiKeyAuth
func UserGroupFindByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.UserGroup
	if err := db.DB.Preload("User").Preload("Group").First(&data, "id = ?", id).Error; err != nil {
		service.HTTPError(c, "查询失败", http.StatusNotFound, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.UserGroupRes{
		Data: data,
	})
}

// UserGroupSingleCreate 创建单条信息
// @Summary				创建单条信息
// @Description			创建单条信息...
// @Tags				usergroup
// @Accept				json
// @Produce				json
// @Param				usergroup body model.UserGroup true "参数"
// @Success				200 {object} model.UserGroupRes
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/usergroups [post]
// @Security			ApiKeyAuth
func UserGroupSingleCreate(c *gin.Context) {

	// 绑定数据
	var data model.UserGroup
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
	c.JSON(http.StatusOK, model.UserGroupRes{
		Data: data,
	})
}

// UserGroupUpdateByPk 更新单条信息
// @Summary				更新单条信息
// @Description			更新单条信息...
// @Tags				usergroup
// @Accept				json
// @Produce				json
// @Param				id path string true "ID"
// @Param				usergroup body model.UserGroup true "更新单条信息"
// @Success				200 {object} model.UserGroupRes
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/usergroups/{id} [patch]
// @Security			ApiKeyAuth
func UserGroupUpdateByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.UserGroup
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
	c.JSON(http.StatusOK, model.UserGroupRes{
		Data: data,
	})
}

// UserGroupDestroyByPk 删除单条信息
// @Summary				删除单条信息
// @Description			删除单条信息...
// @Tags				usergroup
// @Accept				json
// @Produce				json
// @Param				id path string true "ID"
// @Success				200 {object} model.HTTPDeleteSuccess
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/usergroups/{id} [delete]
// @Security			ApiKeyAuth
func UserGroupDestroyByPk(c *gin.Context) {

	// 获取参数
	id := c.Param("id")

	// 查询
	var data model.UserGroup
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
	c.JSON(http.StatusOK, model.HTTPDeleteSuccess{})
}

// UserGroupFindOrCreate 查询或创建单条信息
// @Summary				查询或创建单条信息
// @Description			查询或创建单条信息...
// @Tags				usergroup
// @Accept				json
// @Produce				json
// @Param				usergroup body model.UserGroup true "查询或创建单条信息"
// @Success				200 {object} model.UserGroupRes
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/usergroup [post]
// @Security			ApiKeyAuth
func UserGroupFindOrCreate(c *gin.Context) {

	// 绑定数据
	var data model.UserGroup
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
	c.JSON(http.StatusOK, model.UserGroupRes{
		Data: data,
	})
}

// UserGroupUpdateBulk 批量更新
// @Summary				批量更新
// @Description			批量更新...
// @Tags				usergroup
// @Accept				json
// @Produce				json
// @Param				usergroup body model.BulkUpdate true "批量更新"
// @Success				200 {object} model.HTTPBulkSuccess
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/usergroups [patch]
// @Security			ApiKeyAuth
func UserGroupUpdateBulk(c *gin.Context) {

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
	if err := db.DB.Model(model.UserGroup{}).Where("id IN (?)", data.IDs).Updates(data.Obj).Error; err != nil {
		service.HTTPError(c, "更新失败", http.StatusInternalServerError, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.HTTPBulkSuccess{
		Data: data.IDs,
	})
}

// UserGroupDestroyBulk 批量删除
// @Summary				批量删除
// @Description			批量删除...
// @Tags				usergroup
// @Accept				json
// @Produce				json
// @Param				usergroup body model.BulkDelete true "批量删除"
// @Success				200 {object} model.HTTPBulkSuccess
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/usergroups [delete]
// @Security			ApiKeyAuth
func UserGroupDestroyBulk(c *gin.Context) {

	var data model.BulkDelete

	// 绑定数据
	if err := c.ShouldBind(&data); err != nil {
		service.HTTPError(c, "数据绑定失败", http.StatusBadRequest, err)
		return
	}

	// 删除数据
	if err := db.DB.Delete(model.UserGroup{}, "id IN (?)", data.IDs).Error; err != nil {
		service.HTTPError(c, "删除失败", http.StatusInternalServerError, err)
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, model.HTTPBulkSuccess{
		Data: data.IDs,
	})
}
