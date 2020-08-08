package service

import (
	"errors"
	"mime/multipart"
	"net/http"
	"os"
	"path"

	"github.com/disintegration/imaging"
	"github.com/fishjar/gin-boilerplate/config"
	"github.com/fishjar/gin-boilerplate/model"
	"github.com/fishjar/gin-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

// SaveFile 保存文件
func SaveFile(c *gin.Context, file *multipart.FileHeader) (model.UploadRes, error) {
	// 判断文件类型
	contentType := GetFileHeaderContentType(file)

	// 获取文件扩展名
	orginName := file.Filename
	extName := path.Ext(orginName)
	if len(extName) == 0 { // 限制没有文件扩展名的文件
		return model.UploadRes{}, errors.New("文件扩展名有误")
	}

	// 计算文件md5
	fileName, err := utils.MD5File(file)
	if err != nil {
		return model.UploadRes{}, err
	}

	// 组装返回数据
	rootPath := config.Config.UploadFullPath
	filePath := path.Join(fileName[0:2], fileName[2:4])
	fullPath := path.Join(rootPath, filePath)
	fullName := fileName + extName
	pathName := path.Join(fullPath, fullName)
	url := path.Join(filePath, fullName)
	res := model.UploadRes{
		Origin: orginName,
		Type:   contentType,
		URL:    url,
		Isnew:  true,
	}

	// 判断是否图片
	if contentType == "image/png" || contentType == "image/jpg" {
		thumbName := fileName + "_256x256" + extName
		thumb := path.Join(filePath, thumbName)
		res.Thumb = thumb
	}

	// 判断文件是否已经存在
	if FileExist(pathName) {
		res.Isnew = false
		return res, nil
	}

	// 创建目录
	if err := os.MkdirAll(fullPath, 0755); err != nil {
		return model.UploadRes{}, err
	}

	// 保存文件
	if err := c.SaveUploadedFile(file, pathName); err != nil {
		return model.UploadRes{}, err
	}

	// 如果是图片格式，生成256*256缩略图
	// 其他大小的图片，在nginx中利用image_filter动态生成缩略图
	if contentType == "image/png" || contentType == "image/jpg" {
		thumbName := fileName + "_256x256" + extName
		thumbPathName := path.Join(fullPath, thumbName)
		if err := ImageResize(pathName, thumbPathName); err != nil {
			return model.UploadRes{}, err
		}
	}

	return res, nil
}

// GetFileContentType 获取文件类型
func GetFileContentType(file *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

// GetFileHeaderContentType 获取文件类型
func GetFileHeaderContentType(file *multipart.FileHeader) string {
	buffer := make([]byte, 512)
	tmpFile, _ := file.Open()
	defer tmpFile.Close()

	tmpFile.Read(buffer)
	contentType := http.DetectContentType(buffer)
	return contentType
}

// FileExist 检查文件或目录是否存在
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// ImageResize 图片处理
// https://github.com/h2non/bimg（依赖libvips，不推荐）
// https://github.com/disintegration/imaging（推荐）
func ImageResize(srcPath string, outPath string) error {
	// buffer, err := bimg.Read(srcPath)
	// if err != nil {
	// 	return err
	// }

	// newImage, err := bimg.NewImage(buffer).Thumbnail(256)
	// if err != nil {
	// 	return err
	// }

	// if err := bimg.Write(outPath, newImage); err != nil {
	// 	return err
	// }

	// return nil

	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	// Crop the original image
	src = imaging.Thumbnail(src, 256, 256, imaging.Lanczos)

	// Save the resulting image.
	err = imaging.Save(src, outPath)
	if err != nil {
		return err
	}

	return nil
}
