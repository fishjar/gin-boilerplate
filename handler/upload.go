package handler

import (
	"net/http"

	"github.com/fishjar/gin-boilerplate/model"
	"github.com/fishjar/gin-boilerplate/service"
	"github.com/gin-gonic/gin"
)

// UploudFile 文件上传，单文件
// @Summary				文件上传，单文件
// @Description			文件上传，单文件...
// @Tags				files
// @Accept				mpfd
// @Produce				json
// @Param				file formData file true "文件"
// @Success				200 {object} model.UploadSuccess
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/files/upload [post]
// @Security			ApiKeyAuth
func UploudFile(c *gin.Context) {
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		service.HTTPError(c, "文件上传失败：没有获取到文件", http.StatusBadRequest, err)
		return
	}

	data, err := service.SaveFile(c, file)
	if err != nil {
		service.HTTPError(c, "文件上传失败：保存文件失败", http.StatusInternalServerError, err)
		return
	}

	// 返回成功
	c.JSON(http.StatusOK, model.UploadSuccess{
		HTTPSuccess: model.HTTPSuccess{
			Message: "上传成功",
		},
		Data: data,
	})
}

// UploudMultiFile 文件上传，多文件
// @Summary				文件上传，多文件
// @Description			文件上传，多文件...
// @Tags				files
// @Accept				mpfd
// @Produce				json
// @Param				files formData file true "文件" collectionFormat(multi)
// @Success				200 {object} model.UploadMultiSuccess
// @Failure 			500 {object} model.HTTPError
// @Router				/admin/files/upload/multi [post]
// @Security			ApiKeyAuth
// swagger 不支持多文件上传
// 详见：https://swagger.io/docs/specification/2-0/file-upload/
// https://github.com/OAI/OpenAPI-Specification/issues/254
func UploudMultiFile(c *gin.Context) {
	// 获取文件
	form, err := c.MultipartForm()
	if err != nil {
		service.HTTPError(c, "文件上传失败：请使用form-data上传", http.StatusBadRequest, err)
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		service.HTTPError(c, "文件上传失败：没有获取到文件", http.StatusBadRequest, nil)
		return
	}

	var data = make([]model.UploadRes, len(files))
	for i, file := range files {
		res, _ := service.SaveFile(c, file)
		data[i] = res
	}

	// 返回成功
	c.JSON(http.StatusOK, model.UploadMultiSuccess{
		HTTPSuccess: model.HTTPSuccess{
			Message: "上传成功",
		},
		Data: data,
	})
}
